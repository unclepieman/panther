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

import React from 'react';
import { GlobalModuleDetails } from 'Generated/schema';
import * as Yup from 'yup';
import { Box, Button, Flex } from 'pouncejs';
import ErrorBoundary from 'Components/ErrorBoundary';
import { Formik } from 'formik';
import SubmitButton from 'Components/buttons/SubmitButton/SubmitButton';
import useRouter from 'Hooks/useRouter';
import GlobalModuleFormCoreFields, {
  globalModuleEditableFields,
} from './GlobalModuleFormCoreFields';

// The validation checks that Formik will run
const validationSchema = Yup.object().shape({
  id: Yup.string().required(),
  body: Yup.string().required(),
  description: Yup.string().required(),
});

interface BaseGlobalModuleFormProps<GlobalModuleFormValues> {
  /** The initial values of the form */
  initialValues: GlobalModuleFormValues;

  /** callback for the submission of the form */
  onSubmit: (values: GlobalModuleFormValues) => void;

  /** The validation schema that the form will have */
  validationSchema: Yup.ObjectSchema<Yup.Shape<object, Partial<GlobalModuleFormValues>>>;
}

export type GlobalModuleFormValues = Pick<
  GlobalModuleDetails,
  typeof globalModuleEditableFields[number]
>;

export type GlobalModuleFormProps = Pick<
  BaseGlobalModuleFormProps<GlobalModuleFormValues>,
  'initialValues' | 'onSubmit'
>;

const GlobalModuleForm: React.FC<GlobalModuleFormProps> = ({ initialValues, onSubmit }) => {
  const { history } = useRouter();

  return (
    <Box as="article">
      <ErrorBoundary>
        <Formik<GlobalModuleFormValues>
          initialValues={initialValues}
          onSubmit={onSubmit}
          enableReinitialize
          validationSchema={validationSchema}
        >
          {({ handleSubmit, isSubmitting, isValid, dirty }) => {
            return (
              <form onSubmit={handleSubmit}>
                <GlobalModuleFormCoreFields />
                <Flex
                  borderTop="1px solid"
                  borderColor="grey100"
                  pt={6}
                  mt={10}
                  justifyContent="flex-end"
                >
                  <Flex>
                    <Button variant="default" size="large" onClick={history.goBack} mr={4}>
                      Cancel
                    </Button>
                    <SubmitButton
                      submitting={isSubmitting}
                      disabled={!dirty || !isValid || isSubmitting}
                    >
                      {initialValues.id ? 'Update' : 'Create'}
                    </SubmitButton>
                  </Flex>
                </Flex>
              </form>
            );
          }}
        </Formik>
      </ErrorBoundary>
    </Box>
  );
};

export default GlobalModuleForm;
