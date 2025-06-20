---
page_title: "umbrella_deletesecurewebgatewaydevicesettings Resource - terraform-provider-umbrella"
description: |-
  Remove the Secure Web Gateway (SWG) override setting for the devices in the organization.
Once you remove the override setting on a device, Umbrella applies your organization's SWG setting to the device.
---

# umbrella_deletesecurewebgatewaydevicesettings (Resource)

Remove the Secure Web Gateway (SWG) override setting for the devices in the organization.
Once you remove the override setting on a device, Umbrella applies your organization's SWG setting to the device.

## Example Usage


### Basic Usage

Basic usage of the deletesecurewebgatewaydevicesettings resource

```terraform
resource "umbrella_deletesecurewebgatewaydevicesettings" "example" {
  name        = "example-deletesecurewebgatewaydevicesettings"
  description = "Example deletesecurewebgatewaydevicesettings resource"
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



## Import

umbrella_deletesecurewebgatewaydevicesettings can be imported using the resource ID:

```shell
terraform import umbrella_deletesecurewebgatewaydevicesettings.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

