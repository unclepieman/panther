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

/* eslint-disable import/order, import/no-duplicates, @typescript-eslint/no-unused-vars */

import * as Types from '../../../../__generated__/schema';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type SubscribeToDestinationUpdatesVariables = {};

export type SubscribeToDestinationUpdates = {
  updatedDestination?: Types.Maybe<
    Pick<
      Types.Destination,
      | 'createdBy'
      | 'creationTime'
      | 'displayName'
      | 'lastModifiedBy'
      | 'lastModifiedTime'
      | 'outputId'
      | 'outputType'
      | 'verificationStatus'
      | 'defaultForSeverity'
    > & {
      outputConfig: {
        slack?: Types.Maybe<Pick<Types.SlackConfig, 'webhookURL'>>;
        sns?: Types.Maybe<Pick<Types.SnsConfig, 'topicArn'>>;
        pagerDuty?: Types.Maybe<Pick<Types.PagerDutyConfig, 'integrationKey'>>;
        github?: Types.Maybe<Pick<Types.GithubConfig, 'repoName' | 'token'>>;
        jira?: Types.Maybe<
          Pick<
            Types.JiraConfig,
            'orgDomain' | 'projectKey' | 'userName' | 'apiKey' | 'assigneeId' | 'issueType'
          >
        >;
        opsgenie?: Types.Maybe<Pick<Types.OpsgenieConfig, 'apiKey'>>;
        msTeams?: Types.Maybe<Pick<Types.MsTeamsConfig, 'webhookURL'>>;
        sqs?: Types.Maybe<Pick<Types.SqsConfig, 'queueUrl'>>;
        asana?: Types.Maybe<Pick<Types.AsanaConfig, 'personalAccessToken' | 'projectGids'>>;
      };
    }
  >;
};

export const SubscribeToDestinationUpdatesDocument = gql`
  subscription SubscribeToDestinationUpdates {
    updatedDestination {
      createdBy
      creationTime
      displayName
      lastModifiedBy
      lastModifiedTime
      outputId
      outputType
      outputConfig {
        slack {
          webhookURL
        }
        sns {
          topicArn
        }
        pagerDuty {
          integrationKey
        }
        github {
          repoName
          token
        }
        jira {
          orgDomain
          projectKey
          userName
          apiKey
          assigneeId
          issueType
        }
        opsgenie {
          apiKey
        }
        msTeams {
          webhookURL
        }
        sqs {
          queueUrl
        }
        asana {
          personalAccessToken
          projectGids
        }
      }
      verificationStatus
      defaultForSeverity
    }
  }
`;

/**
 * __useSubscribeToDestinationUpdates__
 *
 * To run a query within a React component, call `useSubscribeToDestinationUpdates` and pass it any options that fit your needs.
 * When your component renders, `useSubscribeToDestinationUpdates` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the subscription, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSubscribeToDestinationUpdates({
 *   variables: {
 *   },
 * });
 */
export function useSubscribeToDestinationUpdates(
  baseOptions?: ApolloReactHooks.SubscriptionHookOptions<
    SubscribeToDestinationUpdates,
    SubscribeToDestinationUpdatesVariables
  >
) {
  return ApolloReactHooks.useSubscription<
    SubscribeToDestinationUpdates,
    SubscribeToDestinationUpdatesVariables
  >(SubscribeToDestinationUpdatesDocument, baseOptions);
}
export type SubscribeToDestinationUpdatesHookResult = ReturnType<
  typeof useSubscribeToDestinationUpdates
>;
export type SubscribeToDestinationUpdatesSubscriptionResult = ApolloReactCommon.SubscriptionResult<
  SubscribeToDestinationUpdates
>;
