---
page_title: "umbrella_categories_by_timerange Data Source - categories_by_timerange"
description: |-
  List the activity volume within the timeframe by category.

**Access Scope:** Reports > Granular Events > Read-Only
---

# umbrella_categories_by_timerange (Data Source)

List the activity volume within the timeframe by category.

**Access Scope:** Reports > Granular Events > Read-Only

## Example Usage


### Basic Usage

Basic usage of the categories_by_timerange data source

```hcl
data "umbrella_categories_by_timerange" "example" {
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



