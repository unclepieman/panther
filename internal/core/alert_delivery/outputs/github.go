package outputs

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
	"strings"

	jsoniter "github.com/json-iterator/go"

	deliverymodel "github.com/panther-labs/panther/api/lambda/delivery/models"
	outputModels "github.com/panther-labs/panther/api/lambda/outputs/models"
)

// Severity colors match those in the Panther UI
const (
	githubEndpoint = "https://api.github.com/repos/"
	requestType    = "/issues"
)

// Github alert send an issue.
func (client *OutputClient) Github(
	ctx context.Context, alert *deliverymodel.Alert, config *outputModels.GithubConfig) *AlertDeliveryResponse {

	description := "**Description:** " + alert.AnalysisDescription
	link := "\n [Click here to view in the Panther UI](" + generateURL(alert) + ")"
	runBook := "\n **Runbook:** " + alert.Runbook
	severity := "\n **Severity:** " + alert.Severity
	tags := "\n **Tags:** " + strings.Join(alert.Tags, ", ")
	// Best effort attempt to marshal Alert Context
	marshaledContext, _ := jsoniter.MarshalToString(alert.Context)
	alertContext := "\n **AlertContext:** " + marshaledContext

	githubRequest := map[string]interface{}{
		"title": generateAlertTitle(alert),
		"body":  description + link + runBook + severity + tags + alertContext,
	}

	token := "token " + config.Token
	repoURL := githubEndpoint + config.RepoName + requestType
	requestHeader := map[string]string{
		AuthorizationHTTPHeader: token,
	}

	postInput := &PostInput{
		url:     repoURL,
		body:    githubRequest,
		headers: requestHeader,
	}
	return client.httpWrapper.post(ctx, postInput)
}
