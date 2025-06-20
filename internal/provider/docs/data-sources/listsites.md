---
page_title: "umbrella_listsites Data Source - terraform-provider-umbrella"
description: |-
  List the sites in the organization.
---

# umbrella_listsites (Data Source)

List the sites in the organization.

## Example Usage


### Basic Usage

Basic usage of the listsites data source

```terraform
data "umbrella_listsites" "example" {
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
- **`vaCount`** (Number) - The number of virtual appliances that are attached to the site.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the site was created.
- **`originId`** (Number) - The origin ID of the site.
- **`name`** (String) - The name of the site.
- **`internalNetworkCount`** (Number) - The number of internal networks that are attached to the site.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the site was modified.
- **`siteId`** (Number) - The ID of the site.
- **`isDefault`** (Boolean) - Specifies whether the site is the default.
- **`type`** (String) - The type of the site.



