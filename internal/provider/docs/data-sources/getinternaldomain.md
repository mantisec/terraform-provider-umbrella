---
page_title: "umbrella_getinternaldomain Data Source - terraform-provider-umbrella"
description: |-
  Get an internal domain.
---

# umbrella_getinternaldomain (Data Source)

Get an internal domain.

## Example Usage


### Basic Usage

Basic usage of the getinternaldomain data source

```terraform
data "umbrella_getinternaldomain" "example" {
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
- **`includeAllVAs`** (Boolean) - Specifies whether to apply the internal domain to all virtual appliances.
- **`includeAllMobileDevices`** (Boolean) - Specifies whether to apply the internal domain to all mobile devices.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was created.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was modified.
- **`siteIds`** (Set of String) - The list of site IDs associated with the domain.
- **`domain`** (String) - The domain name of the internal domain.
- **`description`** (String) - The description of the internal domain.



