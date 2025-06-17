---
page_title: "umbrella_tags Resource - tags"
description: |-
  Add a tag to the organization.
---

# umbrella_tags (Resource)

Add a tag to the organization.

## Example Usage


### Basic Usage

Basic usage of the tags resource

```hcl
resource "umbrella_tags" "example" {
  # Add required attributes here
  name = "example-tags"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_tags.example 12345
```

