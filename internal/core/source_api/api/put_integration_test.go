package api

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
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/sqs"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/panther-labs/panther/api/lambda/source/models"
	pollermodels "github.com/panther-labs/panther/internal/compliance/snapshot_poller/models/poller"
	awspoller "github.com/panther-labs/panther/internal/compliance/snapshot_poller/pollers/aws"
	"github.com/panther-labs/panther/internal/core/source_api/ddb"
	"github.com/panther-labs/panther/internal/core/source_api/ddb/modelstest"
)

func generateMockSQSBatchInputOutput(integration models.SourceIntegrationMetadata) (
	*sqs.SendMessageBatchInput, *sqs.SendMessageBatchOutput, error) {

	// Setup input/output
	var sqsEntries []*sqs.SendMessageBatchRequestEntry
	var err error
	in := &sqs.SendMessageBatchInput{
		QueueUrl: aws.String("test-url"),
	}
	out := &sqs.SendMessageBatchOutput{
		Successful: []*sqs.SendMessageBatchResultEntry{
			{
				Id:               &integration.IntegrationID,
				MessageId:        &integration.IntegrationID,
				MD5OfMessageBody: aws.String("f6255bb01c648fe967714d52a89e8e9c"),
			},
		},
	}

	// Generate all messages for scans
	for resourceType := range awspoller.ServicePollers {
		scanMsg := &pollermodels.ScanMsg{
			Entries: []*pollermodels.ScanEntry{
				{
					AWSAccountID:  &integration.AWSAccountID,
					IntegrationID: &integration.IntegrationID,
					ResourceType:  aws.String(resourceType),
				},
			},
		}

		var messageBody string
		messageBody, err = jsoniter.MarshalToString(scanMsg)
		if err != nil {
			break
		}

		sqsEntries = append(sqsEntries, &sqs.SendMessageBatchRequestEntry{
			Id:          &integration.IntegrationID,
			MessageBody: &messageBody,
		})
	}

	in.Entries = sqsEntries
	return in, out, err
}

// Unit Tests

func TestAddToSnapshotQueue(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.Config.SnapshotPollersQueueURL = "test-url"
	testIntegration := models.SourceIntegrationMetadata{
		AWSAccountID:     testAccountID,
		CreatedAtTime:    time.Now(),
		CreatedBy:        "Bobert",
		IntegrationID:    testIntegrationID,
		IntegrationLabel: "BobertTest",
		IntegrationType:  models.IntegrationTypeAWSScan,
		ScanIntervalMins: 60,
	}

	sqsIn, sqsOut, err := generateMockSQSBatchInputOutput(testIntegration)
	require.NoError(t, err)

	// It's non trivial to mock when the order of a slice is not promised
	apiTest.mockSqs.On("SendMessageBatch", mock.Anything).Return(sqsOut, nil)

	err = apiTest.FullScan(&models.FullScanInput{Integrations: []*models.SourceIntegrationMetadata{&testIntegration}})

	require.NoError(t, err)
	// Check that there is one message per service
	assert.Len(t, sqsIn.Entries, len(awspoller.ServicePollers))
	apiTest.AssertExpectations(t)
}

func TestPutCloudSecIntegration(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.DdbClient = &ddb.DDB{Client: &modelstest.MockDDBClient{TestErr: false}, TableName: "test"}
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	// Message sent to create Cloud Security tables
	// TODO Verify payload
	apiTest.mockSqs.On("SendMessageWithContext", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)
	// This message is to start full scan
	apiTest.mockSqs.On("SendMessageBatch", mock.Anything).Return(&sqs.SendMessageBatchOutput{}, nil)

	out, err := apiTest.PutIntegration(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWSScan,
			ScanIntervalMins: 60,
			UserID:           testUserID,
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, out)
	apiTest.AssertExpectations(t)
}

func TestPutLogIntegrationExists(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	apiTest.DdbClient = &ddb.DDB{
		Client: &modelstest.MockDDBClient{
			MockScanAttributes: []map[string]*dynamodb.AttributeValue{
				{
					"awsAccountId":     {S: aws.String(testAccountID)},
					"integrationType":  {S: aws.String(models.IntegrationTypeAWS3)},
					"integrationlabel": {S: aws.String(testIntegrationLabel)},
				},
			},
			TestErr: false,
		},
		TableName: "test",
	}

	out, err := apiTest.PutIntegration(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWS3,
			ScanIntervalMins: 60,
			UserID:           testUserID,
		},
	})
	require.Error(t, err)
	require.Empty(t, out)
	assert.Equal(t, "Log source for account 123456789012 with label ProdAWS already onboarded", err.Error())
	apiTest.AssertExpectations(t)
}

func TestPutCloudSecIntegrationExists(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	apiTest.DdbClient = &ddb.DDB{
		Client: &modelstest.MockDDBClient{
			MockScanAttributes: []map[string]*dynamodb.AttributeValue{
				{
					"awsAccountId":     {S: aws.String(testAccountID)},
					"integrationType":  {S: aws.String(models.IntegrationTypeAWSScan)},
					"integrationlabel": {S: aws.String(testIntegrationLabel)},
				},
			},
			TestErr: false,
		},
		TableName: "test",
	}

	out, err := apiTest.PutIntegration(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWSScan,
			ScanIntervalMins: 60,
			UserID:           testUserID,
		},
	})
	require.Error(t, err)
	require.Empty(t, out)
	assert.Equal(t, "Source account 123456789012 already onboarded", err.Error())
}

func TestPutIntegrationValidInput(t *testing.T) {
	t.Parallel()
	validator, err := models.Validator()
	require.NoError(t, err)
	assert.NoError(t, validator.Struct(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWSScan,
			ScanIntervalMins: 60,
			UserID:           testUserID,
		},
	}))
}

func TestPutIntegrationInvalidInput(t *testing.T) {
	t.Parallel()
	validator, err := models.Validator()
	require.NoError(t, err)
	assert.Error(t, validator.Struct(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  "type doesn't exist",
			ScanIntervalMins: 60,
			UserID:           testUserID,
		},
	}))
}

func TestPutIntegrationDatabaseError(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	in := &models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWSScan,
			UserID:           testUserID,
		},
	}
	apiTest.DdbClient = &ddb.DDB{
		Client: &modelstest.MockDDBClient{
			TestErr: true,
		},
		TableName: "test",
	}

	out, err := apiTest.PutIntegration(in)
	assert.Error(t, err)
	assert.Empty(t, out)
	assert.Equal(t, "Failed to add source. Please try again later", err.Error())
	apiTest.AssertExpectations(t)
}

func TestPutLogIntegrationUpdateSqsQueuePermissions(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.DdbClient = &ddb.DDB{Client: &modelstest.MockDDBClient{TestErr: false}, TableName: "test"}

	apiTest.Config.LogProcessorQueueURL = "https://sqs.eu-west-1.amazonaws.com/123456789012/testqueue"
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	expectedGetQueueAttributesInput := &sqs.GetQueueAttributesInput{
		AttributeNames: aws.StringSlice([]string{"Policy"}),
		QueueUrl:       &apiTest.Config.LogProcessorQueueURL,
	}
	alreadyExistingAttributes := generateQueueAttributeOutput(t, []string{})
	apiTest.mockSqs.On("GetQueueAttributes", expectedGetQueueAttributesInput).
		Return(&sqs.GetQueueAttributesOutput{Attributes: alreadyExistingAttributes}, nil).Once()
	expectedAttributes := generateQueueAttributeOutput(t, []string{testAccountID})
	expectedSetAttributes := &sqs.SetQueueAttributesInput{
		Attributes: expectedAttributes,
		QueueUrl:   &apiTest.Config.LogProcessorQueueURL,
	}
	apiTest.mockSqs.On("SetQueueAttributes", expectedSetAttributes).Return(&sqs.SetQueueAttributesOutput{}, nil).Once()
	apiTest.mockSqs.On("SendMessageWithContext", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

	out, err := apiTest.PutIntegration(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWS3,
			UserID:           testUserID,
			S3Bucket:         "bucket",
			KmsKey:           "keyarns",
			S3PrefixLogTypes: models.S3PrefixLogtypes{{S3Prefix: "", LogTypes: []string{"AWS.VPCFlow"}}},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, out)
	apiTest.AssertExpectations(t)
}

func TestPutLogIntegrationUpdateSqsQueuePermissionsFailure(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.DdbClient = &ddb.DDB{Client: &modelstest.MockDDBClient{TestErr: false}, TableName: "test"}
	apiTest.Config.LogProcessorQueueURL = "https://sqs.eu-west-1.amazonaws.com/123456789012/testqueue"
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	apiTest.mockSqs.On("GetQueueAttributes", mock.Anything).Return(&sqs.GetQueueAttributesOutput{}, errors.New("error")).Once()
	apiTest.mockSqs.On("SendMessageWithContext", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

	out, err := apiTest.PutIntegration(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			AWSAccountID:     testAccountID,
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeAWS3,
			UserID:           testUserID,
			S3Bucket:         "bucket",
			KmsKey:           "keyarns",
			S3PrefixLogTypes: models.S3PrefixLogtypes{{S3Prefix: "", LogTypes: []string{"AWS.VPCFlow"}}},
		},
	})
	require.Error(t, err)
	require.Empty(t, out)
	apiTest.AssertExpectations(t)
}

func TestPutSqsIntegration(t *testing.T) {
	t.Parallel()
	apiTest := NewAPITest()
	apiTest.DdbClient = &ddb.DDB{Client: &modelstest.MockDDBClient{TestErr: false}, TableName: "test"}

	apiTest.Config.LogProcessorQueueURL = "https://sqs.eu-west-1.amazonaws.com/123456789012/testqueue"
	apiTest.Config.AccountID = "123456789012"
	apiTest.Config.InputDataBucketName = "input-data"
	apiTest.Config.InputDataRoleArn = "role-arn"
	apiTest.Config.Region = "eu-west-1"
	apiTest.EvaluateIntegrationFunc = func(_ *models.CheckIntegrationInput) (string, bool, error) { return "", true, nil }

	// Configuring the Log Processor SQS queue
	alreadyExistingAttributes := generateQueueAttributeOutput(t, []string{})
	apiTest.mockSqs.On("GetQueueAttributes", mock.Anything).
		Return(&sqs.GetQueueAttributesOutput{Attributes: alreadyExistingAttributes}, nil).Once()
	apiTest.mockSqs.On("SetQueueAttributes", mock.Anything).Return(&sqs.SetQueueAttributesOutput{}, nil).Once()

	// Create a new SQS queue - we are verifying the parameters below
	apiTest.mockSqs.On("CreateQueue", mock.Anything).Return(&sqs.CreateQueueOutput{}, nil).Once()
	apiTest.mockSqs.On("SendMessageWithContext", mock.Anything, mock.Anything).Return(&sqs.SendMessageOutput{}, nil)

	apiTest.mockLambda.On("CreateEventSourceMapping", mock.Anything).Return(&lambda.EventSourceMappingConfiguration{}, nil)

	out, err := apiTest.PutIntegration(&models.PutIntegrationInput{
		PutIntegrationSettings: models.PutIntegrationSettings{
			IntegrationLabel: testIntegrationLabel,
			IntegrationType:  models.IntegrationTypeSqs,
			SqsConfig: &models.SqsConfig{
				LogTypes:             []string{"AWS.CloudTrail"},
				AllowedPrincipalArns: []string{"arn:aws:iam::123456789012:root"},
				AllowedSourceArns:    []string{"arn:aws:sns:*:415773754570:*"},
			},
		},
	})

	// Verify returned values
	require.NoError(t, err)
	require.NotEmpty(t, out)
	bucket, prefixes := out.S3Info()
	assert.Equal(t, "input-data", bucket)
	assert.Equal(t, []string{"forwarder"}, prefixes)
	assert.Equal(t, "role-arn", out.RequiredLogProcessingRole())
	assert.Equal(t, []string{"AWS.CloudTrail"}, out.RequiredLogTypes())

	// Verify SQS queue was created the appropriate permissions
	createQueueRequest := apiTest.mockSqs.Calls[3].Arguments.Get(0).(*sqs.CreateQueueInput)
	// nolint:lll
	expectedSqsQueuePolicy := `
{
"Version":"2008-10-17",
"Statement":[
{"Action":"sqs:SendMessage", "Effect":"Allow", "Principal":{"AWS":"arn:aws:iam::123456789012:root"}, "Resource":"*", "Sid":"arn:aws:iam::123456789012:root"},
{"Action":"sqs:SendMessage", "Effect":"Allow", "Principal":{"AWS":"*"}, "Resource":"*", "Condition":{"ArnLike":{"aws:SourceArn":"arn:aws:sns:*:415773754570:*"}}, "Sid":"arn:aws:sns:*:415773754570:*"}
]
}
`
	assert.JSONEq(t, expectedSqsQueuePolicy, *createQueueRequest.Attributes["Policy"])
	apiTest.AssertExpectations(t)
}
