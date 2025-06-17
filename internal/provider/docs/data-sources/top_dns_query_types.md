---
page_title: "umbrella_top_dns_query_types Data Source - top_dns_query_types"
description: |-
  List the top types of DNS query.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_dns_query_types (Data Source)

List the top types of DNS query.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_dns_query_types data source

```hcl
data "umbrella_top_dns_query_types" "example" {
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



