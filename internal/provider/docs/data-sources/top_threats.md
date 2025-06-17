---
page_title: "umbrella_top_threats Data Source - top_threats"
description: |-
  Get the top threats within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_threats (Data Source)

Get the top threats within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_threats data source

```hcl
data "umbrella_top_threats" "example" {
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



