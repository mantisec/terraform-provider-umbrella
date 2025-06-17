---
page_title: "umbrella_amp_retrospective Data Source - amp_retrospective"
description: |-
  List all AMP retrospective activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_amp_retrospective (Data Source)

List all AMP retrospective activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the amp_retrospective data source

```hcl
data "umbrella_amp_retrospective" "example" {
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



