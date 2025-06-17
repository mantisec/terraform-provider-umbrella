---
page_title: "umbrella_top_identities Data Source - top_identities"
description: |-
  List the identities for the specific traffic type by the number of requests.
Sort the results in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_identities (Data Source)

List the identities for the specific traffic type by the number of requests.
Sort the results in descending order.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_identities data source

```hcl
data "umbrella_top_identities" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `data` (List of String) 
- `meta` (String) 



