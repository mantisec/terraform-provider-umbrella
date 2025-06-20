---
page_title: "umbrella_updatecustomerdeals Resource - terraform-provider-umbrella"
description: |-
  Update the customer deal.
---

# umbrella_updatecustomerdeals (Resource)

Update the customer deal.

## Example Usage


### Basic Usage

Basic usage of the updatecustomerdeals resource

```terraform
resource "umbrella_updatecustomerdeals" "example" {
  name        = "example-updatecustomerdeals"
  description = "Example updatecustomerdeals resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`dealId`** (String) - The deal ID.
- **`endCustomer`** (String) - The type of customer.
- **`majorLineItems`** (Set of String) - The list of essential components of the deal.
- **`canStampDeal`** (Boolean) - Specify whether to approve the deal.
- **`trialIds`** (Set of String) - The list of trial IDs.



## Import

umbrella_updatecustomerdeals can be imported using the resource ID:

```shell
terraform import umbrella_updatecustomerdeals.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

