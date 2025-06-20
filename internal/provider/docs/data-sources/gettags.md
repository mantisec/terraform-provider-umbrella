---
page_title: "umbrella_gettags Data Source - terraform-provider-umbrella"
description: |-
  List the tags in the organization.
---

# umbrella_gettags (Data Source)

List the tags in the organization.

## Example Usage


### Basic Usage

Basic usage of the gettags data source

```terraform
data "umbrella_gettags" "example" {
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
- **`organizationId`** (Number) - The organization ID.
- **`name`** (String) - The name of the tag.
- **`originsModifiedAt`** (String) - The date and time (timestamp) that Umbrella modified the roaming computers. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.
- **`createdAt`** (String) - The date and time (timestamp) that Umbrella added the tag to the roaming computer. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.
- **`modifiedAt`** (String) - The date and time (timestamp) that Umbrella updated the tag on the roaming computer. The timestamp is an ISO 8601 formatted string. For example: `2023-04-12T23:20:50.52Z`.



