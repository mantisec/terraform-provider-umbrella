---
page_title: "umbrella_listusers Data Source - terraform-provider-umbrella"
description: |-
  List the user accounts in the organization.
---

# umbrella_listusers (Data Source)

List the user accounts in the organization.

## Example Usage


### Basic Usage

Basic usage of the listusers data source

```terraform
data "umbrella_listusers" "example" {
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
- **`password`** (String) - The user's password.
- **`role`** (String) - The user's role.
- **`timezone`** (String) - The user's timezone.
- **`firstname`** (String) - The user's first name.
- **`lastname`** (String) - The user's last name.
- **`email`** (String) - The user's email address.
- **`roleId`** (Number) - The role ID.
- **`lastLoginTime`** (String) - The user's last login date and time (ISO8601 timestamp).
- **`twoFactorEnable`** (Boolean) - Specifies whether two-factor authentication is enabled.



