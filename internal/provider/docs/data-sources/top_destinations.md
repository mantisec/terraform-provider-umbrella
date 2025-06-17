---
page_title: "umbrella_top_destinations Data Source - top_destinations"
description: |-
  List the destinations by type of destination and the number of requests made
to this destination. Return the collection in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_destinations (Data Source)

List the destinations by type of destination and the number of requests made
to this destination. Return the collection in descending order.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_destinations data source

```hcl
data "umbrella_top_destinations" "example" {
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



