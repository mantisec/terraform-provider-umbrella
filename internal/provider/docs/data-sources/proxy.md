---
page_title: "umbrella_proxy Data Source - proxy"
description: |-
  List all proxy entries within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_proxy (Data Source)

List all proxy entries within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the proxy data source

```hcl
data "umbrella_proxy" "example" {
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



