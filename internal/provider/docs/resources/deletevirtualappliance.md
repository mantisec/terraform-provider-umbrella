---
page_title: "umbrella_deletevirtualappliance Resource - terraform-provider-umbrella"
description: |-
  Delete a virtual appliance in the organization.
---

# umbrella_deletevirtualappliance (Resource)

Delete a virtual appliance in the organization.

## Example Usage


### Basic Usage

Basic usage of the deletevirtualappliance resource

```terraform
resource "umbrella_deletevirtualappliance" "example" {
  name        = "example-deletevirtualappliance"
  description = "Example deletevirtualappliance resource"
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

umbrella_deletevirtualappliance can be imported using the resource ID:

```shell
terraform import umbrella_deletevirtualappliance.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

