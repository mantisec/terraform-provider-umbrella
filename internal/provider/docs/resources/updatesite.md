---
page_title: "umbrella_updatesite Resource - terraform-provider-umbrella"
description: |-
  Update a site.
---

# umbrella_updatesite (Resource)

Update a site.

## Example Usage


### Basic Usage

Basic usage of the updatesite resource

```terraform
resource "umbrella_updatesite" "example" {
  name        = "example-updatesite"
  description = "Example updatesite resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the site was modified.
- **`siteId`** (Number) - The ID of the site.
- **`isDefault`** (Boolean) - Specifies whether the site is the default.
- **`type`** (String) - The type of the site.
- **`vaCount`** (Number) - The number of virtual appliances that are attached to the site.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the site was created.
- **`originId`** (Number) - The origin ID of the site.
- **`name`** (String) - The name of the site.
- **`internalNetworkCount`** (Number) - The number of internal networks that are attached to the site.



## Import

umbrella_updatesite can be imported using the resource ID:

```shell
terraform import umbrella_updatesite.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

