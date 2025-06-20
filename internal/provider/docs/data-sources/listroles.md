---
page_title: "umbrella_listroles Data Source - terraform-provider-umbrella"
description: |-
  List the roles in the organization.
---

# umbrella_listroles (Data Source)

List the roles in the organization.

## Example Usage


### Basic Usage

Basic usage of the listroles data source

```terraform
data "umbrella_listroles" "example" {
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
- **`roleId`** (Number) - The ID of the role.
- **`label`** (String) - The label for the role.
- **`organizationId`** (Number) - The organization ID.



