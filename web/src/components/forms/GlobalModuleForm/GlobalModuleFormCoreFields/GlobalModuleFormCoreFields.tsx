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
import { Field, useFormikContext } from 'formik';
import FormikTextInput from 'Components/fields/TextInput';
import { InputElementLabel, Grid, Box, InputElementErrorLabel } from 'pouncejs';
import { capitalize } from 'Helpers/utils';
import FormikTextArea from 'Components/fields/TextArea';
import FormikEditor from 'Components/fields/Editor';
import { GlobalModuleFormValues } from 'Components/forms/GlobalModuleForm';

export const globalModuleEditableFields = ['id', 'body', 'description'] as const;

type FormValues = Required<Pick<GlobalModuleFormValues, typeof globalModuleEditableFields[number]>>;

const GlobalModuleFormCoreFields: React.FC = () => {
  // Read the values from the "parent" form. We expect a formik to be declared in the upper scope
  // since this is a "partial" form. If no Formik context is found this will error out intentionally
  const { errors, touched, initialValues } = useFormikContext<FormValues>();

  return (
    <section>
      <Grid gridTemplateColumns="1fr 1fr" gridRowGap={2} gridColumnGap={9}>
        <Field
          as={FormikTextInput}
          label="* ID"
          placeholder={`The unique ID of the global`}
          name="id"
          disabled={initialValues.id}
          aria-required
        />
        <Field
          as={FormikTextArea}
          label="Description"
          placeholder={`Additional context about this global module`}
          name="description"
        />
      </Grid>
      <Box my={6}>
        <InputElementLabel htmlFor="enabled">{`* ${capitalize(
          'Global'
        )} Function`}</InputElementLabel>
        <Field
          as={FormikEditor}
          placeholder={`# Enter the body of the global here...`}
          name="body"
          width="100%"
          minLines={16}
          mode="python"
          aria-required
        />
        {errors.body && touched.body && (
          <InputElementErrorLabel mt={6}>{errors.body}</InputElementErrorLabel>
        )}
      </Box>
    </section>
  );
};

export default GlobalModuleFormCoreFields;
