---
page_title: "umbrella_customers Data Source - customers"
description: |-
  List the customers for the managed provider.
---

# umbrella_customers (Data Source)

List the customers for the managed provider.

## Example Usage


### Basic Usage

Basic usage of the customers data source

```hcl
data "umbrella_customers" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



