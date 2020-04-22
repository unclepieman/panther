# CloudFormation Stack

## Resource Type

`AWS.CloudFormation.Stack`

## Resource ID Format

For CloudFormation Stacks, the resource ID is the ARN.

`arn:aws:cloudformation:ap-northeast-2:123456789012:stack/example-stack/11111111`

## Background

A stack is a collection of AWS resources that you can manage as code within a template.

## Fields

| Field | Type | Description |
| :--- | :--- | :--- |


| `Capabilities` | `List` | Certain capabilities required in order for AWS CloudFormation to create the stack. |
| :--- | :--- | :--- |


| `ChangeSetId` | `String` | The ID of the change set |
| :--- | :--- | :--- |


| `DeletionTime` | `Time` | The time the stack was deleted. |
| :--- | :--- | :--- |


| `Description` | `String` | A user-defined description associated with the stack. |
| :--- | :--- | :--- |


<table>
  <thead>
    <tr>
      <th style="text-align:left"><code>DisableRollback</code>
      </th>
      <th style="text-align:left"><code>Bool</code>
      </th>
      <th style="text-align:left">
        <p>Boolean to enable or disable rollback on stack creation failures: <code>true</code>disables
          rollback,</p>
        <p><code>false</code> enables rollback.</p>
      </th>
    </tr>
  </thead>
  <tbody></tbody>
</table>| `DriftInformation` | `Map` | [Information](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackDriftInformation.html) on whether a stack's actual configuration differs from its expected configuration |
| :--- | :--- | :--- |


| `EnableTerminationProtection` | `Bool` | Whether termination protection is enabled for the stack. |
| :--- | :--- | :--- |


| `LastUpdatedTime` | `Time` | The time the stack was last updated. This field will only be returned if the stack has been updated at least once. |
| :--- | :--- | :--- |


| `NotificationARNs` | `List` | SNS topic ARNs to which stack related events are published. |
| :--- | :--- | :--- |


| `Outputs` | `List` | A list of [output](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Output.html) structures. |
| :--- | :--- | :--- |


| `RoleARN` | `String` | The associated IAM service role. |
| :--- | :--- | :--- |


| `Drifts` | `List` | Details on the drifted resources. |
| :--- | :--- | :--- |


```javascript
{
    "AccountId": "123456789012",
    "Arn": "arn:aws:cloudformation:ap-northeast-2:123456789012:stack/example-stack",
    "Capabilities": null,
    "ChangeSetId": null,
    "DeletionTime": null,
    "Description": "This is an example stack",
    "DisableRollback": false,
    "DriftInformation": {
        "LastCheckTimestamp": "2019-01-01T00:00:00.00Z",
        "StackDriftStatus": "IN_SYNC"
    },
    "Drifts": null,
    "EnableTerminationProtection": null,
    "Id": "arn:aws:cloudformation:ap-northeast-2:123456789012:stack/example-stack",
    "LastUpdatedTime": null,
    "Name": "example-stack",
    "NotificationARNs": [
        "arn:aws:sns:ap-northeast-2:123456789012:example-topic"
    ],
    "Outputs": null,
    "Parameters": [
        {
            "ParameterKey": "Parameter1",
            "ParameterValue": "Value1",
            "ResolvedValue": null,
            "UsePreviousValue": null
        },
        {
            "ParameterKey": "Parameter2",
            "ParameterValue": "Value2",
            "ResolvedValue": null,
            "UsePreviousValue": null
        }
    ],
    "ParentId": null,
    "Region": "ap-northeast-2",
    "ResourceId": "arn:aws:cloudformation:ap-northeast-2:123456789012:stack/example-stack",
    "ResourceType": "AWS.CloudFormation.Stack",
    "RoleARN": null,
    "RollbackConfiguration": {
        "MonitoringTimeInMinutes": null,
        "RollbackTriggers": null
    },
    "RootId": null,
    "StackId": "arn:aws:cloudformation:ap-northeast-2:123456789012:stack/example-stack",
    "StackStatus": "CREATE_COMPLETE",
    "StackStatusReason": null,
    "Tags": null,
    "TimeCreated": "2019-01-01T00:00:00.000Z",
    "TimeoutInMinutes": null
}
```

