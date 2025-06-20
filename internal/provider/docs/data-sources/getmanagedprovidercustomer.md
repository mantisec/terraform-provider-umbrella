---
page_title: "umbrella_getmanagedprovidercustomer Data Source - terraform-provider-umbrella"
description: |-
  Get a customer for the managed provider.
---

# umbrella_getmanagedprovidercustomer (Data Source)

Get a customer for the managed provider.

## Example Usage


### Basic Usage

Basic usage of the getmanagedprovidercustomer data source

```terraform
data "umbrella_getmanagedprovidercustomer" "example" {
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
- **`seats`** (Number) - The number of users.
- **`customerId`** (Number) - The ID of the customer.
- **`customerName`** (String) - The customer's organization name.



