---
page_title: "umbrella_updatecontact Resource - terraform-provider-umbrella"
description: |-
  Update a contact for service providers console.
---

# umbrella_updatecontact (Resource)

Update a contact for service providers console.

## Example Usage


### Basic Usage

Basic usage of the updatecontact resource

```terraform
resource "umbrella_updatecontact" "example" {
  name        = "example-updatecontact"
  description = "Example updatecontact resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`createdAt`** (Number) - The time when the contact information was created. The time is specified in milliseconds.
- **`primaryEmailAddress`** (String) - The email address of primary contact of distributor organization.
- **`distributorVisibility`** (Boolean) - Specify whether distributors primary contact has visibility into trials.
- **`orgName`** (String) - The organization name of the distributor.
- **`primaryContact`** (String) - Specify if the primary contact.
- **`city`** (String) - The city where the contact is located.
- **`state`** (String) - The state where the contact is located.
- **`modifiedAt`** (Number) - The time when the contact information was last modified. The time is specified in milliseconds.
- **`firstName`** (String) - The first name of the contact.
- **`phoneNumber`** (String) - The phone number for the contact.
- **`phoneNumber2`** (String) - The second phone number for the contact.
- **`streetAddress2`** (String) - A secondary street address for the contact.
- **`zipCode`** (String) - The US zip code where the contact is located.
- **`faxNumber`** (String) - The fax number for the contact.
- **`contactId`** (Number) - The contact ID.
- **`organizationId`** (Number) - The organization ID.
- **`streetAddress`** (String) - The street address for the contact.
- **`emailAddress`** (String) - The email address for the contact.
- **`contactType`** (String) - The type of contact.
- **`lastName`** (String) - The last name of the contact.
- **`countryCode`** (String) - The country code for the contact.



## Import

umbrella_updatecontact can be imported using the resource ID:

```shell
terraform import umbrella_updatecontact.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

