---
page_title: "umbrella_updatemanagedprovidercustomer Resource - terraform-provider-umbrella"
description: |-
  Update a customer for a managed provider.
---

# umbrella_updatemanagedprovidercustomer (Resource)

Update a customer for a managed provider.

## Example Usage


### Basic Usage

Basic usage of the updatemanagedprovidercustomer resource

```terraform
resource "umbrella_updatemanagedprovidercustomer" "example" {
  name        = "example-updatemanagedprovidercustomer"
  description = "Example updatemanagedprovidercustomer resource"
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



## Import

umbrella_updatemanagedprovidercustomer can be imported using the resource ID:

```shell
terraform import umbrella_updatemanagedprovidercustomer.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

