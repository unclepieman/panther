# AWS Create GuardDuty Detector

## Remediation Id

`AWS.GuardDuty.CreateDetector`

## Description

Remediation that creates a GuardDuty detector if one doesn't exist.

## Resource Parameters

| Name | Description |
| :--- | :--- |
| `AccountId` | The AWS Account Id |
| `Region` | The AWS region |

## Additional Parameters

| Name | Description |
| :--- | :--- |


<table>
  <thead>
    <tr>
      <th style="text-align:left"><code>FindingPublishingFrequency</code>
      </th>
      <th style="text-align:left">
        <p>A enum value that specifies how frequently finding updates will be published.</p>
        <p>Possible values:</p>
        <ul>
          <li>FIFTEEN_MINUTES</li>
          <li>ONE_HOUR</li>
          <li>SIX_HOURS</li>
        </ul>
      </th>
    </tr>
  </thead>
  <tbody></tbody>
</table>* [https://docs.aws.amazon.com/cli/latest/reference/guardduty/create-detector.html](https://docs.aws.amazon.com/cli/latest/reference/guardduty/create-detector.html)
* [https://aws.amazon.com/guardduty/](https://aws.amazon.com/guardduty/)

