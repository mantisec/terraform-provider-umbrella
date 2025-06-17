---
page_title: "umbrella_summaries_by_category Data Source - summaries_by_category"
description: |-
  List the summaries of requests by category.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_summaries_by_category (Data Source)

List the summaries of requests by category.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the summaries_by_category data source

```hcl
data "umbrella_summaries_by_category" "example" {
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



