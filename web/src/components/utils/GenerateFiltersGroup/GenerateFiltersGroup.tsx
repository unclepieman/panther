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
import { Formik, Field, Form } from 'formik';
import { Box, Button, SimpleGrid } from 'pouncejs';
import mapValues from 'lodash/mapValues';
import map from 'lodash/map';
import FormikTextInput from 'Components/fields/TextInput';
import FormikTextArea from 'Components/fields/TextArea';
import FormikCombobox from 'Components/fields/ComboBox';
import FormikMultiCombobox from 'Components/fields/MultiComboBox';

interface FiltersGroupData<T> {
  /** The component to render for this particular form entry */
  component:
    | typeof FormikCombobox
    | typeof FormikMultiCombobox
    | typeof FormikTextInput
    | typeof FormikTextArea;

  /** The props that should be given to the form entry */
  props: { [key: string]: any };
}

interface GenerateFiltersGroupProps<T> {
  /** The initial values of the filters group */
  initialValues: T;

  /** A callback for when the `apply` button is clicked or - in general - the filters are applied */
  onSubmit: (values: T) => void;

  /** A filter configuration */
  filters: {
    [K in Extract<keyof T, string>]: FiltersGroupData<T>;
  };
}

const getFilterGroupDefaultValue = (
  component: FiltersGroupData<any>['component']
): [] | '' | null => {
  switch (component) {
    case FormikCombobox:
      return null;
    case FormikMultiCombobox:
      return [];
    default:
      return '';
  }
};

function GenerateFiltersGroup<T extends { [key: string]: any }>({
  filters,
  onSubmit,
  initialValues,
}: GenerateFiltersGroupProps<T>): React.ReactElement<GenerateFiltersGroupProps<T>> {
  // These are the default values that each field should have. This is related to the type of the
  // field (a.k.a. `component`) and is the "fallback initial value of the field" (since its actual
  // initial value comes from the URL)
  const defaultValues = React.useMemo(() => {
    return mapValues(filters, filterData => getFilterGroupDefaultValue(filterData.component)) as T;
  }, []);

  // We initialize the values of the form based on the current URL. This only happens during mount
  // time. The value of `initialValues` doesn't get updated as the component updates, since we only
  // need it during form initialization (a.k.a. component mount-time)
  const initialValuesWithDefaults = React.useMemo(() => {
    return mapValues(filters, (value, name) =>
      initialValues[name] === undefined ? defaultValues[name] : initialValues[name]
    ) as T;
  }, []);

  // On a successful submit, the URL params are updated and the page query gets re-fetched, since
  // the page query depends on what the URL is. Essentially, we are using the URL params as a
  // "store" that we observe on the index
  return (
    <Formik<T> initialValues={initialValuesWithDefaults} onSubmit={onSubmit}>
      {({ setValues, submitForm }) => (
        <Form>
          <SimpleGrid columns={4} spacing={6} mb={6}>
            {map(filters, (filterData, filterName) => (
              <Box as="fieldset" key={filterName}>
                <Field as={filterData.component} name={filterName} {...filterData.props} />
              </Box>
            ))}
          </SimpleGrid>
          <SimpleGrid inline columns={2} spacingX={4}>
            <Button type="submit">Apply</Button>
            <Button
              variantColor="red"
              onClick={() => {
                setValues(defaultValues);
                submitForm();
              }}
            >
              Clear all
            </Button>
          </SimpleGrid>
        </Form>
      )}
    </Formik>
  );
}

export default GenerateFiltersGroup;
