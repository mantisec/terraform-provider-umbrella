---
page_title: "umbrella_internal Data Source - internal"
description: |-
  List the top internal IP addresses.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_internal (Data Source)

List the top internal IP addresses.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the internal data source

```hcl
data "umbrella_internal" "example" {
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



