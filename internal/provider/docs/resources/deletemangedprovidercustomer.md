---
page_title: "umbrella_deletemangedprovidercustomer Resource - terraform-provider-umbrella"
description: |-
  Delete an customer for the managed provider.
---

# umbrella_deletemangedprovidercustomer (Resource)

Delete an customer for the managed provider.

## Example Usage


### Basic Usage

Basic usage of the deletemangedprovidercustomer resource

```terraform
resource "umbrella_deletemangedprovidercustomer" "example" {
  name        = "example-deletemangedprovidercustomer"
  description = "Example deletemangedprovidercustomer resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
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

umbrella_deletemangedprovidercustomer can be imported using the resource ID:

```shell
terraform import umbrella_deletemangedprovidercustomer.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

