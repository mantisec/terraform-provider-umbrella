---
page_title: "umbrella_top_urls Data Source - top_urls"
description: |-
  List the top number of URLs that are requested for a certain domain.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_urls (Data Source)

List the top number of URLs that are requested for a certain domain.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_urls data source

```hcl
data "umbrella_top_urls" "example" {
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



