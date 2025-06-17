---
page_title: "umbrella_categories_by_hour Data Source - categories_by_hour"
description: |-
  List the activity volume within the timeframe by type of category.

**Access Scope:** Reports > Granular Events > Read-Only
---

# umbrella_categories_by_hour (Data Source)

List the activity volume within the timeframe by type of category.

**Access Scope:** Reports > Granular Events > Read-Only

## Example Usage


### Basic Usage

Basic usage of the categories_by_hour data source

```hcl
data "umbrella_categories_by_hour" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `meta` (String) 
- `data` (List of String) 



