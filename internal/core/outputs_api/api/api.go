// Package api defines CRUD actions for Panther alert outputs.
package api

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
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
	"os"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/panther-labs/panther/internal/core/outputs_api/encryption"
	"github.com/panther-labs/panther/internal/core/outputs_api/table"
	"github.com/panther-labs/panther/internal/core/outputs_api/verification"
)

// The API consists of receiver methods for each of the handlers.
type API struct{}

var (
	awsSession = session.Must(session.NewSession())

	encryptionKey encryption.API = encryption.New(os.Getenv("KEY_ID"), awsSession)

	outputsTable table.OutputsAPI = table.NewOutputs(
		os.Getenv("OUTPUTS_TABLE_NAME"),
		os.Getenv("OUTPUTS_DISPLAY_NAME_INDEX_NAME"),
		awsSession)
	defaultsTable table.DefaultsAPI = table.NewDefaults(
		os.Getenv("DEFAULTS_TABLE_NAME"),
		awsSession)

	outputVerification verification.OutputVerificationAPI = verification.NewVerification(awsSession)
)
