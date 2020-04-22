# AWS

{% hint style="info" %}
Required fields are in **bold**.
{% endhint %}

## AWS.ALB

Application Load Balancer logs Layer 7 network logs for your application load balancer. Reference: [https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| `type` | `string` | The type of request or connection. |
| **`timestamp`** | `timestamp` | The time when the load balancer generated a response to the client \(UTC\). For WebSockets, this is the time when the connection is closed. |
| `elb` | `string` | The resource ID of the load balancer. If you are parsing access log entries, note that resources IDs can contain forward slashes \(/\). |
| `clientIp` | `string` | The IP address of the requesting client. |
| `clientPort` | `bigint` | The port of the requesting client. |
| `targetIp` | `string` | The IP address of the target that processed this request. |
| `targetPort` | `bigint` | The port of the target that processed this request. |
| `requestProcessingTime` | `double` | The total time elapsed \(in seconds, with millisecond precision\) from the time the load balancer received the request until the time it sent it to a target. This value is set to -1 if the load balancer can't dispatch the request to a target. This can ha... |
| `targetProcessingTime` | `double` | The total time elapsed \(in seconds, with millisecond precision\) from the time the load balancer sent the request to a target until the target started to send the response headers. This value is set to -1 if the load balancer can't dispatch the request ... |
| `responseProcessingTime` | `double` | The total time elapsed \(in seconds, with millisecond precision\) from the time the load balancer received the response header from the target until it started to send the response to the client. This includes both the queuing time at the load balancer a... |
| `elbStatusCode` | `bigint` | The status code of the response from the load balancer. |
| `targetStatusCode` | `bigint` | The status code of the response from the target. This value is recorded only if a connection was established to the target and the target sent a response. |
| `receivedBytes` | `bigint` | The size of the request, in bytes, received from the client \(requester\). For HTTP requests, this includes the headers. For WebSockets, this is the total number of bytes received from the client on the connection. |
| `sentBytes` | `bigint` | The size of the response, in bytes, sent to the client \(requester\). For HTTP requests, this includes the headers. For WebSockets, this is the total number of bytes sent to the client on the connection. |
| `requestHttpMethod` | `string` | The HTTP method parsed from the request. |
| `requestUrl` | `string` | The HTTP URL parsed from the request. |
| `requestHttpVersion` | `string` | The HTTP version parsed from the request. |
| `userAgent` | `string` | A User-Agent string that identifies the client that originated the request. The string consists of one or more product identifiers, product\[/version\]. If the string is longer than 8 KB, it is truncated. |
| `sslCipher` | `string` | \[HTTPS listener\] The SSL cipher. This value is set to NULL if the listener is not an HTTPS listener. |
| `sslProtocol` | `string` | \[HTTPS listener\] The SSL protocol. This value is set to NULL if the listener is not an HTTPS listener. |
| `targetGroupArn` | `string` | The Amazon Resource Name \(ARN\) of the target group. |
| `traceId` | `string` | The contents of the X-Amzn-Trace-Id header. |
| `domainName` | `string` | \[HTTPS listener\] The SNI domain provided by the client during the TLS handshake. This value is set to NULL if the client doesn't support SNI or the domain doesn't match a certificate and the default certificate is presented to the client. |
| `chosenCertArn` | `string` | \[HTTPS listener\] The ARN of the certificate presented to the client. This value is set to session-reused if the session is reused. This value is set to NULL if the listener is not an HTTPS listener. |
| `matchedRulePriority` | `bigint` | The priority value of the rule that matched the request. If a rule matched, this is a value from 1 to 50,000. If no rule matched and the default action was taken, this value is set to 0. If an error occurs during rules evaluation, it is set to -1. For ... |
| `requestCreationTime` | `timestamp` | The time when the load balancer received the request from the client. |
| `actionsExecuted` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | The actions taken when processing the request. This value is a comma-separated list that can include the values described in Actions Taken. If no action was taken, such as for a malformed request, this value is set to NULL. |
| `redirectUrl` | `string` | The URL of the redirect target for the location header of the HTTP response. If no redirect actions were taken, this value is set to NULL. |
| `errorReason` | `string` | The error reason code. If the request failed, this is one of the error codes described in Error Reason Codes. If the actions taken do not include an authenticate action or the target is not a Lambda function, this value is set to NULL. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |
| `p_any_aws_account_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws account ids associated with the row |
| `p_any_aws_instance_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws instance ids associated with the row |
| `p_any_aws_arns` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws arns associated with the row |
| `p_any_aws_tags` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws tags associated with the row |

## AWS.AuroraMySQLAudit

AuroraMySQLAudit is an RDS Aurora audit log which contains context around database calls. Reference: [https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Auditing.html](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Auditing.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| `timestamp` | `timestamp` | The timestamp for the logged event with microsecond precision \(UTC\). |
| `serverHost` | `string` | The name of the instance that the event is logged for. |
| `username` | `string` | The connected user name of the user. |
| `host` | `string` | The host that the user connected from. |
| `connectionId` | `bigint` | The connection ID number for the logged operation. |
| `queryId` | `bigint` | The query ID number, which can be used for finding the relational table events and related queries. For TABLE events, multiple lines are added. |
| `operation` | `string` | The recorded action type. Possible values are: CONNECT, QUERY, READ, WRITE, CREATE, ALTER, RENAME, and DROP. |
| `database` | `string` | The active database, as set by the USE command. |
| `object` | `string` | For QUERY events, this value indicates the executed query. For TABLE events, it indicates the table name. |
| `retCode` | `bigint` | The return code of the logged operation. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |
| `p_any_aws_account_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws account ids associated with the row |
| `p_any_aws_instance_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws instance ids associated with the row |
| `p_any_aws_arns` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws arns associated with the row |
| `p_any_aws_tags` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws tags associated with the row |

## AWS.CloudTrail

AWSCloudTrail represents the content of a CloudTrail S3 object. Log format & samples can be seen here: [https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference.html](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| `additionalEventData` | `string` | Additional data about the event that was not part of the request or response. |
| `apiVersion` | `string` | Identifies the API version associated with the AwsApiCall eventType value. |
| **`awsRegion`** | `string` | The AWS region that the request was made to, such as us-east-2. |
| `errorCode` | `string` | The AWS service error if the request returns an error. |
| `errorMessage` | `string` | If the request returns an error, the description of the error. This message includes messages for authorization failures. CloudTrail captures the message logged by the service in its exception handling. |
| **`eventId`** | `string` | GUID generated by CloudTrail to uniquely identify each event. You can use this value to identify a single event. For example, you can use the ID as a primary key to retrieve log data from a searchable database. |
| **`eventName`** | `string` | The requested action, which is one of the actions in the API for that service. |
| **`eventSource`** | `string` | The service that the request was made to. This name is typically a short form of the service name without spaces plus .amazonaws.com. |
| **`eventTime`** | `timestamp` | The date and time the request was made, in coordinated universal time \(UTC\). |
| **`eventType`** | `string` | Identifies the type of event that generated the event record. This can be the one of the following values: AwsApiCall, AwsServiceEvent, AwsConsoleSignIn |
| **`eventVersion`** | `string` | The version of the log event format. |
| `managementEvent` | `boolean` | A Boolean value that identifies whether the event is a management event. managementEvent is shown in an event record if eventVersion is 1.06 or higher, and the event type is one of the following: AwsApiCall, AwsConsoleAction, AwsConsoleSignIn, AwsServ... |
| `readOnly` | `boolean` | Identifies whether this operation is a read-only operation. |
| `recipientAccountId` | `string` | Represents the account ID that received this event. The recipientAccountID may be different from the CloudTrail userIdentity Element accountId. This can occur in cross-account resource access. |
| `requestId` | `string` | The value that identifies the request. The service being called generates this value. |
| `requestParameters` | `string` | The parameters, if any, that were sent with the request. These parameters are documented in the API reference documentation for the appropriate AWS service. |
| `resources` | `"CloudTrailResources":{   "arn": {     "type": "string"   },   "accountId": {     "type": "string"   },   "type": {     "type": "string"   } }  {   "items": {          "$ref": "CloudTrailResources"   },   "type": "array" }`                                 | A list of resources accessed in the event. |
| `responseElements` | `string` | The response element for actions that make changes \(create, update, or delete actions\). If an action does not change state \(for example, a request to get or list objects\), this element is omitted. These actions are documented in the API reference docum... |
| `serviceEventDetails` | `string` | Identifies the service event, including what triggered the event and the result. |
| `sharedEventId` | `string` | GUID generated by CloudTrail to uniquely identify CloudTrail events from the same AWS action that is sent to different AWS accounts. |
| **`sourceIpAddress`** | `string` | The IP address that the request was made from. For actions that originate from the service console, the address reported is for the underlying customer resource, not the console web server. For services in AWS, only the DNS name is displayed. |
| `userAgent` | `string` | The agent through which the request was made, such as the AWS Management Console, an AWS service, the AWS SDKs or the AWS CLI. |
| **`userIdentity`** | `"CloudTrailSessionContext":{   "attributes": {          "$ref": "CloudTrailSessionContextAttributes"   },   "sessionIssuer": {          "$ref": "CloudTrailSessionContextSessionIssuer"   },   "webIdFederationData": {          "$ref": "CloudTrailSessionContextWebIDFederationData"   } }  "CloudTrailSessionContextAttributes":{   "mfaAuthenticated": {     "type": "string"   },   "creationDate": {     "type": "string"   } }  "CloudTrailSessionContextSessionIssuer":{   "type": {     "type": "string"   },   "principalId": {     "type": "string"   },   "arn": {     "type": "string"   },   "accountId": {     "type": "string"   },   "userName": {     "type": "string"   } }  "CloudTrailSessionContextWebIDFederationData":{   "federatedProvider": {     "type": "string"   },   "attributes": {     "items": {       "type": "integer"     },     "type": "array"   } }  {   "type": {     "type": "string"   },   "principalId": {     "type": "string"   },   "arn": {     "type": "string"   },   "accountId": {     "type": "string"   },   "accessKeyId": {     "type": "string"   },   "userName": {     "type": "string"   },   "sessionContext": {          "$ref": "CloudTrailSessionContext"   },   "invokedBy": {     "type": "string"   },   "identityProvider": {     "type": "string"   } }`                                 | Information about the user that made a request. |
| `vpcEndpointId` | `string` | Identifies the VPC endpoint in which requests were made from a VPC to another AWS service, such as Amazon S3. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |
| `p_any_aws_account_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws account ids associated with the row |
| `p_any_aws_instance_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws instance ids associated with the row |
| `p_any_aws_arns` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws arns associated with the row |
| `p_any_aws_tags` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws tags associated with the row |

## AWS.GuardDuty

Amazon GuardDuty is a threat detection service that continuously monitors for malicious activity and unauthorized behavior inside AWS Accounts. See also GuardDuty Finding Format : [https://docs.aws.amazon.com/guardduty/latest/ug/guardduty\_finding-format.html](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_finding-format.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| **`schemaVersion`** | `string` | The schema format version of this record. |
| `accountId` | `string` | The ID of the AWS account in which the activity took place that prompted GuardDuty to generate this finding. |
| **`region`** | `string` | The AWS region in which the finding was generated. |
| **`partition`** | `string` | The AWS partition in which the finding was generated. |
| **`id`** | `string` | A unique identifier for the finding. |
| **`arn`** | `string` | A unique identifier formatted as an ARN for the finding. |
| **`type`** | `string` | A concise yet readable description of the potential security issue. |
| **`resource`** | `string` | The AWS resource against which the activity took place that prompted GuardDuty to generate this finding. |
| **`severity`** | `float` | The value of the severity can fall anywhere within the 0.1 to 8.9 range. |
| **`createdAt`** | `timestamp` | The initial creation time of the finding \(UTC\). |
| **`updatedAt`** | `timestamp` | The last update time of the finding \(UTC\). |
| **`title`** | `string` | A short description of the finding. |
| **`description`** | `string` | A long description of the finding. |
| **`service`** | `"RFC3339": {   "type": "timestamp" }  {   "additionalInfo": {     "items": {       "type": "integer"     },     "type": "array"   },   "action": {     "items": {       "type": "integer"     },     "type": "array"   },   "serviceName": {     "type": "string"   },   "detectorId": {     "type": "string"   },   "resourceRole": {     "type": "string"   },   "eventFirstSeen": {          "$ref": "RFC3339"   },   "eventLastSeen": {     "$ref": "RFC3339"   },   "archived": {     "type": "boolean"   },   "count": {     "type": "integer"   } }`                                 | Additional information about the affected service. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |
| `p_any_aws_account_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws account ids associated with the row |
| `p_any_aws_instance_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws instance ids associated with the row |
| `p_any_aws_arns` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws arns associated with the row |
| `p_any_aws_tags` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws tags associated with the row |

## AWS.S3ServerAccess

S3ServerAccess is an AWS S3 Access Log. Log format & samples can be seen here: [https://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html](https://docs.aws.amazon.com/AmazonS3/latest/dev/LogFormat.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| **`bucketowner`** | `string` | The canonical user ID of the owner of the source bucket. The canonical user ID is another form of the AWS account ID. |
| `bucket` | `string` | The name of the bucket that the request was processed against. If the system receives a malformed request and cannot determine the bucket, the request will not appear in any server access log. |
| `time` | `timestamp` | The time at which the request was received \(UTC\). |
| `remoteip` | `string` | The apparent internet address of the requester. Intermediate proxies and firewalls might obscure the actual address of the machine making the request. |
| `requester` | `string` | The canonical user ID of the requester, or NULL for unauthenticated requests. If the requester was an IAM user, this field returns the requester's IAM user name along with the AWS root account that the IAM user belongs to. This identifier is the same o... |
| `requestid` | `string` | A string generated by Amazon S3 to uniquely identify each request. |
| `operation` | `string` | The operation listed here is declared as SOAP.operation, REST.HTTP\_method.resource\_type, WEBSITE.HTTP\_method.resource\_type, or BATCH.DELETE.OBJECT. |
| `key` | `string` | The key part of the request, URL encoded, or NULL if the operation does not take a key parameter. |
| `requesturi` | `string` | The Request-URI part of the HTTP request message. |
| `httpstatus` | `bigint` | The numeric HTTP status code of the response. |
| `errorcode` | `string` | The Amazon S3 Error Code, or NULL if no error occurred. |
| `bytessent` | `bigint` | The number of response bytes sent, excluding HTTP protocol overhead, or NULL if zero. |
| `objectsize` | `bigint` | The total size of the object in question. |
| `totaltime` | `bigint` | The number of milliseconds the request was in flight from the server's perspective. This value is measured from the time your request is received to the time that the last byte of the response is sent. Measurements made from the client's perspective mi... |
| `turnaroundtime` | `bigint` | The number of milliseconds that Amazon S3 spent processing your request. This value is measured from the time the last byte of your request was received until the time the first byte of the response was sent. |
| `referrer` | `string` | The value of the HTTP Referer header, if present. HTTP user-agents \(for example, browsers\) typically set this header to the URL of the linking or embedding page when making a request. |
| `useragent` | `string` | The value of the HTTP User-Agent header. |
| `versionid` | `string` | The version ID in the request, or NULL if the operation does not take a versionId parameter. |
| `hostid` | `string` | The x-amz-id-2 or Amazon S3 extended request ID. |
| `signatureversion` | `string` | The signature version, SigV2 or SigV4, that was used to authenticate the request or NULL for unauthenticated requests. |
| `ciphersuite` | `string` | The Secure Sockets Layer \(SSL\) cipher that was negotiated for HTTPS request or NULL for HTTP. |
| `authenticationtype` | `string` | The type of request authentication used, AuthHeader for authentication headers, QueryString for query string \(pre-signed URL\) or NULL for unauthenticated requests. |
| `hostheader` | `string` | The endpoint used to connect to Amazon S3. |
| `tlsVersion` | `string` | The Transport Layer Security \(TLS\) version negotiated by the client. The value is one of following: TLSv1, TLSv1.1, TLSv1.2; or NULL if TLS wasn't used. |
| `additionalFields` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | The remaining columns in the record as an array. |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |
| `p_any_aws_account_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws account ids associated with the row |
| `p_any_aws_instance_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws instance ids associated with the row |
| `p_any_aws_arns` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws arns associated with the row |
| `p_any_aws_tags` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws tags associated with the row |

## AWS.VPCFlow

VPCFlow is a VPC NetFlow log, which is a layer 3 representation of network traffic in EC2. Log format & samples can be seen here: [https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-records-examples.html](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-records-examples.html)

| Column | Type | Description |
| :--- | :--- | :--- |
| `version` | `bigint` | The VPC Flow Logs version. If you use the default format, the version is 2. If you specify a custom format, the version is 3. |
| `account` | `string` | The AWS account ID for the flow log. |
| `interfaceId` | `string` | The ID of the network interface for which the traffic is recorded. |
| `srcAddr` | `string` | The source address for incoming traffic, or the IPv4 or IPv6 address of the network interface for outgoing traffic on the network interface. The IPv4 address of the network interface is always its private IPv4 address. |
| `dstAddr` | `string` | The destination address for outgoing traffic, or the IPv4 or IPv6 address of the network interface for incoming traffic on the network interface. The IPv4 address of the network interface is always its private IPv4 address. |
| `srcPort` | `bigint` | The source port of the traffic. |
| `dstPort` | `bigint` | The destination port of the traffic. |
| `protocol` | `bigint` | The IANA protocol number of the traffic. |
| `packets` | `bigint` | The number of packets transferred during the flow. |
| `bytes` | `bigint` | The number of bytes transferred during the flow. |
| **`start`** | `timestamp` | The time of the start of the flow \(UTC\). |
| **`end`** | `timestamp` | The time of the end of the flow \(UTC\). |
| `action` | `string` | The action that is associated with the traffic. ACCEPT: The recorded traffic was permitted by the security groups or network ACLs. REJECT: The recorded traffic was not permitted by the security groups or network ACLs. |
| `status` | `string` | The logging status of the flow log. OK: Data is logging normally to the chosen destinations. NODATA: There was no network traffic to or from the network interface during the capture window. SKIPDATA: Some flow log records were skipped during the captur... |
| `vpcId` | `string` | The ID of the VPC that contains the network interface for which the traffic is recorded. |
| `subNetId` | `string` | The ID of the subnet that contains the network interface for which the traffic is recorded. |
| `instanceId` | `string` | The ID of the instance that's associated with network interface for which the traffic is recorded, if the instance is owned by you. Returns a '-' symbol for a requester-managed network interface; for example, the network interface for a NAT gateway. |
| `tcpFlags` | `bigint` | The bitmask value for the following TCP flags: SYN: 2, SYN-ACK: 18, FIN: 1, RST: 4. ACK is reported only when it's accompanied with SYN. TCP flags can be OR-ed during the aggregation interval. For short connections, the flags might be set on the same l... |
| `trafficType` | `string` | The type of traffic: IPv4, IPv6, or EFA. |
| `pktSrcAddr` | `string` | The packet-level \(original\) source IP address of the traffic. Use this field with the srcaddr field to distinguish between the IP address of an intermediate layer through which traffic flows, and the original source IP address of the traffic. For examp... |
| `pktDstAddr` | `string` | The packet-level \(original\) destination IP address for the traffic. Use this field with the dstaddr field to distinguish between the IP address of an intermediate layer through which traffic flows, and the final destination IP address of the traffic. F... |
| **`p_log_type`** | `string` | Panther added field with type of log |
| **`p_row_id`** | `string` | Panther added field with unique id \(within table\) |
| **`p_event_time`** | `timestamp` | Panther added standardize event time \(UTC\) |
| **`p_parse_time`** | `timestamp` | Panther added standardize log parse time \(UTC\) |
| `p_any_ip_addresses` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of ip addresses associated with the row |
| `p_any_domain_names` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of domain names associated with the row |
| `p_any_sha1_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of SHA1 hashes associated with the row |
| `p_any_md5_hashes` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of MD5 hashes associated with the row |
| `p_any_aws_account_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws account ids associated with the row |
| `p_any_aws_instance_ids` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws instance ids associated with the row |
| `p_any_aws_arns` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws arns associated with the row |
| `p_any_aws_tags` | `{   "items": {     "type": "string"   },   "type": "array" }`                                 | Panther added field with collection of aws tags associated with the row |

