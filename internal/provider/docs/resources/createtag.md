---
page_title: "umbrella_createtag Resource - terraform-provider-umbrella"
description: |-
  Add a tag to the organization.
---

# umbrella_createtag (Resource)

Add a tag to the organization.

## Example Usage


### Basic Usage

Basic usage of the createtag resource

```terraform
resource "umbrella_createtag" "example" {
  name        = "example-createtag"
  description = "Example createtag resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`organizationId`** (Number) - The organization ID.
- **`name`** (String) - The name of the tag.
- **`originsModifiedAt`** (String) - The date and time (timestamp) that Umbrella modified the roaming computers. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.
- **`createdAt`** (String) - The date and time (timestamp) that Umbrella added the tag to the roaming computer. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.
- **`modifiedAt`** (String) - The date and time (timestamp) that Umbrella updated the tag on the roaming computer. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.



## Import

umbrella_createtag can be imported using the resource ID:

```shell
terraform import umbrella_createtag.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

