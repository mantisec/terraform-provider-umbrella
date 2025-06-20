---
page_title: "umbrella_getdeviceswithtag Data Source - terraform-provider-umbrella"
description: |-
  List the roaming computers that have a tag with the specified tag ID.
---

# umbrella_getdeviceswithtag (Data Source)

List the roaming computers that have a tag with the specified tag ID.

## Example Usage


### Basic Usage

Basic usage of the getdeviceswithtag data source

```terraform
data "umbrella_getdeviceswithtag" "example" {
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
- **`label`** (String) - The display name of the device.
- **`type`** (String) - The device type.
- **`createdAt`** (String) - The date and time (timestamp) that Umbrella added the roaming computer to the organization. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.
- **`modifiedAt`** (String) - The date and time (timestamp) that Umbrella updated the roaming computer in the organization. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.
- **`originId`** (Number) - The unique identifier of the device.
- **`organizationId`** (Number) - The organization ID.



