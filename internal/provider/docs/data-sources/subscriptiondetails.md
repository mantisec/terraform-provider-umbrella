---
page_title: "umbrella_subscriptiondetails Data Source - subscriptiondetails"
description: |-
  Get the subscription details for the customer's organization.
---

# umbrella_subscriptiondetails (Data Source)

Get the subscription details for the customer's organization.

## Example Usage


### Basic Usage

Basic usage of the subscriptiondetails data source

```hcl
data "umbrella_subscriptiondetails" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



