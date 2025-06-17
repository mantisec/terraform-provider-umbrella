---
page_title: "umbrella_total_requests Data Source - total_requests"
description: |-
  Get the count of the total requests.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_total_requests (Data Source)

Get the count of the total requests.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the total_requests data source

```hcl
data "umbrella_total_requests" "example" {
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



