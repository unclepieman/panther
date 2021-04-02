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
// TODO: uncomment when event latency data are fixed (PR #2509, Ticket #2492)
import { Box, Card /* TabList, TabPanel, TabPanels, Tabs */ } from 'pouncejs';
// import { BorderedTab, BorderTabDivider } from 'Components/BorderedTab';
import EventsByLogType from 'Pages/LogAnalysisOverview/EventsByLogType/EventsByLogType';
import { LongSeriesData /* FloatSeriesData */ } from 'Generated/schema';
import Panel from 'Components/Panel';
// import EventsByLatency from '../EventsByLatency';

interface LogTypeChartsProps {
  eventsProcessed: LongSeriesData;
  // eventsLatency: FloatSeriesData;
}

const LogTypeCharts: React.FC<LogTypeChartsProps> = ({ eventsProcessed /* eventsLatency */ }) => {
  return (
    <Card as="section">
      <Panel title="Events by Log Type">
        <Box height={289} py={5} pl={4} backgroundColor="navyblue-500">
          <EventsByLogType events={eventsProcessed} />
        </Box>
      </Panel>
      {/*
      // TODO: uncomment when event latency data are fixed (PR #2509, Ticket #2492)
      <Tabs>
        <Box position="relative" pl={2} pr={4}>
          <TabList>
            <BorderedTab>Events by Log Type</BorderedTab>
            <BorderedTab>Data Latency by Log Type</BorderedTab>
          </TabList>
          <BorderTabDivider />
        </Box>
        <Box p={6}>
          <TabPanels>
            <TabPanel lazy>
              <Box height={289} py={5} pl={4} backgroundColor="navyblue-500">
                <EventsByLogType events={eventsProcessed} />
              </Box>
            </TabPanel>
            <TabPanel lazy>
              <Box height={289} py={5} pl={4} backgroundColor="navyblue-500">
                <EventsByLatency events={eventsLatency} />
              </Box>
            </TabPanel>
          </TabPanels>
        </Box>
      </Tabs> */}
    </Card>
  );
};

export default LogTypeCharts;
