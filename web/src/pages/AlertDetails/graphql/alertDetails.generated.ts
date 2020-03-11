/* eslint-disable import/order, import/no-duplicates */
import * as Types from '../../../../__generated__/schema';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type AlertDetailsVariables = {
  input: Types.GetAlertInput;
};

export type AlertDetails = {
  alert: Types.Maybe<
    Pick<
      Types.AlertDetails,
      | 'alertId'
      | 'ruleId'
      | 'creationTime'
      | 'eventsMatched'
      | 'updateTime'
      | 'eventsLastEvaluatedKey'
      | 'events'
    >
  >;
};

export type RuleTeaserVariables = {
  input: Types.GetRuleInput;
};

export type RuleTeaser = {
  rule: Types.Maybe<
    Pick<
      Types.RuleDetails,
      'description' | 'displayName' | 'id' | 'logTypes' | 'runbook' | 'severity' | 'tags'
    >
  >;
};

export const AlertDetailsDocument = gql`
  query AlertDetails($input: GetAlertInput!) {
    alert(input: $input) {
      alertId
      ruleId
      creationTime
      eventsMatched
      updateTime
      eventsLastEvaluatedKey
      events
    }
  }
`;

/**
 * __useAlertDetails__
 *
 * To run a query within a React component, call `useAlertDetails` and pass it any options that fit your needs.
 * When your component renders, `useAlertDetails` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAlertDetails({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useAlertDetails(
  baseOptions?: ApolloReactHooks.QueryHookOptions<AlertDetails, AlertDetailsVariables>
) {
  return ApolloReactHooks.useQuery<AlertDetails, AlertDetailsVariables>(
    AlertDetailsDocument,
    baseOptions
  );
}
export function useAlertDetailsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<AlertDetails, AlertDetailsVariables>
) {
  return ApolloReactHooks.useLazyQuery<AlertDetails, AlertDetailsVariables>(
    AlertDetailsDocument,
    baseOptions
  );
}
export type AlertDetailsHookResult = ReturnType<typeof useAlertDetails>;
export type AlertDetailsLazyQueryHookResult = ReturnType<typeof useAlertDetailsLazyQuery>;
export type AlertDetailsQueryResult = ApolloReactCommon.QueryResult<
  AlertDetails,
  AlertDetailsVariables
>;
export const RuleTeaserDocument = gql`
  query RuleTeaser($input: GetRuleInput!) {
    rule(input: $input) {
      description
      displayName
      id
      logTypes
      runbook
      severity
      tags
    }
  }
`;

/**
 * __useRuleTeaser__
 *
 * To run a query within a React component, call `useRuleTeaser` and pass it any options that fit your needs.
 * When your component renders, `useRuleTeaser` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useRuleTeaser({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useRuleTeaser(
  baseOptions?: ApolloReactHooks.QueryHookOptions<RuleTeaser, RuleTeaserVariables>
) {
  return ApolloReactHooks.useQuery<RuleTeaser, RuleTeaserVariables>(
    RuleTeaserDocument,
    baseOptions
  );
}
export function useRuleTeaserLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<RuleTeaser, RuleTeaserVariables>
) {
  return ApolloReactHooks.useLazyQuery<RuleTeaser, RuleTeaserVariables>(
    RuleTeaserDocument,
    baseOptions
  );
}
export type RuleTeaserHookResult = ReturnType<typeof useRuleTeaser>;
export type RuleTeaserLazyQueryHookResult = ReturnType<typeof useRuleTeaserLazyQuery>;
export type RuleTeaserQueryResult = ApolloReactCommon.QueryResult<RuleTeaser, RuleTeaserVariables>;
