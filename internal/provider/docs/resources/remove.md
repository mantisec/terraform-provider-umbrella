---
page_title: "umbrella_remove Resource - remove"
description: |-
  Remove the Secure Web Gateway (SWG) override setting for the devices in the organization.
Once you remove the override setting on a device, Umbrella applies your organization's SWG setting to the device.
---

# umbrella_remove (Resource)

Remove the Secure Web Gateway (SWG) override setting for the devices in the organization.
Once you remove the override setting on a device, Umbrella applies your organization's SWG setting to the device.

## Example Usage


### Basic Usage

Basic usage of the remove resource

```hcl
resource "umbrella_remove" "example" {
  # Add required attributes here
  name = "example-remove"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `status` (String) Deleted SWG override setting on the devices.



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_remove.example 12345
```

