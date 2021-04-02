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
import { Badge, Box, Table } from 'pouncejs';
import dayjs from 'dayjs';
import { getUserDisplayName } from 'Helpers/utils';
import { ListUsers } from '../graphql/listUsers.generated';
import ListUsersTableRowOptions from './ListUsersTableRowOptions';

type ListUsersTableProps = Pick<ListUsers, 'users'>;

const ListUsersTable: React.FC<ListUsersTableProps> = ({ users }) => {
  return (
    <Table>
      <Table.Head>
        <Table.Row>
          <Table.HeaderCell>Name</Table.HeaderCell>
          <Table.HeaderCell>Email</Table.HeaderCell>
          <Table.HeaderCell>Role</Table.HeaderCell>
          <Table.HeaderCell>Invited At</Table.HeaderCell>
          <Table.HeaderCell align="center">Status</Table.HeaderCell>
          <Table.HeaderCell />
        </Table.Row>
      </Table.Head>
      <Table.Body>
        {users.map(user => (
          <Table.Row key={user.id}>
            <Table.Cell>{getUserDisplayName(user)}</Table.Cell>
            <Table.Cell>{user.email}</Table.Cell>
            <Table.Cell>Admin</Table.Cell>
            <Table.Cell>
              {dayjs(user.createdAt * 1000).format('MM/DD/YYYY, HH:mm G[M]TZZ')}
            </Table.Cell>
            <Table.Cell align="center">
              <Box my={-1} display="inline-block">
                <Badge color="blue-300">{user.status.replace(/_/g, ' ')}</Badge>
              </Box>
            </Table.Cell>
            <Table.Cell>
              <Box my={-2}>
                <ListUsersTableRowOptions user={user} />
              </Box>
            </Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table>
  );
};

export default React.memo(ListUsersTable);
