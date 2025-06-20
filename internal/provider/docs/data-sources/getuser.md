---
page_title: "umbrella_getuser Data Source - terraform-provider-umbrella"
description: |-
  Get a user account.
---

# umbrella_getuser (Data Source)

Get a user account.

## Example Usage


### Basic Usage

Basic usage of the getuser data source

```terraform
data "umbrella_getuser" "example" {
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
- **`firstname`** (String) - The user's first name.
- **`lastname`** (String) - The user's last name.
- **`email`** (String) - The user's email address.
- **`roleId`** (Number) - The role ID.
- **`lastLoginTime`** (String) - The user's last login date and time (ISO8601 timestamp).
- **`twoFactorEnable`** (Boolean) - Specifies whether two-factor authentication is enabled.
- **`password`** (String) - The user's password.
- **`role`** (String) - The user's role.
- **`timezone`** (String) - The user's timezone.



