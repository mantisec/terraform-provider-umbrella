---
page_title: "umbrella_createcontact Resource - terraform-provider-umbrella"
description: |-
  Create a contact for the service providers console.
---

# umbrella_createcontact (Resource)

Create a contact for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the createcontact resource

```terraform
resource "umbrella_createcontact" "example" {
  name        = "example-createcontact"
  description = "Example createcontact resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`contactType`** (String) - The type of contact. Example: `"example"`
- **`emailAddress`** (String) - The email address for the contact. Example: `"user@example.com"`


### Optional

- **`zipCode`** (String) - The US zip code where the contact is located. Example: `"example"`
- **`countryCode`** (String) - The country code for the contact. Example: `"example"`
- **`settings`** (String) - The contact information of the distributor. Example: `"example"`
- **`primaryContact`** (String) - Specify whether the contact is the primary contact. Example: `"example"`
- **`firstName`** (String) - The first name of the contact. Example: `"example-name"`
- **`lastName`** (String) - The last name of the contact. Example: `"example-name"`
- **`streetAddress2`** (String) - A secondary street address for the contact. Example: `"example"`
- **`state`** (String) - The state where the contact is located. Example: `"example"`
- **`city`** (String) - The city where the contact is located. Example: `"example"`
- **`phoneNumber`** (String) - The phone number for the contact. Example: `"example"`
- **`streetAddress`** (String) - The street address for the contact. Example: `"example"`
- **`phoneNumber2`** (String) - The second phone number for the contact. Example: `"example"`
- **`faxNumber`** (String) - The fax number for the contact. Example: `"example"`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`createdAt`** (Number) - The time when the contact information was created. The time is specified in milliseconds.
- **`primaryEmailAddress`** (String) - The email address of primary contact of distributor organization.
- **`distributorVisibility`** (Boolean) - Specify whether distributors primary contact has visibility into trials.
- **`orgName`** (String) - The organization name of the distributor.
- **`modifiedAt`** (Number) - The time when the contact information was last modified. The time is specified in milliseconds.
- **`organizationId`** (Number) - The organization ID.
- **`contactId`** (Number) - The contact ID.



## Import

umbrella_createcontact can be imported using the resource ID:

```shell
terraform import umbrella_createcontact.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

