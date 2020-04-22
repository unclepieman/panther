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
import { Box, Button, Grid, Label, Text } from 'pouncejs';
import { formatDatetime } from 'Helpers/utils';
import Panel from 'Components/Panel';
import Linkify from 'Components/Linkify';
import { GlobalModuleDetails } from 'Generated/schema';
import { Link } from 'react-router-dom';
import urls from 'Source/urls';

interface GlobalModuleDetailsInfoProps {
  global?: GlobalModuleDetails;
}

const GlobalModuleDetailsInfo: React.FC<GlobalModuleDetailsInfoProps> = ({ global }) => {
  return (
    <Panel
      size="large"
      title="Global Module Details"
      actions={
        <Box>
          <Button
            size="large"
            variant="default"
            mr={4}
            as={Link}
            to={urls.settings.globalModule.edit(global.id)}
          >
            Edit
          </Button>
        </Box>
      }
    >
      <Grid gridTemplateColumns="repeat(2, 1fr)" gridGap={6}>
        <Box my={1}>
          <Label mb={1} as="div" size="small" color="grey300">
            ID
          </Label>
          <Text size="medium" color="black">
            {global.id}
          </Text>
        </Box>
        <Box my={1}>
          <Label mb={1} as="div" size="small" color="grey300">
            DESCRIPTION
          </Label>
          <Text size="medium" color={global.description ? 'black' : 'grey200'}>
            <Linkify>{global.description || 'No description available'}</Linkify>
          </Text>
        </Box>
        <Box my={1}>
          <Label mb={1} as="div" size="small" color="grey300">
            CREATED
          </Label>
          <Text size="medium" color="black">
            {formatDatetime(global.createdAt)}
          </Text>
        </Box>
        <Box my={1}>
          <Label mb={1} as="div" size="small" color="grey300">
            LAST MODIFIED
          </Label>
          <Text size="medium" color="black">
            {formatDatetime(global.lastModified)}
          </Text>
        </Box>
      </Grid>
    </Panel>
  );
};

export default React.memo(GlobalModuleDetailsInfo);
