---
page_title: "umbrella_listvirtualappliances Data Source - terraform-provider-umbrella"
description: |-
  List the virtual appliances in the organization.
---

# umbrella_listvirtualappliances (Data Source)

List the virtual appliances in the organization.

## Example Usage


### Basic Usage

Basic usage of the listvirtualappliances data source

```terraform
data "umbrella_listvirtualappliances" "example" {
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
- **`createdAt`** (String) - The date and time (ISO8601 timestamp) when the VA was created.
- **`modifiedAt`** (String) - The date and time (ISO8601 timestamp) when the VA was modified.
- **`siteId`** (Number) - The site ID of the virtual appliance.
- **`isUpgradable`** (Boolean) - Specifies whether you can upgrade the virtual appliance (VA) to the latest VA version.
- **`state`** (String) - The properties for the state of the virtual appliance.
- **`health`** (String) - A description of the health of the virtual appliance.
- **`type`** (String) - The type of the virtual appliance.
- **`settings`** (String) - The properties of the settings on the virtual appliance.
- **`stateUpdatedAt`** (String) - The date and time (ISO8601 timestamp) when the virtual appliance's state was updated.
- **`originId`** (Number) - The origin ID of the virtual appliance.
- **`name`** (String) - The name of the virtual appliance.



