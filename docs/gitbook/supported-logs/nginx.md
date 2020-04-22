# Nginx

{% hint style="info" %}
Required fields are in **bold**.
{% endhint %}

## Nginx.Access

Access Logs for your Nginx server. We currently support Nginx 'combined' format. Reference: [http://nginx.org/en/docs/http/ngx\_http\_log\_module.html\#log\_format](http://nginx.org/en/docs/http/ngx_http_log_module.html#log_format)

| Column | Type | Description |
| :--- | :--- | :--- |
| `remoteAddr` | `string` | The IP address of the client \(remote host\) which made the request to the server. |
| `remoteUser` | `string` | The userid of the person making the request. Usually empty unless .htaccess has requested authentication. |
| **`time`** | `timestamp` | The time that the request was received \(UTC\). |
| `request` | `string` | The request line from the client. It includes the HTTP method, the resource requested, and the HTTP protocol. |
| `status` | `smallint` | The HTTP status code returned to the client. |
| `bodyBytesSent` | `bigint` | The size of the object returned to the client, measured in bytes. |
| `httpReferer` | `string` | The HTTP referrer if any. |
| `httpUserAgent` | `string` | The agent the user used when making the request. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |

