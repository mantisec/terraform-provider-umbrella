---
page_title: "umbrella_summary Data Source - summary"
description: |-
  Get the summary of requests by the traffic type.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_summary (Data Source)

Get the summary of requests by the traffic type.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the summary data source

```hcl
data "umbrella_summary" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `data` (String) 
- `meta` (String) 



