# Fluentd

{% hint style="info" %}
Required fields are in **bold**.
{% endhint %}

## Fluentd.Syslog3164

Fluentd syslog parser for the RFC3164 format \(ie. BSD-syslog messages\) Reference: [https://docs.fluentd.org/parser/syslog\#rfc3164-log](https://docs.fluentd.org/parser/syslog#rfc3164-log)

| Column | Type | Description |
| :--- | :--- | :--- |
| `pri` | `smallint` | Priority is calculated by \(Facility \* 8 + Severity\). The lower this value, the higher importance of the log message. |
| **`host`** | `string` | Hostname identifies the machine that originally sent the syslog message. |
| **`ident`** | `string` | Appname identifies the device or application that originated the syslog message. |
| **`pid`** | `bigint` | ProcID is often the process ID, but can be any value used to enable log analyzers to detect discontinuities in syslog reporting. |
| **`message`** | `string` | Message contains free-form text that provides information about the event. |
| **`time`** | `timestamp` | Timestamp of the syslog message in UTC. |
| **`tag`** | `string` | Tag of the syslog message |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |

## Fluentd.Syslog5424

Fluentd syslog parser for the RFC5424 format \(ie. BSD-syslog messages\) Reference: [https://docs.fluentd.org/parser/syslog\#rfc5424-log](https://docs.fluentd.org/parser/syslog#rfc5424-log)

| Column | Type | Description |
| :--- | :--- | :--- |
| `pri` | `smallint` | Priority is calculated by \(Facility \* 8 + Severity\). The lower this value, the higher importance of the log message. |
| **`host`** | `string` | Hostname identifies the machine that originally sent the syslog message. |
| **`ident`** | `string` | Appname identifies the device or application that originated the syslog message. |
| **`pid`** | `bigint` | ProcID is often the process ID, but can be any value used to enable log analyzers to detect discontinuities in syslog reporting. |
| **`msgid`** | `string` | MsgID identifies the type of message. For example, a firewall might use the MsgID 'TCPIN' for incoming TCP traffic. |
| **`extradata`** | `string` | ExtraData contains syslog strucured data as string |
| **`message`** | `string` | Message contains free-form text that provides information about the event. |
| **`time`** | `timestamp` | Timestamp of the syslog message in UTC. |
| **`tag`** | `string` | Tag of the syslog message |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |

