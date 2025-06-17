---
page_title: "umbrella_firewall Data Source - firewall"
description: |-
  List all firewall activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_firewall (Data Source)

List all firewall activity within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the firewall data source

```hcl
data "umbrella_firewall" "example" {
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



