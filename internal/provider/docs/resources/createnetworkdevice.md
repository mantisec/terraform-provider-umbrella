---
page_title: "umbrella_createnetworkdevice Resource - terraform-provider-umbrella"
description: |-
  Create a network device.
---

# umbrella_createnetworkdevice (Resource)

Create a network device.

## Example Usage


### Basic Usage

Basic usage of the createnetworkdevice resource

```terraform
resource "umbrella_createnetworkdevice" "example" {
  name        = "example-network"
  description = "Example network configuration"
  enabled     = true
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`model`** (String) - The model name of the device. Must be unique to your organization. Example: `"example"`
- **`macAddress`** (String) - The MAC address of the device. The unique MAC address may include up to 12 characters and must not contain hyphens or colons. Example: `"example"`
- **`name`** (String) - The name of the device. The name is a sequence of characters with a length from 1 through 50. The name must be unique in your organization. Example: `"example-name"`
- **`serialNumber`** (String) - The serial number of the device. Example: `"example"`


### Optional

- **`tag`** (String) - A text tag that describes the device or this origin, which is assigned to the device. Provide a tag that is unique to your organization. Example: `"example"`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`organizationId`** (Number) - The ID of the organization.
- **`originId`** (Number) - The unique global identifier for this traffic source (origin). Use the origin ID to manage the device. The origin ID is not used by the device.
- **`deviceId`** (String) - The unique identifier for the specific network device. Insert the identifier into the EDNS packets.
- **`deviceKey`** (String) - A descriptive unique identifier for the device. Not used by the device.
- **`createdAt`** (String) - The time when the device was created. Specify an ISO 8601-formatted timestamp.



## Import

umbrella_createnetworkdevice can be imported using the resource ID:

```shell
terraform import umbrella_createnetworkdevice.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

