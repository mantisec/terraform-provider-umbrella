---
page_title: "umbrella_getcustomerdeals Data Source - terraform-provider-umbrella"
description: |-
  Get the information about deals available to customer.
---

# umbrella_getcustomerdeals (Data Source)

Get the information about deals available to customer.

## Example Usage


### Basic Usage

Basic usage of the getcustomerdeals data source

```terraform
data "umbrella_getcustomerdeals" "example" {
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
- **`dealId`** (String) - The deal ID.
- **`endCustomer`** (String) - The type of customer.
- **`majorLineItems`** (Set of String) - The list of essential components of the deal.
- **`canStampDeal`** (Boolean) - Specify whether to approve the deal.
- **`trialIds`** (Set of String) - The list of trial IDs.



