---
page_title: "umbrella_destinations Resource - destinations"
description: |-
  Add destinations to a destination list.
---

# umbrella_destinations (Resource)

Add destinations to a destination list.

## Example Usage


### Basic Usage

Basic usage of the destinations resource

```hcl
resource "umbrella_destinations" "example" {
  # Add required attributes here
  name = "example-destinations"
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
terraform import umbrella_destinations.example 12345
```

