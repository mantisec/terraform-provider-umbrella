---
page_title: "umbrella_createinternaldomain Resource - terraform-provider-umbrella"
description: |-
  Create an internal domain. If you do not assign a list of sites to the internal domain, the internal domain
is associated with all sites in the organization.
---

# umbrella_createinternaldomain (Resource)

Create an internal domain. If you do not assign a list of sites to the internal domain, the internal domain
is associated with all sites in the organization.

## Example Usage


### Basic Usage

Basic usage of the createinternaldomain resource

```terraform
resource "umbrella_createinternaldomain" "example" {
  name        = "example-createinternaldomain"
  description = "Example createinternaldomain resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`siteIds`** (Set of String) - The list of site IDs associated with the domain.
- **`domain`** (String) - The domain name of the internal domain.
- **`description`** (String) - The description of the internal domain.
- **`includeAllVAs`** (Boolean) - Specifies whether to apply the internal domain to all virtual appliances.
- **`includeAllMobileDevices`** (Boolean) - Specifies whether to apply the internal domain to all mobile devices.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was created.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was modified.



## Import

umbrella_createinternaldomain can be imported using the resource ID:

```shell
terraform import umbrella_createinternaldomain.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

