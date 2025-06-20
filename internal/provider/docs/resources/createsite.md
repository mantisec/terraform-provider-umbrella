---
page_title: "umbrella_createsite Resource - terraform-provider-umbrella"
description: |-
  Create a site.
---

# umbrella_createsite (Resource)

Create a site.

## Example Usage


### Basic Usage

Basic usage of the createsite resource

```terraform
resource "umbrella_createsite" "example" {
  name        = "example-createsite"
  description = "Example createsite resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`name`** (String) - The name of the site. The name is a sequence of characters with a length from 1 through 255. Example: `"example-name"`


### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`originId`** (Number) - The origin ID of the site.
- **`internalNetworkCount`** (Number) - The number of internal networks that are attached to the site.
- **`vaCount`** (Number) - The number of virtual appliances that are attached to the site.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the site was created.
- **`siteId`** (Number) - The ID of the site.
- **`isDefault`** (Boolean) - Specifies whether the site is the default.
- **`type`** (String) - The type of the site.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the site was modified.



## Import

umbrella_createsite can be imported using the resource ID:

```shell
terraform import umbrella_createsite.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

