---
page_title: "umbrella_getallnetworkdevices Data Source - terraform-provider-umbrella"
description: |-
  List the network devices.
---

# umbrella_getallnetworkdevices (Data Source)

List the network devices.

## Example Usage


### Basic Usage

Basic usage of the getallnetworkdevices data source

```terraform
data "umbrella_getallnetworkdevices" "example" {
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
- **`name`** (String) - The name of the device.
- **`serialNumber`** (String) - The serial number of the device.
- **`createdAt`** (String) - The time when the device was created. Specify an ISO 8601-formatted timestamp.
- **`organizationId`** (Number) - The ID of the organization.
- **`originId`** (Number) - The unique global identifier for this traffic source (origin). Use the origin ID to manage the device. The origin ID is not used by the device.
- **`deviceId`** (String) - The unique identifier for the specific network device. Insert the identifier into the EDNS packets.
- **`deviceKey`** (String) - A descriptive unique identifier for the device. Not used by the device.



