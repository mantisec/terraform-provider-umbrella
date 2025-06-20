---
page_title: "umbrella_getcontacts Data Source - terraform-provider-umbrella"
description: |-
  List the contacts for the service providers console.
---

# umbrella_getcontacts (Data Source)

List the contacts for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the getcontacts data source

```terraform
data "umbrella_getcontacts" "example" {
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
- **`phoneNumber2`** (String) - The second phone number for the contact.
- **`modifiedAt`** (Number) - The time when the contact information was last modified. The time is specified in milliseconds.
- **`firstName`** (String) - The first name of the contact.
- **`phoneNumber`** (String) - The phone number for the contact.
- **`streetAddress`** (String) - The street address for the contact.
- **`streetAddress2`** (String) - A secondary street address for the contact.
- **`zipCode`** (String) - The US zip code where the contact is located.
- **`faxNumber`** (String) - The fax number for the contact.
- **`contactId`** (Number) - The contact ID.
- **`organizationId`** (Number) - The organization ID.
- **`countryCode`** (String) - The country code for the contact.
- **`emailAddress`** (String) - The email address for the contact.
- **`contactType`** (String) - The type of contact.
- **`lastName`** (String) - The last name of the contact.
- **`state`** (String) - The state where the contact is located.
- **`createdAt`** (Number) - The time when the contact information was created. The time is specified in milliseconds.
- **`primaryEmailAddress`** (String) - The email address of primary contact of distributor organization.
- **`distributorVisibility`** (Boolean) - Specify whether distributors primary contact has visibility into trials.
- **`orgName`** (String) - The organization name of the distributor.
- **`primaryContact`** (String) - Specify if the primary contact.
- **`city`** (String) - The city where the contact is located.



