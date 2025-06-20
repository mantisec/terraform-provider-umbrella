---
page_title: "umbrella_createuser Resource - terraform-provider-umbrella"
description: |-
  Create an Umbrella user account with a designated role.

Once a user account is deleted, you can recreate the account through the `POST` operation.
When you recreate the user account, only set the `email` and `roleId` fields in the Request Body.

For example: {
  "email": "DEVWKS-22@mailinator.com",
  "roleId": 1
}

If you provide all fields for the user account in the Request Body, Umbrella returns an `HTTP/400` (Bad Request) with the
message: `Email already in use`.
---

# umbrella_createuser (Resource)

Create an Umbrella user account with a designated role.

Once a user account is deleted, you can recreate the account through the `POST` operation.
When you recreate the user account, only set the `email` and `roleId` fields in the Request Body.

For example: {
  "email": "DEVWKS-22@mailinator.com",
  "roleId": 1
}

If you provide all fields for the user account in the Request Body, Umbrella returns an `HTTP/400` (Bad Request) with the
message: `Email already in use`.

## Example Usage


### Basic Usage

Basic usage of the createuser resource

```terraform
resource "umbrella_createuser" "example" {
  name        = "example-user"
  description = "Example user account"
  enabled     = true
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`lastname`** (String) - The user's last name. Example: `"example-name"`
- **`email`** (String) - The user's email address. Example: `"user@example.com"`
- **`password`** (String) - The user's password. Example: `"example"`
- **`roleId`** (Number) - The role ID. Example: `123`
- **`timezone`** (String) - The user's timezone. Example: `"example"`
- **`firstname`** (String) - The user's first name. Example: `"example-name"`


### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`lastLoginTime`** (String) - The user's last login date and time (ISO8601 timestamp).
- **`twoFactorEnable`** (Boolean) - Specifies whether two-factor authentication is enabled.
- **`role`** (String) - The user's role.



## Import

umbrella_createuser can be imported using the resource ID:

```shell
terraform import umbrella_createuser.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

