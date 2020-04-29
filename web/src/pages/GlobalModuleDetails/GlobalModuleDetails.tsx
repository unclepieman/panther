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
import useRouter from 'Hooks/useRouter';
import { Alert } from 'pouncejs';
import { extractErrorMessage } from 'Helpers/utils';
import ErrorBoundary from 'Components/ErrorBoundary';
import GlobalModuleDetailsInfo from './GlobalModuleDetailsInfo';
import GlobalModuleDetailsPageSkeleton from './Skeleton';
import { useGlobalModuleDetails } from './graphql/globalModuleDetails.generated';

const GlobalModuleDetailsPage = () => {
  const { match } = useRouter<{ id: string }>();

  const { error, data, loading } = useGlobalModuleDetails({
    fetchPolicy: 'cache-and-network',
    variables: {
      globalModuleDetailsInput: {
        globalId: match.params.id,
      },
    },
  });

  if (loading && !data) {
    return <GlobalModuleDetailsPageSkeleton />;
  }

  if (error) {
    return (
      <Alert
        variant="error"
        title="Couldn't load global module"
        description={
          extractErrorMessage(error) ||
          "An unknown error occurred and we couldn't load the global module from the server"
        }
        mb={6}
      />
    );
  }

  return (
    <article>
      <ErrorBoundary>
        <GlobalModuleDetailsInfo global={data.getGlobalPythonModule} />
      </ErrorBoundary>
    </article>
  );
};

export default GlobalModuleDetailsPage;
