---
page_title: "umbrella_passwordresets Resource - passwordresets"
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

# umbrella_passwordresets (Resource)

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

Basic usage of the passwordresets resource

```hcl
resource "umbrella_passwordresets" "example" {
  # Add required attributes here
  name = "example-passwordresets"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_passwordresets.example 12345
```

