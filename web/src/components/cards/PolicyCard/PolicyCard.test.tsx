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
import { buildPolicy, render } from 'test-utils';
import { SeverityEnum } from 'Generated/schema';
import { SelectProvider } from 'Components/utils/SelectContext';
import urls from 'Source/urls';
import PolicyCard from './index';

describe('PolicyCard', () => {
  it('displays the correct data in the card', async () => {
    const policy = buildPolicy();

    const { getByText } = render(<PolicyCard policy={policy} />);

    expect(getByText(policy.displayName)).toBeInTheDocument();
    expect(getByText(SeverityEnum.Medium)).toBeInTheDocument();
    expect(getByText(policy.enabled ? 'ENABLED' : 'DISABLED')).toBeInTheDocument();

    policy.resourceTypes.forEach(resourceType => {
      expect(getByText(resourceType)).toBeInTheDocument();
    });
  });

  it('should have valid links', async () => {
    const policy = buildPolicy();

    const { getByAriaLabel } = render(<PolicyCard policy={policy} />);
    expect(getByAriaLabel('Link to Policy')).toHaveAttribute(
      'href',
      urls.compliance.policies.details(policy.id)
    );
  });

  it('renders a checkbox when selection is enabled', () => {
    const { getByAriaLabel } = render(
      <SelectProvider>
        <PolicyCard policy={buildPolicy()} selectionEnabled />
      </SelectProvider>
    );

    expect(getByAriaLabel(`select item`)).toBeInTheDocument();
  });
});
