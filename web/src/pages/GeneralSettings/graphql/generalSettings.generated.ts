/* eslint-disable import/order, import/no-duplicates */
import * as Types from '../../../../__generated__/schema';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type GetGeneralSettingsVariables = {};

export type GetGeneralSettings = {
  generalSettings: Pick<Types.GeneralSettings, 'displayName' | 'email' | 'errorReportingConsent'>;
};

export type UpdateGeneralSettingsVariables = {
  input: Types.UpdateGeneralSettingsInput;
};

export type UpdateGeneralSettings = {
  updateGeneralSettings: Pick<
    Types.GeneralSettings,
    'displayName' | 'email' | 'errorReportingConsent'
  >;
};

export const GetGeneralSettingsDocument = gql`
  query GetGeneralSettings {
    generalSettings {
      displayName
      email
      errorReportingConsent
    }
  }
`;

/**
 * __useGetGeneralSettings__
 *
 * To run a query within a React component, call `useGetGeneralSettings` and pass it any options that fit your needs.
 * When your component renders, `useGetGeneralSettings` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetGeneralSettings({
 *   variables: {
 *   },
 * });
 */
export function useGetGeneralSettings(
  baseOptions?: ApolloReactHooks.QueryHookOptions<GetGeneralSettings, GetGeneralSettingsVariables>
) {
  return ApolloReactHooks.useQuery<GetGeneralSettings, GetGeneralSettingsVariables>(
    GetGeneralSettingsDocument,
    baseOptions
  );
}
export function useGetGeneralSettingsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    GetGeneralSettings,
    GetGeneralSettingsVariables
  >
) {
  return ApolloReactHooks.useLazyQuery<GetGeneralSettings, GetGeneralSettingsVariables>(
    GetGeneralSettingsDocument,
    baseOptions
  );
}
export type GetGeneralSettingsHookResult = ReturnType<typeof useGetGeneralSettings>;
export type GetGeneralSettingsLazyQueryHookResult = ReturnType<
  typeof useGetGeneralSettingsLazyQuery
>;
export type GetGeneralSettingsQueryResult = ApolloReactCommon.QueryResult<
  GetGeneralSettings,
  GetGeneralSettingsVariables
>;
export const UpdateGeneralSettingsDocument = gql`
  mutation UpdateGeneralSettings($input: UpdateGeneralSettingsInput!) {
    updateGeneralSettings(input: $input) {
      displayName
      email
      errorReportingConsent
    }
  }
`;
export type UpdateGeneralSettingsMutationFn = ApolloReactCommon.MutationFunction<
  UpdateGeneralSettings,
  UpdateGeneralSettingsVariables
>;

/**
 * __useUpdateGeneralSettings__
 *
 * To run a mutation, you first call `useUpdateGeneralSettings` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateGeneralSettings` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateGeneralSettings, { data, loading, error }] = useUpdateGeneralSettings({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateGeneralSettings(
  baseOptions?: ApolloReactHooks.MutationHookOptions<
    UpdateGeneralSettings,
    UpdateGeneralSettingsVariables
  >
) {
  return ApolloReactHooks.useMutation<UpdateGeneralSettings, UpdateGeneralSettingsVariables>(
    UpdateGeneralSettingsDocument,
    baseOptions
  );
}
export type UpdateGeneralSettingsHookResult = ReturnType<typeof useUpdateGeneralSettings>;
export type UpdateGeneralSettingsMutationResult = ApolloReactCommon.MutationResult<
  UpdateGeneralSettings
>;
export type UpdateGeneralSettingsMutationOptions = ApolloReactCommon.BaseMutationOptions<
  UpdateGeneralSettings,
  UpdateGeneralSettingsVariables
>;
