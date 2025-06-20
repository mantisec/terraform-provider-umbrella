---
page_title: "umbrella_updateinternaldomain Resource - terraform-provider-umbrella"
description: |-
  Update an internal domain.
---

# umbrella_updateinternaldomain (Resource)

Update an internal domain.

## Example Usage


### Basic Usage

Basic usage of the updateinternaldomain resource

```terraform
resource "umbrella_updateinternaldomain" "example" {
  name        = "example-updateinternaldomain"
  description = "Example updateinternaldomain resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`description`** (String) - The description of the internal domain.
- **`includeAllVAs`** (Boolean) - Specifies whether to apply the internal domain to all virtual appliances.
- **`includeAllMobileDevices`** (Boolean) - Specifies whether to apply the internal domain to all mobile devices.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was created.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the internal domain was modified.
- **`siteIds`** (Set of String) - The list of site IDs associated with the domain.
- **`domain`** (String) - The domain name of the internal domain.



## Import

umbrella_updateinternaldomain can be imported using the resource ID:

```shell
terraform import umbrella_updateinternaldomain.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

