---
page_title: "umbrella_customers Resource - customers"
description: |-
  Create a customer for a provider.
---

# umbrella_customers (Resource)

Create a customer for a provider.

## Example Usage


### Basic Usage

Basic usage of the customers resource

```hcl
resource "umbrella_customers" "example" {
  # Add required attributes here
  name = "example-customers"
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
terraform import umbrella_customers.example 12345
```

