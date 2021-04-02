package awsglue

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/glue/glueiface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/panther-labs/panther/internal/log_analysis/awsglue/glueschema"
	"github.com/panther-labs/panther/internal/log_analysis/pantherdb"
	"github.com/panther-labs/panther/pkg/awsutils"
	"github.com/panther-labs/panther/pkg/box"
)

type PartitionKey struct {
	Name string
	Type string
}

// Metadata about Glue table
type GlueTableMetadata struct {
	databaseName string
	tableName    string
	description  string
	prefix       string
	timebin      GlueTableTimebin // at what time resolution is this table partitioned
	eventStruct  interface{}
}

// Creates a new GlueTableMetadata object for Panther log sources
func NewGlueTableMetadata(
	database, table, logDescription string, timebin GlueTableTimebin, eventStruct interface{}) *GlueTableMetadata {

	tablePrefix := TablePrefix(database, table)
	return &GlueTableMetadata{
		databaseName: database,
		tableName:    table,
		description:  logDescription,
		timebin:      timebin,
		prefix:       tablePrefix,
		eventStruct:  eventStruct,
	}
}

func (gm *GlueTableMetadata) DatabaseName() string {
	return gm.databaseName
}

func (gm *GlueTableMetadata) TableName() string {
	return gm.tableName
}

func (gm *GlueTableMetadata) Description() string {
	return gm.description
}

// All data for this table are stored in this S3 prefix
func (gm *GlueTableMetadata) Prefix() string {
	return gm.prefix
}

func (gm *GlueTableMetadata) Timebin() GlueTableTimebin {
	return gm.timebin
}

func (gm *GlueTableMetadata) EventStruct() interface{} {
	return gm.eventStruct
}

func (gm *GlueTableMetadata) HasPartitions(glueClient glueiface.GlueAPI) (bool, error) {
	return TableHasPartitions(glueClient, gm.databaseName, gm.tableName)
}

// The partition keys for this table
func (gm *GlueTableMetadata) PartitionKeys() (partitions []PartitionKey) {
	partitions = []PartitionKey{{Name: "year", Type: "int"}}

	if gm.Timebin() >= GlueTableMonthly {
		partitions = append(partitions, PartitionKey{Name: "month", Type: "int"})
	}
	if gm.Timebin() >= GlueTableDaily {
		partitions = append(partitions, PartitionKey{Name: "day", Type: "int"})
	}
	if gm.Timebin() >= GlueTableHourly {
		partitions = append(partitions, PartitionKey{Name: "hour", Type: "int"})
		partitions = append(partitions, PartitionKey{Name: "partition_time", Type: "bigint"})
	}
	return partitions
}

func (gm *GlueTableMetadata) RuleTable() *GlueTableMetadata {
	if gm.databaseName == pantherdb.RuleMatchDatabase {
		return gm
	}
	// the corresponding rule table shares the same structure as the log table + some columns
	return NewGlueTableMetadata(pantherdb.RuleMatchDatabase, gm.tableName, gm.Description(), GlueTableHourly, gm.EventStruct())
}

func (gm *GlueTableMetadata) RuleErrorTable() *GlueTableMetadata {
	if gm.databaseName == pantherdb.RuleMatchDatabase {
		return gm
	}
	// the corresponding rule table shares the same structure as the log table + some columns
	return NewGlueTableMetadata(pantherdb.RuleErrorsDatabase, gm.tableName, gm.Description(), GlueTableHourly, gm.EventStruct())
}

func (gm *GlueTableMetadata) glueTableInput(bucketName string) (*glue.TableInput, error) {
	// partition keys -> []*glue.Column
	partitionKeys := gm.PartitionKeys()
	partitionColumns := make([]*glue.Column, len(partitionKeys))
	for i := range partitionKeys {
		partitionColumns[i] = &glue.Column{
			Name: &partitionKeys[i].Name,
			Type: &partitionKeys[i].Type,
		}
	}

	// columns -> []*glue.Column
	columns, mappings, err := glueschema.InferColumnsWithMappings(gm.EventStruct())
	if err != nil {
		return nil, err
	}
	switch gm.databaseName {
	case pantherdb.RuleMatchDatabase:
		// append the columns added by the rule engine
		columns = append(columns, RuleMatchColumns...)
	case pantherdb.RuleErrorsDatabase:
		// append the rule error columns
		columns = append(columns, RuleErrorColumns...)
	}
	glueColumns := make([]*glue.Column, len(columns))
	for i := range columns {
		glueColumns[i] = &glue.Column{
			Name:    &columns[i].Name,
			Type:    (*string)(&columns[i].Type),
			Comment: &columns[i].Comment,
		}
	}

	// Need to be case sensitive to deal with columns that have same name but different casing
	// https://github.com/rcongiu/Hive-JSON-Serde#case-sensitivity-in-mappings
	descriptorParameters := map[string]*string{
		"serialization.format": aws.String("1"),
		"case.insensitive":     aws.String("false"),
	}

	// Add mapping for all field names. This is required when columns are case sensitive
	for from, to := range mappings {
		to := to
		descriptorParameters[fmt.Sprintf("mapping.%s", from)] = &to
	}

	return &glue.TableInput{
		Name:          &gm.tableName,
		Description:   &gm.description,
		PartitionKeys: partitionColumns,
		StorageDescriptor: &glue.StorageDescriptor{ // configure as JSON
			Columns:      glueColumns,
			Location:     aws.String("s3://" + bucketName + "/" + gm.prefix),
			InputFormat:  aws.String("org.apache.hadoop.mapred.TextInputFormat"),
			OutputFormat: aws.String("org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat"),
			SerdeInfo: &glue.SerDeInfo{
				SerializationLibrary: aws.String("org.openx.data.jsonserde.JsonSerDe"),
				Parameters:           descriptorParameters,
			},
		},
		TableType: aws.String("EXTERNAL_TABLE"),
	}, nil
}

func (gm *GlueTableMetadata) UpdateTableIfExists(ctx context.Context, glueAPI glueiface.GlueAPI, bucketName string) (bool, error) {
	tableInput, err := gm.glueTableInput(bucketName)
	if err != nil {
		return false, err
	}
	updateTableInput := &glue.UpdateTableInput{
		DatabaseName: &gm.databaseName,
		TableInput:   tableInput,
	}
	if _, err := glueAPI.UpdateTableWithContext(ctx, updateTableInput); err != nil {
		if awsutils.IsAnyError(err, glue.ErrCodeEntityNotFoundException) {
			return false, nil
		}
		return false, errors.Wrapf(err, "failed to update table %s.%s", gm.databaseName, gm.tableName)
	}

	return true, nil
}

func (gm *GlueTableMetadata) CreateTableIfNotExists(ctx context.Context, glueAPI glueiface.GlueAPI, bucketName string) (bool, error) {
	tableInput, err := gm.glueTableInput(bucketName)
	if err != nil {
		return false, err
	}
	createTableInput := &glue.CreateTableInput{
		DatabaseName: &gm.databaseName,
		TableInput:   tableInput,
		PartitionIndexes: []*glue.PartitionIndex{
			{
				IndexName: aws.String("month_idx"),
				Keys: []*string{
					aws.String("year"),
					aws.String("month"),
				},
			},
			{
				IndexName: aws.String("day_idx"),
				Keys: []*string{
					aws.String("year"),
					aws.String("month"),
					aws.String("day"),
				},
			},
			{
				IndexName: aws.String("partition_time_idx"),
				Keys: []*string{
					aws.String("partition_time"),
				},
			},
		},
	}
	if _, err := glueAPI.CreateTableWithContext(ctx, createTableInput); err != nil {
		if awsutils.IsAnyError(err, glue.ErrCodeAlreadyExistsException) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (gm *GlueTableMetadata) CreateOrUpdateTable(glueClient glueiface.GlueAPI, bucketName string) error {
	tableInput, err := gm.glueTableInput(bucketName)
	if err != nil {
		return err
	}

	createTableInput := &glue.CreateTableInput{
		DatabaseName: &gm.databaseName,
		TableInput:   tableInput,
		PartitionIndexes: []*glue.PartitionIndex{
			{
				IndexName: aws.String("month_idx"),
				Keys: []*string{
					aws.String("year"),
					aws.String("month"),
				},
			},
			{
				IndexName: aws.String("day_idx"),
				Keys: []*string{
					aws.String("year"),
					aws.String("month"),
					aws.String("day"),
				},
			},
			{
				IndexName: aws.String("partition_time_idx"),
				Keys: []*string{
					aws.String("partition_time"),
				},
			},
		},
	}
	if _, err := glueClient.CreateTable(createTableInput); err != nil {
		if awsutils.IsAnyError(err, glue.ErrCodeAlreadyExistsException) {
			// need to do an update
			updateTableInput := &glue.UpdateTableInput{
				DatabaseName: &gm.databaseName,
				TableInput:   tableInput,
			}
			_, errUpdate := glueClient.UpdateTable(updateTableInput)
			if errUpdate != nil {
				return errors.Wrapf(errUpdate, "failed to update table %s.%s", gm.databaseName, gm.tableName)
			}
			for _, index := range createTableInput.PartitionIndexes {
				createPartitionIndexInput := &glue.CreatePartitionIndexInput{
					DatabaseName:   &gm.databaseName,
					PartitionIndex: index,
					TableName:      &gm.tableName,
				}
				_, errLoop := glueClient.CreatePartitionIndex(createPartitionIndexInput)
				if errLoop != nil {
					if awsutils.IsAnyError(errLoop, glue.ErrCodeAlreadyExistsException) {
						continue
					}
					return errors.Wrapf(errLoop, "failed to create index %s for table %s.%s",
						*index.IndexName, gm.databaseName, gm.tableName)
				}
			}
			return nil
		}
		return errors.Wrapf(err, "failed to create table %s.%s", gm.databaseName, gm.tableName)
	}
	return nil
}

// Based on Timebin(), return an S3 prefix for objects of this table
func (gm *GlueTableMetadata) PartitionPrefix(t time.Time) string {
	return gm.Prefix() + gm.timebin.PartitionPathS3(t)
}

// SyncPartitions updates a table's partitions using the latest table schema. Used when schemas change.
// If deadline is non-nil, it will stop when execution time has passed the deadline and will return the
// _next_ time period needing evaluation. Deadlines are used when this is called in Lambdas to avoid
// running past the lambda deadline.
func (gm *GlueTableMetadata) SyncPartitions(glueClient glueiface.GlueAPI, s3Client s3iface.S3API,
	startDate time.Time, deadline *time.Time) (*time.Time, error) {

	// inherit StorageDescriptor from table
	tableOutput, err := GetTable(glueClient, gm.databaseName, gm.tableName)
	if err != nil {
		return nil, err
	}

	columns := tableOutput.Table.StorageDescriptor.Columns
	if startDate.IsZero() {
		startDate = *tableOutput.Table.CreateTime
	}
	startDate = startDate.Truncate(time.Hour * 24) // clip to beginning of day
	// update to current day at last hour
	endDay := time.Now().UTC().Truncate(time.Hour * 24).Add(time.Hour * 23)

	const concurrency = 10
	updateChan := make(chan time.Time, concurrency)
	errChan := make(chan error, concurrency)
	// update concurrently cuz the Glue API is very slow
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			failed := false
			for update := range updateChan {
				if failed {
					continue // drain channel
				}

				values := gm.timebin.PartitionValuesFromTime(update)

				getPartitionOutput, err := GetPartition(glueClient, gm.databaseName, gm.tableName, values)
				if err != nil {
					// skip time period with no partition UNLESS there is data, then create
					var awsErr awserr.Error
					if !errors.As(err, &awsErr) || awsErr.Code() != glue.ErrCodeEntityNotFoundException {
						failed = true
						errChan <- err
					} else { // no partition, check if there is data in S3, if so, create
						if hasData, err := gm.timebin.PartitionHasData(s3Client, update, tableOutput); err != nil {
							failed = true
							errChan <- err
						} else if hasData {
							if _, err = gm.createPartition(glueClient, update, tableOutput); err != nil {
								failed = true
								errChan <- err
							}
						}
					}
					continue
				}

				// leave _everything_ the same except the schema, and the serde info to get column mappings
				storageDescriptor := *getPartitionOutput.Partition.StorageDescriptor // copy because we will mutate
				storageDescriptor.Columns = columns
				// we need to update the SerDeInfo for JSON partitions to get the column mappings
				if IsJSONPartition(&storageDescriptor) {
					storageDescriptor.SerdeInfo = tableOutput.Table.StorageDescriptor.SerdeInfo
				}
				_, err = UpdatePartition(glueClient, gm.databaseName, gm.tableName, values,
					&storageDescriptor, nil)
				if err != nil {
					failed = true
					errChan <- err
					continue
				}
			}
		}()
	}

	var nextTimeBin *time.Time // set and returned if deadline passed
	var isDeadlinePassed func() bool
	if deadline != nil {
		isDeadlinePassed = func() bool {
			return time.Now().UTC().After(*deadline)
		}
	} else {
		isDeadlinePassed = func() bool { return false }
	}

	// loop over each partition updating, stop if past deadline
	for timeBin := startDate; !timeBin.After(endDay); timeBin = gm.Timebin().Next(timeBin) {
		if isDeadlinePassed() {
			nextTimeBin = box.Time(timeBin)
			break
		}
		updateChan <- timeBin
	}

	close(updateChan)
	wg.Wait()

	close(errChan)
	return nextTimeBin, <-errChan
}

func (gm *GlueTableMetadata) CreateJSONPartition(client glueiface.GlueAPI, t time.Time) (created bool, err error) {
	// inherit StorageDescriptor from table
	tableOutput, err := GetTable(client, gm.databaseName, gm.tableName)
	if err != nil {
		return false, err
	}

	// ensure this is a JSON table, use Contains() because there are multiple json serdes
	if !IsJSONPartition(tableOutput.Table.StorageDescriptor) {
		return false, errors.Errorf("not a JSON table: %#v", *tableOutput.Table.StorageDescriptor)
	}

	return gm.createPartition(client, t, tableOutput)
}

func (gm *GlueTableMetadata) createPartition(client glueiface.GlueAPI, t time.Time,
	tableOutput *glue.GetTableOutput) (created bool, err error) {

	bucket, _, err := ParseS3URL(*tableOutput.Table.StorageDescriptor.Location)
	if err != nil {
		return false, err
	}

	storageDescriptor := *tableOutput.Table.StorageDescriptor // copy because we will mutate
	storageDescriptor.Location = aws.String("s3://" + bucket + "/" + gm.PartitionPrefix(t))

	// FIXME: This can be removed in a few releases after 1.16 after everyone has upgraded.
	// FIXME: Release 1.16 adds a new partition column to the tables in Glue.
	// FIXME: The custom resource will run AFTER this lambda has been updated and run a sync to update all tables
	partitionValues := gm.timebin.PartitionValuesFromTime(t)
	if len(partitionValues) != len(tableOutput.Table.PartitionKeys) { // this means we have new code and old schema
		// the last value is partition_time, take off to match old schema
		partitionValues = partitionValues[0 : len(partitionValues)-1]
		zap.L().Warn("creating backwards compatible partition", // log so we know then this is happening
			zap.String("partition", t.String()))
	}

	_, err = CreatePartition(client, gm.databaseName, gm.tableName, partitionValues, &storageDescriptor, nil)
	if err != nil {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) && awsErr.Code() == glue.ErrCodeAlreadyExistsException {
			return false, nil // no error
		}
		return false, err
	}
	return true, nil
}

// get partition, return nil if it does not exist
func (gm *GlueTableMetadata) GetPartition(client glueiface.GlueAPI, t time.Time) (output *glue.GetPartitionOutput, err error) {
	output, err = GetPartition(client, gm.databaseName, gm.tableName, gm.timebin.PartitionValuesFromTime(t))
	if err != nil {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) && awsErr.Code() == glue.ErrCodeEntityNotFoundException {
			return nil, nil // not there, no error
		}
		return nil, err
	}
	return output, err
}

func (gm *GlueTableMetadata) deletePartition(client glueiface.GlueAPI, t time.Time) (output *glue.DeletePartitionOutput, err error) {
	return DeletePartition(client, gm.databaseName, gm.tableName, gm.timebin.PartitionValuesFromTime(t))
}
