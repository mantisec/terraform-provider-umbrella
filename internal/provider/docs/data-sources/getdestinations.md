---
page_title: "umbrella_getdestinations Data Source - terraform-provider-umbrella"
description: |-
  Get destinations in a destination list.
---

# umbrella_getdestinations (Data Source)

Get destinations in a destination list.

## Example Usage


### Basic Usage

Basic usage of the getdestinations data source

```terraform
data "umbrella_getdestinations" "example" {
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
- **`data`** (Set of String) - The list of destinations in the destination list.
- **`destination`** (String) - A domain, URL, or IP.
- **`type`** (String) - The type of the destination.
- **`comment`** (String) - The comment about the destination.
- **`createdAt`** (String) - The date and time when the destination list was created.



