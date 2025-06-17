---
page_title: "umbrella_identity_distribution Data Source - identity_distribution"
description: |-
  List the number of requests by identity types.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_identity_distribution (Data Source)

List the number of requests by identity types.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the identity_distribution data source

```hcl
data "umbrella_identity_distribution" "example" {
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



