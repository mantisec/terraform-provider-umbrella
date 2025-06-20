---
page_title: "umbrella_deletetagondevices Resource - terraform-provider-umbrella"
description: |-
  For the tag ID, remove the tag on the roaming computers in the organization.
After the delete, if the tag is not assigned to any roaming computers in the organization, Umbrella removes the tag from
the organization. You can recreate the tag for your organization.
---

# umbrella_deletetagondevices (Resource)

For the tag ID, remove the tag on the roaming computers in the organization.
After the delete, if the tag is not assigned to any roaming computers in the organization, Umbrella removes the tag from
the organization. You can recreate the tag for your organization.

## Example Usage


### Basic Usage

Basic usage of the deletetagondevices resource

```terraform
resource "umbrella_deletetagondevices" "example" {
  name        = "example-deletetagondevices"
  description = "Example deletetagondevices resource"
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

umbrella_deletetagondevices can be imported using the resource ID:

```shell
terraform import umbrella_deletetagondevices.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

