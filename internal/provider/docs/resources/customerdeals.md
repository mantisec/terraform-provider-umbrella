---
page_title: "umbrella_customerdeals Resource - customerdeals"
description: |-
  Update the customer deal.
---

# umbrella_customerdeals (Resource)

Update the customer deal.

## Example Usage


### Basic Usage

Basic usage of the customerdeals resource

```hcl
resource "umbrella_customerdeals" "example" {
  # Add required attributes here
  name = "example-customerdeals"
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
terraform import umbrella_customerdeals.example 12345
```

