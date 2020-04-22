# AWS Application Load Balancer Has Web ACL

| Risk | Remediation Effort |
| :--- | :--- |
| **High** | **Medium** |

This policy validates that each AWS Elastic Load Balancer is protected by the correct AWS WAF Web ACL. This can prevent many attacks before they reach your web servers, including XSS and SQL injection attacks.

This policy requires configuration before it can be enabled.

**Remediation**

To remediate this, assign a WAF Web ACL to the load balancer from the AWS [WAF panel](https://console.aws.amazon.com/wafv2/home?#/webacls/rules/).

| Using the AWS Console |
| :--- |


| 1. Selecting the region that the WAF and load balancer exist in from the `Filter` dropdown |
| :--- |


| 2. Selecting the Web ACL you would like to associate to the load balancer \(one must be created if one does not already exist in the specified region\) |
| :--- |


| 3. Selecting the `Rules` tab |
| :--- |


| 4. Selecting the `Add association` button |
| :--- |


| 5. Selecting the appropriate resource type in the `Resource type` dropdown |
| :--- |


| 6. Selecting the desired load balancer from the `Resource` dropdown |
| :--- |


| 7. Selecting the `Add` button |
| :--- |


* AWS [WAF Web ACL](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl-associating-cloudfront-distribution.html) documentation

