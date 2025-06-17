---
page_title: "umbrella_top_ips Data Source - top_ips"
description: |-
  List the top IP addresses.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_ips (Data Source)

List the top IP addresses.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_ips data source

```hcl
data "umbrella_top_ips" "example" {
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



