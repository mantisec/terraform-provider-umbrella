---
page_title: "umbrella_top_threat_types Data Source - top_threat_types"
description: |-
  List the top threat-types within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_threat_types (Data Source)

List the top threat-types within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_threat_types data source

```hcl
data "umbrella_top_threat_types" "example" {
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



