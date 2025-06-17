---
page_title: "umbrella_devices Resource - devices"
description: |-
  For the tag ID, remove the tag on the roaming computers in the organization.
After the delete, if the tag is not assigned to any roaming computers in the organization, Umbrella removes the tag from
the organization. You can recreate the tag for your organization.
---

# umbrella_devices (Resource)

For the tag ID, remove the tag on the roaming computers in the organization.
After the delete, if the tag is not assigned to any roaming computers in the organization, Umbrella removes the tag from
the organization. You can recreate the tag for your organization.

## Example Usage


### Basic Usage

Basic usage of the devices resource

```hcl
resource "umbrella_devices" "example" {
  # Add required attributes here
  name = "example-devices"
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
terraform import umbrella_devices.example 12345
```

