---
page_title: "umbrella_ip Data Source - ip"
description: |-
  (Deprecated) List all IP activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_ip (Data Source)

(Deprecated) List all IP activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the ip data source

```hcl
data "umbrella_ip" "example" {
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



