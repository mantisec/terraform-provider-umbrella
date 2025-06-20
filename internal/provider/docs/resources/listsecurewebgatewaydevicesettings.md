---
page_title: "umbrella_listsecurewebgatewaydevicesettings Resource - terraform-provider-umbrella"
description: |-
  List the Secure Web Gateway (SWG) override setting for devices in the organization.
---

# umbrella_listsecurewebgatewaydevicesettings (Resource)

List the Secure Web Gateway (SWG) override setting for devices in the organization.

## Example Usage


### Basic Usage

Basic usage of the listsecurewebgatewaydevicesettings resource

```terraform
resource "umbrella_listsecurewebgatewaydevicesettings" "example" {
  name        = "example-listsecurewebgatewaydevicesettings"
  description = "Example listsecurewebgatewaydevicesettings resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`originIds`** (Set of String) - The list of origin IDs. The list can contain 1–100 origin IDs. Example: `["item1", "item2"]`


### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`modifiedAt`** (String) - The date and time when the settings on the device were modified. The timestamp is in the ISO 8601 date format.
- **`originId`** (Number) - The origin ID of the device.
- **`name`** (String) - The name of the device setting.
- **`value`** (String) - Specifies whether to enable the Secure Web Gateway (SWG) device settings. Valid values are: `0` or `1` where `1` indicates enable.



## Import

umbrella_listsecurewebgatewaydevicesettings can be imported using the resource ID:

```shell
terraform import umbrella_listsecurewebgatewaydevicesettings.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

