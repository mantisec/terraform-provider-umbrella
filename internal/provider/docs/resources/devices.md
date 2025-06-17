---
page_title: "umbrella_devices Resource - devices"
description: |-
  For the tag ID, add the tag to the roaming computers.
---

# umbrella_devices (Resource)

For the tag ID, add the tag to the roaming computers.

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

