---
page_title: "umbrella_intrusion Data Source - intrusion"
description: |-
  List all Intrusion Prevention System (IPS) activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_intrusion (Data Source)

List all Intrusion Prevention System (IPS) activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the intrusion data source

```hcl
data "umbrella_intrusion" "example" {
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



