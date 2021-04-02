// Package nginxlogs provides parsers for NGINX server logs
package nginxlogs

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

import (
	"github.com/panther-labs/panther/internal/log_analysis/log_processor/logtypes"
	"github.com/panther-labs/panther/internal/log_analysis/log_processor/parsers"
)

const (
	TypeAccess = `Nginx.Access`
)

func LogTypes() logtypes.Group {
	return logTypes
}

var logTypes = logtypes.Must("Nginx",
	logtypes.Config{
		Name:         TypeAccess,
		Description:  `Access Logs for your Nginx server. We currently support Nginx 'combined' format.`,
		ReferenceURL: `http://nginx.org/en/docs/http/ngx_http_log_module.html#log_format`,
		Schema:       Access{},
		NewParser:    parsers.AdapterFactory(&AccessParser{}),
	},
)
