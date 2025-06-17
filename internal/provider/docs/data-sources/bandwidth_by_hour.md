---
page_title: "umbrella_bandwidth_by_hour Data Source - bandwidth_by_hour"
description: |-
  List the bandwidth in bytes within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Granular Events > Read-Only
---

# umbrella_bandwidth_by_hour (Data Source)

List the bandwidth in bytes within the timeframe. Only returns proxy data.

**Access Scope:** Reports > Granular Events > Read-Only

## Example Usage


### Basic Usage

Basic usage of the bandwidth_by_hour data source

```hcl
data "umbrella_bandwidth_by_hour" "example" {
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



