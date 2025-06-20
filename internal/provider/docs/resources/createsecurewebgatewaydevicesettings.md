---
page_title: "umbrella_createsecurewebgatewaydevicesettings Resource - terraform-provider-umbrella"
description: |-
  Add a list of origin ID and the Secure Web Gateway setting for devices in the organization.
The SWG device setting overrides the organization-level SWG setting.

**Note:** Before you can add an SWG override setting to a device, you must register the device as a roaming computer with Umbrella.
Umbrella applies the SWG override setting to a device using the device's origin ID.
You can list the roaming computers in your organization by sending a request to the Umbrella Roaming Computers API.
---

# umbrella_createsecurewebgatewaydevicesettings (Resource)

Add a list of origin ID and the Secure Web Gateway setting for devices in the organization.
The SWG device setting overrides the organization-level SWG setting.

**Note:** Before you can add an SWG override setting to a device, you must register the device as a roaming computer with Umbrella.
Umbrella applies the SWG override setting to a device using the device's origin ID.
You can list the roaming computers in your organization by sending a request to the Umbrella Roaming Computers API.

## Example Usage


### Basic Usage

Basic usage of the createsecurewebgatewaydevicesettings resource

```terraform
resource "umbrella_createsecurewebgatewaydevicesettings" "example" {
  name        = "example-createsecurewebgatewaydevicesettings"
  description = "Example createsecurewebgatewaydevicesettings resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`originIds`** (Set of String) - The list of origin IDs. The list can contain 1–100 origin IDs. Example: `["item1", "item2"]`
- **`value`** (String) - Specifies whether to enable the Secure Web Gateway (SWG) device settings. Valid values are: `0` or `1` where `1` indicates enable. Example: `"example"`


### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`failCount`** (Number) - The number of devices that failed to change the device setting.
- **`items`** (Set of String) - The list of device setting status properties.
- **`totalCount`** (Number) - The total number of devices that requested to update the device setting.
- **`successCount`** (Number) - The number of devices that successfully changed the device setting.



## Import

umbrella_createsecurewebgatewaydevicesettings can be imported using the resource ID:

```shell
terraform import umbrella_createsecurewebgatewaydevicesettings.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

