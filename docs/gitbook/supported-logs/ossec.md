# OSSEC

{% hint style="info" %}
Required fields are in **bold**.
{% endhint %}

## OSSEC.EventInfo

OSSEC EventInfo alert parser. Currently only JSON output is supported. Reference: [https://www.ossec.net/docs/docs/formats/alerts.html](https://www.ossec.net/docs/docs/formats/alerts.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| **`id`** | `string` | Unique id of the event. |
| **`rule`** | `{   "comment": {     "type": "string"   },   "group": {     "type": "string"   },   "level": {     "type": "integer"   },   "sidid": {     "type": "integer"   },   "CIS": {     "items": {       "type": "string"     },     "type": "array"   },   "cve": {     "type": "string"   },   "firedtimes": {     "type": "integer"   },   "frequency": {     "type": "integer"   },   "groups": {     "items": {       "type": "string"     },     "type": "array"   },   "info": {     "type": "string"   },   "PCI_DSS": {     "items": {       "type": "string"     },     "type": "array"   } }`                                 | Information about the rule that created the event. |
| **`TimeStamp`** | `timestamp` | Timestamp in UTC. |
| **`location`** | `string` | Source of the event \(filename, command, etc\). |
| **`hostname`** | `string` | Hostname of the host that created the event. |
| **`full_log`** | `string` | The full captured log of the event. |
| `action` | `string` | The event action \(drop, deny, accept, etc\). |
| `agentip` | `string` | The IP address of an agent extracted from the hostname. |
| `agent_name` | `string` | The name of an agent extracted from the hostname. |
| `command` | `string` | The command extracted by the decoder. |
| `data` | `string` | Additional data extracted by the decoder. For example a filename. |
| `decoder` | `string` | The name of the decoder used to parse the logs. |
| `decoder_desc` | `{   "accumulate": {     "type": "integer"   },   "fts": {     "type": "integer"   },   "ftscomment": {     "type": "string"   },   "name": {     "type": "string"   },   "parent": {     "type": "string"   } }`                                 | Information about the decoder used to parse the logs. |
| `decoder_parent` | `string` | In the case of a nested decoder, the name of it's parent. |
| `dstgeoip` | `string` | GeoIP location information about the destination IP address. |
| `dstip` | `string` | The destination IP address. |
| `dstport` | `string` | The destination port. |
| `dstuser` | `string` | The destination \(target\) username. |
| `logfile` | `string` | The source log file that was decoded to generate the event. |
| `previous_output` | `string` | The full captured log of the previous event. |
| `program_name` | `string` | The executable name extracted from the log by the decoder used to match a rule. |
| `protocol` | `string` | The protocol \(ip, tcp, udp, etc\) extracted by the decoder. |
| `srcgeoip` | `string` | GeoIP location information about the source IP address. |
| `srcip` | `string` | The source IP address. |
| `srcport` | `string` | The source port. |
| `srcuser` | `string` | The source username. |
| `status` | `string` | Event status \(success, failure, etc\). |
| `SyscheckFile` | `{   "gowner_after": {     "type": "string"   },   "gowner_before": {     "type": "string"   },   "md5_after": {     "type": "string"   },   "md5_before": {     "type": "string"   },   "owner_after": {     "type": "string"   },   "owner_before": {     "type": "string"   },   "path": {     "type": "string"   },   "perm_after": {     "type": "integer"   },   "perm_before": {     "type": "integer"   },   "sha1_after": {     "type": "string"   },   "sha1_before": {     "type": "string"   } }`                                 | Information about a file integrity check. |
| `systemname` | `string` | The system name extracted by the decoder. |
| `url` | `string` | URL of the event. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |

