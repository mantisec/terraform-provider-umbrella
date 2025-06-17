---
page_title: "umbrella_accessrequests Data Source - accessrequests"
description: |-
  Get the access request details for the customer's organization.
---

# umbrella_accessrequests (Data Source)

Get the access request details for the customer's organization.

## Example Usage


### Basic Usage

Basic usage of the accessrequests data source

```hcl
data "umbrella_accessrequests" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



