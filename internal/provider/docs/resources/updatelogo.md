---
page_title: "umbrella_updatelogo Resource - terraform-provider-umbrella"
description: |-
  Update a logo for the service providers console.
---

# umbrella_updatelogo (Resource)

Update a logo for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the updatelogo resource

```terraform
resource "umbrella_updatelogo" "example" {
  name        = "example-updatelogo"
  description = "Example updatelogo resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`imageKey`** (String) - The key for the logo.
- **`token`** (String) - The randomly generated number for the logo.
- **`createdAt`** (Number) - The time when the logo was created. The time is specified in milliseconds.
- **`enabled`** (Boolean) - Specify whether the logo is in use.
- **`organizationId`** (String) - The ID of the organization associated with the logo.
- **`imageUrl`** (String) - The URL for the logo.
- **`modifiedAt`** (Number) - The time when the logo was last modified. The time is specified in milliseconds.
- **`brandingTypeId`** (Number) - The ID of the type of branding.



## Import

umbrella_updatelogo can be imported using the resource ID:

```shell
terraform import umbrella_updatelogo.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

