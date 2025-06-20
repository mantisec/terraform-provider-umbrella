---
page_title: "umbrella_createpasswordresets Resource - terraform-provider-umbrella"
description: |-
  Renew the passwords for the user accounts in the child (customer) organization of a provider organization.

For each email address, Umbrella identifies the user's account and renews the password for the user account.
Umbrella does not send a renewal request to the user account's email address.

Once an admin resets the user account password through the `/passwordResets/{customerId}` API endpoint,
a user can sign into their account on the child (customer) organization's managed console
and reset the password for their user account.

The `/passwordResets/{customerId}` API endpoint is only available for parent (provider) organizations
on the Multi-org or provider console.
---

# umbrella_createpasswordresets (Resource)

Renew the passwords for the user accounts in the child (customer) organization of a provider organization.

For each email address, Umbrella identifies the user's account and renews the password for the user account.
Umbrella does not send a renewal request to the user account's email address.

Once an admin resets the user account password through the `/passwordResets/{customerId}` API endpoint,
a user can sign into their account on the child (customer) organization's managed console
and reset the password for their user account.

The `/passwordResets/{customerId}` API endpoint is only available for parent (provider) organizations
on the Multi-org or provider console.

## Example Usage


### Basic Usage

Basic usage of the createpasswordresets resource

```terraform
resource "umbrella_createpasswordresets" "example" {
  name        = "example-createpasswordresets"
  description = "Example createpasswordresets resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional

- **`adminEmails`** (Set of String) - A list of email addresses for the user accounts in a child (customer) organization. Example: `["item1", "item2"]`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier



## Import

umbrella_createpasswordresets can be imported using the resource ID:

```shell
terraform import umbrella_createpasswordresets.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

