# AWS Update Account Password Policy

## Remediation Id

`AWS.IAM.UpdateAccountPasswordPolicy`

## Description

Remediation that sets an account's password policy.

## Resource Parameters

| Name | Description |
| :--- | :--- |
| `AccountId` | The AWS Account Id |

## Additional Parameters

| Name | Description |
| :--- | :--- |


| `MinimumPasswordLength` | The minimum number of characters allowed in an IAM user password |
| :--- | :--- |


<table>
  <thead>
    <tr>
      <th style="text-align:left"><code>RequireSymbols</code>
      </th>
      <th style="text-align:left">
        <p>Boolean that specifies whether IAM user passwords must contain at least
          one of the following non-alphanumeric characters:</p>
        <p><code>! @ # $ % ^ * ( ) _ + - = [ ] { } | &apos;</code>
        </p>
      </th>
    </tr>
  </thead>
  <tbody></tbody>
</table>| `RequireNumbers` | Boolean that specifies whether IAM user passwords must contain at least one numeric character \(0 to 9\) |
| :--- | :--- |


| `RequireUppercaseCharacters` | Boolean that specifies whether IAM user passwords must contain at least one uppercase character from the ISO basic Latin alphabet \(A to Z\) |
| :--- | :--- |


| `RequireLowercaseCharacters` | Boolean that specifies whether IAM user passwords must contain at least one lowercase character from the ISO basic Latin alphabet \(a to z\) |
| :--- | :--- |


| `AllowUsersToChangePassword` | Boolean that specifies if IAM users in the account can change their own passwords |
| :--- | :--- |


| `MaxPasswordAge` | The number of days that an IAM user password is valid |
| :--- | :--- |


| `PasswordReusePrevention` | Specifies the number of previous passwords that IAM users are prevented from reusing |
| :--- | :--- |


* [https://docs.aws.amazon.com/cli/latest/reference/iam/update-account-password-policy.html](https://docs.aws.amazon.com/cli/latest/reference/iam/update-account-password-policy.html)

