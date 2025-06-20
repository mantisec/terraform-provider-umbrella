---
page_title: "umbrella_listmanagedprovidercustomers Data Source - terraform-provider-umbrella"
description: |-
  List the customers for the managed provider.
---

# umbrella_listmanagedprovidercustomers (Data Source)

List the customers for the managed provider.

## Example Usage


### Basic Usage

Basic usage of the listmanagedprovidercustomers data source

```terraform
data "umbrella_listmanagedprovidercustomers" "example" {
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



