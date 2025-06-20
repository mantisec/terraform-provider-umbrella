---
page_title: "umbrella_updatevirtualappliance Resource - terraform-provider-umbrella"
description: |-
  Update a virtual appliance in the organization.
---

# umbrella_updatevirtualappliance (Resource)

Update a virtual appliance in the organization.

## Example Usage


### Basic Usage

Basic usage of the updatevirtualappliance resource

```terraform
resource "umbrella_updatevirtualappliance" "example" {
  name        = "example-updatevirtualappliance"
  description = "Example updatevirtualappliance resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`isUpgradable`** (Boolean) - Specifies whether you can upgrade the virtual appliance (VA) to the latest VA version.
- **`createdAt`** (String) - The date and time (ISO8601 timestamp) when the VA was created.
- **`modifiedAt`** (String) - The date and time (ISO8601 timestamp) when the VA was modified.
- **`siteId`** (Number) - The site ID of the virtual appliance.
- **`name`** (String) - The name of the virtual appliance.
- **`state`** (String) - The properties for the state of the virtual appliance.
- **`health`** (String) - A description of the health of the virtual appliance.
- **`type`** (String) - The type of the virtual appliance.
- **`settings`** (String) - The properties of the settings on the virtual appliance.
- **`stateUpdatedAt`** (String) - The date and time (ISO8601 timestamp) when the virtual appliance's state was updated.
- **`originId`** (Number) - The origin ID of the virtual appliance.



## Import

umbrella_updatevirtualappliance can be imported using the resource ID:

```shell
terraform import umbrella_updatevirtualappliance.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

