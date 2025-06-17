---
page_title: "umbrella_cnames Resource - cnames"
description: |-
  Update the cname information for the service providers console.
---

# umbrella_cnames (Resource)

Update the cname information for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the cnames resource

```hcl
resource "umbrella_cnames" "example" {
  # Add required attributes here
  name = "example-cnames"
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
terraform import umbrella_cnames.example 12345
```

