---
page_title: "umbrella_dns Data Source - dns"
description: |-
  List all DNS entries within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_dns (Data Source)

List all DNS entries within the timeframe.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the dns data source

```hcl
data "umbrella_dns" "example" {
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



