---
page_title: "umbrella_deletedestinationlist Resource - terraform-provider-umbrella"
description: |-
  Delete a destination list from your organization.
---

# umbrella_deletedestinationlist (Resource)

Delete a destination list from your organization.

## Example Usage


### Basic Usage

Basic usage of the deletedestinationlist resource

```terraform
resource "umbrella_deletedestinationlist" "example" {
  name      = "example-destination-list"
  access    = "allow"
  is_global = false
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`data`** (Set of String) - 
- **`name`** (String) - The name of the resource
- **`description`** (String) - The description of the resource
- **`enabled`** (Boolean) - Whether the resource is enabled
- **`active`** (Boolean) - Whether the resource is active
- **`status`** (String) - The status of the resource
- **`organization_id`** (Number) - The organization ID
- **`created_at`** (Number) - The date and time when the resource was created
- **`modified_at`** (Number) - The date and time when the resource was modified
- **`updated_at`** (Number) - The date and time when the resource was updated
- **`created_by`** (String) - The user who created the resource
- **`modified_by`** (String) - The user who modified the resource



## Import

umbrella_deletedestinationlist can be imported using the resource ID:

```shell
terraform import umbrella_deletedestinationlist.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

