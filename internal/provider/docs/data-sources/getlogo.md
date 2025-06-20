---
page_title: "umbrella_getlogo Data Source - terraform-provider-umbrella"
description: |-
  Get a logo for the service providers console.
---

# umbrella_getlogo (Data Source)

Get a logo for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the getlogo data source

```terraform
data "umbrella_getlogo" "example" {
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
- **`imageUrl`** (String) - The URL for the logo.
- **`modifiedAt`** (Number) - The time when the logo was last modified. The time is specified in milliseconds.
- **`brandingTypeId`** (Number) - The ID of the type of branding.
- **`imageKey`** (String) - The key for the logo.
- **`token`** (String) - The randomly generated number for the logo.
- **`createdAt`** (Number) - The time when the logo was created. The time is specified in milliseconds.
- **`enabled`** (Boolean) - Specify whether the logo is in use.
- **`organizationId`** (String) - The ID of the organization associated with the logo.



