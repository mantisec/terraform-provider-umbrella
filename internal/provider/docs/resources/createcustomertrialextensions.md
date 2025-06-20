---
page_title: "umbrella_createcustomertrialextensions Resource - terraform-provider-umbrella"
description: |-
  Create the extension for the customer's trial.
---

# umbrella_createcustomertrialextensions (Resource)

Create the extension for the customer's trial.

## Example Usage


### Basic Usage

Basic usage of the createcustomertrialextensions resource

```terraform
resource "umbrella_createcustomertrialextensions" "example" {
  name        = "example-createcustomertrialextensions"
  description = "Example createcustomertrialextensions resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional

- **`trialExtensionDays`** (Number) - The number of days to extend the trial. Example: `123`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`modifiedAt`** (Number) - The time when the logo was last modified. Specify the time in milliseconds.
- **`packageInternalName`** (String) - The internal name for the package.
- **`strength`** (String) - The number of features consumed by a customer trial.
- **`dealId`** (String) - The deal ID.
- **`ppovLifecycle`** (String) - The email details about lifecycle events from a customer trial.
- **`hasDistributorVisibility`** (Boolean) - Specify whether the distributor has visibility into the trial.
- **`trialPeriod`** (String) - The period of the trial, including the start and end times.
- **`organizationTypeId`** (Number) - The type ID of the customer's organization.
- **`packageId`** (Number) - The ID of the Umbrella package. To create or update a customer with either the SIG Essentials or SIG Advantage package, you must have a license for the selected package. | Package Id | Package Name | |:----------:|----------------------------------| | `99` | Umbrella Professional | | `101` | Umbrella Platform | | `107` | Umbrella Insights | | `171` | Cisco Umbrella for Wireless LAN | | `202` | Cisco Umbrella for EDU | | `246` | Umbrella DNS Security Essentials | | `248` | Umbrella DNS Security Advantage | | `250` | Umbrella SIG Essentials | | `252` | Umbrella SIG Advantage | | `312` | Umbrella Not for Resale (NFR) MSP DNS Advantage | Note: The Umbrella NFR MSP DNS Advantage package (`312`) is only available in the Umbrella Secure MSP console.
- **`accountManagerEmails`** (Set of String) - The emails of the account managers.
- **`originId`** (Number) - The origin ID created for the customer.
- **`adminEmails`** (Set of String) - The list of the administrator email addresses.
- **`streetAddress`** (String) - The street address for the customer.
- **`isOnboardingWizardCompleted`** (Boolean) - Specify whether the customer has logged into Umbrella.
- **`trialId`** (String) - The MD5 value of the organization ID.
- **`organizationId`** (Number) - The ID of the organization.
- **`endsAt`** (String) - The end date of the subscription.
- **`city`** (String) - The name of the city where the customer's organization is located.
- **`zipCode`** (String) - The zip code of the customer's organization.
- **`accessRequestId`** (Number) - The ID of the access request to enable access to the child organization.
- **`accessRequestState`** (String) - The state of the access request.
- **`trialExtensionCount`** (Number) - The number of extensions that are applied to the trial.
- **`organizationName`** (String) - The name of the customer's organization.
- **`subscriptionId`** (Number) - The ID of the subscription.
- **`packageName`** (String) - The name of the Umbrella package. To create or update a customer with either the SIG Essentials or SIG Advantage package, you must have a license for the selected package. | Package Id | Package Name | |:----------:|----------------------------------| | `99` | Umbrella Professional | | `101` | Umbrella Platform | | `107` | Umbrella Insights | | `171` | Cisco Umbrella for Wireless LAN | | `202` | Cisco Umbrella for EDU | | `246` | Umbrella DNS Security Essentials | | `248` | Umbrella DNS Security Advantage | | `250` | Umbrella SIG Essentials | | `252` | Umbrella SIG Advantage | | `312` | Umbrella Not for Resale (NFR) MSP DNS Advantage | Note: The Umbrella NFR MSP DNS Advantage package (`312`) is only available in the Umbrella Secure MSP console.
- **`streetAddress2`** (String) - The second street address for the customer.
- **`trialExtendedDays`** (Number) - The number of days to extend the trial.
- **`createdAt`** (Number) - The time when the logo was created. Specify the time in milliseconds.
- **`countryCode`** (String) - The country code of the customer's organization.
- **`users`** (Number) - The number of users in the subscription.
- **`startsAt`** (String) - The start date of the subscription.
- **`state`** (String) - The name of the customer's state.



## Import

umbrella_createcustomertrialextensions can be imported using the resource ID:

```shell
terraform import umbrella_createcustomertrialextensions.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

