---
page_title: "umbrella_activity Data Source - activity"
description: |-
  List all activities (dns/proxy/firewall/intrusion) within the timeframe.
**Note:** The IP activity report is not available.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_activity (Data Source)

List all activities (dns/proxy/firewall/intrusion) within the timeframe.
**Note:** The IP activity report is not available.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the activity data source

```hcl
data "umbrella_activity" "example" {
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



