---
page_title: "umbrella_listprovidercustomers Data Source - terraform-provider-umbrella"
description: |-
  List the customers for the provider.
---

# umbrella_listprovidercustomers (Data Source)

List the customers for the provider.

## Example Usage


### Basic Usage

Basic usage of the listprovidercustomers data source

```terraform
data "umbrella_listprovidercustomers" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`isTrial`** (Boolean) - Specify whether the customer is a trial with an MSSP with SPLA (MSLA) license.
- **`addonRbi`** (String) - Specify whether remote browser isolation (RBI) is added to the subscription. Remote Browser Isolation Add-On is applicable for SIG Essentials or SIG Advantage only: `0` = No selection `1` = RBI Isolate Risky `2` = RBI Isolate Web Applications `3` = RBI Isolate All
- **`customerId`** (Number) - The ID for the customer.
- **`dealId`** (String) - The deal ID.
- **`ccwDealOwnerEmails`** (Set of String) - The list emails for the CCW deal owner.
- **`addonCdfwL7`** (Boolean) - Specify whether if cloud delivered firewall (CDFW) is added to the subscription. The L7 Cloud Delivered Firewall Solution is applicable only for SIG E.
- **`streetAddress`** (String) - The street address for the customer.
- **`adminEmails`** (Set of String) - The list of the administrator email addresses.
- **`modifiedAt`** (String) - The time when the customer information was last modified. The timestamp is specified in the ISO 8601 format.
- **`city`** (String) - The name of the city where the customer's organization is located.
- **`state`** (String) - The name of the customer's state.
- **`countryCode`** (String) - The country code of the customer's organization.
- **`createdAt`** (String) - The time when the customer information was created. The timestamp is specified in the ISO 8601 format.
- **`customerName`** (String) - The name of the customer's organization.
- **`packageName`** (String) - The name of the Umbrella package. To create or update a customer with either the SIG Essentials or SIG Advantage package, you must have a license for the selected package. | Package Id | Package Name | |:----------:|----------------------------------| | `99` | Umbrella Professional | | `101` | Umbrella Platform | | `107` | Umbrella Insights | | `171` | Cisco Umbrella for Wireless LAN | | `202` | Cisco Umbrella for EDU | | `246` | Umbrella DNS Security Essentials | | `248` | Umbrella DNS Security Advantage | | `250` | Umbrella SIG Essentials | | `252` | Umbrella SIG Advantage | | `312` | Umbrella Not for Resale (NFR) MSP DNS Advantage | Note: The Umbrella NFR MSP DNS Advantage package (`312`) is only available in the Umbrella Secure MSP console.
- **`streetAddress2`** (String) - The second street address for the customer.
- **`packageId`** (Number) - The ID of the Umbrella package. To create or update a customer with either the SIG Essentials or SIG Advantage package, you must have a license for the selected package. | Package Id | Package Name | |:----------:|----------------------------------| | `99` | Umbrella Professional | | `101` | Umbrella Platform | | `107` | Umbrella Insights | | `171` | Cisco Umbrella for Wireless LAN | | `202` | Cisco Umbrella for EDU | | `246` | Umbrella DNS Security Essentials | | `248` | Umbrella DNS Security Advantage | | `250` | Umbrella SIG Essentials | | `252` | Umbrella SIG Advantage | | `312` | Umbrella Not for Resale (NFR) MSP DNS Advantage | Note: The Umbrella NFR MSP DNS Advantage package (`312`) is only available in the Umbrella Secure MSP console.
- **`addonDlp`** (Boolean) - Specify whether data loss prevention (DLP) is added to the subscription. The Data Loss Prevention Add-On is applicable only for SIG E.
- **`licenseType`** (String) - The type of license for the customer's organization.
- **`seats`** (Number) - The number of users.
- **`zipCode`** (String) - The zip code of the customer's organization.



