---
page_title: "umbrella_tags Data Source - tags"
description: |-
  List the tags in the organization.
---

# umbrella_tags (Data Source)

List the tags in the organization.

## Example Usage


### Basic Usage

Basic usage of the tags data source

```hcl
data "umbrella_tags" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



