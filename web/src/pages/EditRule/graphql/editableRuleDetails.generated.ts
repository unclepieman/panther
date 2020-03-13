/* eslint-disable import/order, import/no-duplicates */
import * as Types from '../../../../__generated__/schema';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type EditableRuleDetailsVariables = {
  input: Types.GetRuleInput;
};

export type EditableRuleDetails = {
  rule: Types.Maybe<
    Pick<
      Types.RuleDetails,
      | 'description'
      | 'displayName'
      | 'enabled'
      | 'id'
      | 'reference'
      | 'logTypes'
      | 'runbook'
      | 'severity'
      | 'tags'
      | 'body'
    > & {
      tests: Types.Maybe<
        Array<
          Types.Maybe<
            Pick<Types.PolicyUnitTest, 'expectedResult' | 'name' | 'resource' | 'resourceType'>
          >
        >
      >;
    }
  >;
};

export const EditableRuleDetailsDocument = gql`
  query EditableRuleDetails($input: GetRuleInput!) {
    rule(input: $input) {
      description
      displayName
      enabled
      id
      reference
      logTypes
      runbook
      severity
      tags
      body
      tests {
        expectedResult
        name
        resource
        resourceType
      }
    }
  }
`;

/**
 * __useEditableRuleDetails__
 *
 * To run a query within a React component, call `useEditableRuleDetails` and pass it any options that fit your needs.
 * When your component renders, `useEditableRuleDetails` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useEditableRuleDetails({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useEditableRuleDetails(
  baseOptions?: ApolloReactHooks.QueryHookOptions<EditableRuleDetails, EditableRuleDetailsVariables>
) {
  return ApolloReactHooks.useQuery<EditableRuleDetails, EditableRuleDetailsVariables>(
    EditableRuleDetailsDocument,
    baseOptions
  );
}
export function useEditableRuleDetailsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    EditableRuleDetails,
    EditableRuleDetailsVariables
  >
) {
  return ApolloReactHooks.useLazyQuery<EditableRuleDetails, EditableRuleDetailsVariables>(
    EditableRuleDetailsDocument,
    baseOptions
  );
}
export type EditableRuleDetailsHookResult = ReturnType<typeof useEditableRuleDetails>;
export type EditableRuleDetailsLazyQueryHookResult = ReturnType<
  typeof useEditableRuleDetailsLazyQuery
>;
export type EditableRuleDetailsQueryResult = ApolloReactCommon.QueryResult<
  EditableRuleDetails,
  EditableRuleDetailsVariables
>;
