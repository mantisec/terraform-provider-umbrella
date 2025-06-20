---
page_title: "umbrella_createmanagedprovidercustomer Resource - terraform-provider-umbrella"
description: |-
  Create a customer for a managed provider.
---

# umbrella_createmanagedprovidercustomer (Resource)

Create a customer for a managed provider.

## Example Usage


### Basic Usage

Basic usage of the createmanagedprovidercustomer resource

```terraform
resource "umbrella_createmanagedprovidercustomer" "example" {
  name        = "example-createmanagedprovidercustomer"
  description = "Example createmanagedprovidercustomer resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`customerName`** (String) - The customer's organization name. Example: `"example-name"`
- **`seats`** (Number) - The number of users. Example: `123`


### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`customerId`** (Number) - The ID of the customer.



## Import

umbrella_createmanagedprovidercustomer can be imported using the resource ID:

```shell
terraform import umbrella_createmanagedprovidercustomer.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

