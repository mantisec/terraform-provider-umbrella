---
page_title: "umbrella_listinternaldomains Data Source - terraform-provider-umbrella"
description: |-
  List the internal domains.
---

# umbrella_listinternaldomains (Data Source)

List the internal domains.

## Example Usage


### Basic Usage

Basic usage of the listinternaldomains data source

```terraform
data "umbrella_listinternaldomains" "example" {
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
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was created.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was modified.
- **`siteIds`** (Set of String) - The list of site IDs associated with the domain.
- **`domain`** (String) - The domain name of the internal domain.
- **`description`** (String) - The description of the internal domain.
- **`includeAllVAs`** (Boolean) - Specifies whether to apply the internal domain to all virtual appliances.
- **`includeAllMobileDevices`** (Boolean) - Specifies whether to apply the internal domain to all mobile devices.



