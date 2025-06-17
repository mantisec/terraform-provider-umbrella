---
page_title: "umbrella_set Resource - set"
description: |-
  Add a list of origin ID and the Secure Web Gateway setting for devices in the organization.
The SWG device setting overrides the organization-level SWG setting.

**Note:** Before you can add an SWG override setting to a device, you must register the device as a roaming computer with Umbrella.
Umbrella applies the SWG override setting to a device using the device's origin ID.
You can list the roaming computers in your organization by sending a request to the Umbrella Roaming Computers API.
---

# umbrella_set (Resource)

Add a list of origin ID and the Secure Web Gateway setting for devices in the organization.
The SWG device setting overrides the organization-level SWG setting.

**Note:** Before you can add an SWG override setting to a device, you must register the device as a roaming computer with Umbrella.
Umbrella applies the SWG override setting to a device using the device's origin ID.
You can list the roaming computers in your organization by sending a request to the Umbrella Roaming Computers API.

## Example Usage


### Basic Usage

Basic usage of the set resource

```hcl
resource "umbrella_set" "example" {
  # Add required attributes here
  name = "example-set"
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
terraform import umbrella_set.example 12345
```

