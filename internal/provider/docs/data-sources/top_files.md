---
page_title: "umbrella_top_files Data Source - top_files"
description: |-
  List the top files within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_files (Data Source)

List the top files within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_files data source

```hcl
data "umbrella_top_files" "example" {
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



