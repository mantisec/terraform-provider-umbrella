---
page_title: "umbrella_addtagtodevices Resource - terraform-provider-umbrella"
description: |-
  For the tag ID, add the tag to the roaming computers.
---

# umbrella_addtagtodevices (Resource)

For the tag ID, add the tag to the roaming computers.

## Example Usage


### Basic Usage

Basic usage of the addtagtodevices resource

```terraform
resource "umbrella_addtagtodevices" "example" {
  name        = "example-addtagtodevices"
  description = "Example addtagtodevices resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`tagId`** (Number) - The unique ID of the tag.
- **`addOrigins`** (Set of String) - The list of roaming computers (origin IDs) that have the tag.
- **`removeOrigins`** (Set of String) - The list of roaming computers (origin IDs) that have the tag removed.



## Import

umbrella_addtagtodevices can be imported using the resource ID:

```shell
terraform import umbrella_addtagtodevices.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

