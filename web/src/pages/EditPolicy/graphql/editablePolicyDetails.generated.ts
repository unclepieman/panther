/* eslint-disable import/order, import/no-duplicates */
import * as Types from '../../../../__generated__/schema';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type EditablePolicyDetailsVariables = {
  input: Types.GetPolicyInput;
};

export type EditablePolicyDetails = {
  policy: Types.Maybe<
    Pick<
      Types.PolicyDetails,
      | 'autoRemediationId'
      | 'autoRemediationParameters'
      | 'description'
      | 'displayName'
      | 'enabled'
      | 'suppressions'
      | 'id'
      | 'reference'
      | 'resourceTypes'
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

export const EditablePolicyDetailsDocument = gql`
  query EditablePolicyDetails($input: GetPolicyInput!) {
    policy(input: $input) {
      autoRemediationId
      autoRemediationParameters
      description
      displayName
      enabled
      suppressions
      id
      reference
      resourceTypes
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
 * __useEditablePolicyDetails__
 *
 * To run a query within a React component, call `useEditablePolicyDetails` and pass it any options that fit your needs.
 * When your component renders, `useEditablePolicyDetails` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useEditablePolicyDetails({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useEditablePolicyDetails(
  baseOptions?: ApolloReactHooks.QueryHookOptions<
    EditablePolicyDetails,
    EditablePolicyDetailsVariables
  >
) {
  return ApolloReactHooks.useQuery<EditablePolicyDetails, EditablePolicyDetailsVariables>(
    EditablePolicyDetailsDocument,
    baseOptions
  );
}
export function useEditablePolicyDetailsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    EditablePolicyDetails,
    EditablePolicyDetailsVariables
  >
) {
  return ApolloReactHooks.useLazyQuery<EditablePolicyDetails, EditablePolicyDetailsVariables>(
    EditablePolicyDetailsDocument,
    baseOptions
  );
}
export type EditablePolicyDetailsHookResult = ReturnType<typeof useEditablePolicyDetails>;
export type EditablePolicyDetailsLazyQueryHookResult = ReturnType<
  typeof useEditablePolicyDetailsLazyQuery
>;
export type EditablePolicyDetailsQueryResult = ApolloReactCommon.QueryResult<
  EditablePolicyDetails,
  EditablePolicyDetailsVariables
>;
