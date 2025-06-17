---
page_title: "umbrella_top_categories Data Source - top_categories"
description: |-
  List the categories that received the greatest number of requests.
Order the number of requests in descending order.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_categories (Data Source)

List the categories that received the greatest number of requests.
Order the number of requests in descending order.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_categories data source

```hcl
data "umbrella_top_categories" "example" {
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



