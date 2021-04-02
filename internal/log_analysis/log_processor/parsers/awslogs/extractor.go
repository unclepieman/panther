package awslogs

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
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/tidwall/gjson"
)

// extracts useful AWS features that can be detected generically (w/context)
type AWSExtractor struct {
	pl *AWSPantherLog
}

func NewAWSExtractor(pl *AWSPantherLog) *AWSExtractor {
	return &AWSExtractor{pl: pl}
}

func (e *AWSExtractor) Extract(key, value gjson.Result) {
	// NOTE: add tests as you add new extractions!
	// NOTE: be very careful returning early, keep sure following code does not need to execute

	// value based matching
	if strings.HasPrefix(value.Str, "arn:") {
		/* arns may contain an embedded account id as well as interesting resources
		   See: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
		   Formats:
		    arn:partition:service:region:account-id:resource-id
		    arn:partition:service:region:account-id:resource-type/resource-id
		    arn:partition:service:region:account-id:resource-type:resource-id
		*/
		parsedARN, err := arn.Parse(value.Str)
		if err == nil {
			e.pl.AppendAnyAWSARNs(value.Str)
			e.pl.AppendAnyAWSAccountIds(parsedARN.AccountID)
			// instanceId: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-policy-structure.html#EC2_ARN_Format
			if strings.HasPrefix(parsedARN.Resource, "instance/") {
				slashIndex := strings.LastIndex(parsedARN.Resource, "/")
				if slashIndex < len(parsedARN.Resource)-2 { // not if ends in "/"
					instanceID := parsedARN.Resource[slashIndex+1:]
					if strings.HasPrefix(instanceID, "i-") {
						e.pl.AppendAnyAWSInstanceIds(instanceID)
					}
				}
			}
		}
		return
	}

	// key based matching (not exact)
	if key.Str == "instanceId" || strings.HasSuffix(key.Str, "InstanceId") {
		if strings.HasPrefix(value.Str, "i-") {
			e.pl.AppendAnyAWSInstanceIds(value.Str)
		}
		return
	}

	if key.Str == "accountId" || strings.HasSuffix(key.Str, "AccountId") {
		e.pl.AppendAnyAWSAccountIds(value.Str)
		return
	}

	// exact key based matching
	switch key.Str {
	case "tags": // found in many objects that use ASW tags
		if value.IsArray() {
			value.ForEach(func(tagListKey, tagListValue gjson.Result) bool {
				tagKey := tagListValue.Get("key")
				tagValue := tagListValue.Get("value")
				if tagKey.Exists() && tagValue.Exists() {
					e.pl.AppendAnyAWSTags(tagKey.Str + ":" + tagValue.Str)
				}
				return true
			})
		}

	case "ipv6Addresses": // found in instanceDetails in CloudTrail and GuardDuty (perhaps others)
		if value.IsArray() {
			value.ForEach(func(v6ListKey, v6ListValue gjson.Result) bool {
				e.pl.AppendAnyIPAddress(v6ListValue.Str)
				return true
			})
		}

	case
		"publicIp",         // found in instanceDetails in CloudTrail and GuardDuty (perhaps others)
		"privateIpAddress", // found in instanceDetails in CloudTrail and GuardDuty (perhaps others)
		"ipAddressV4":      // found in GuardDuty findings
		e.pl.AppendAnyIPAddress(value.Str)

	case
		"publicDnsName",  // found in instanceDetails in CloudTrail and GuardDuty (perhaps others)
		"privateDnsName", // found in instanceDetails in CloudTrail and GuardDuty (perhaps others)
		"domain":         // found in GuardDuty findings
		e.pl.AppendAnyDomainNames(value.Str)
	}
}
