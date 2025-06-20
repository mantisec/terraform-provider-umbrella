---
page_title: "umbrella_updatenetworkdevice Resource - terraform-provider-umbrella"
description: |-
  Update a network device.
---

# umbrella_updatenetworkdevice (Resource)

Update a network device.

## Example Usage


### Basic Usage

Basic usage of the updatenetworkdevice resource

```terraform
resource "umbrella_updatenetworkdevice" "example" {
  name        = "example-network"
  description = "Example network configuration"
  enabled     = true
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`originId`** (Number) - The unique global identifier for this traffic source (origin). Use the origin ID to manage the device. The origin ID is not used by the device.
- **`deviceId`** (String) - The unique identifier for the specific network device. Insert the identifier into the EDNS packets.
- **`deviceKey`** (String) - A descriptive unique identifier for the device. Not used by the device.
- **`name`** (String) - The name of the device.
- **`serialNumber`** (String) - The serial number of the device.
- **`createdAt`** (String) - The time when the device was created. Specify an ISO 8601-formatted timestamp.
- **`organizationId`** (Number) - The ID of the organization.



## Import

umbrella_updatenetworkdevice can be imported using the resource ID:

```shell
terraform import umbrella_updatenetworkdevice.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

