---
page_title: "umbrella_summaries_by_destination Data Source - summaries_by_destination"
description: |-
  List the summaries by destination for the type of traffic.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_summaries_by_destination (Data Source)

List the summaries by destination for the type of traffic.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the summaries_by_destination data source

```hcl
data "umbrella_summaries_by_destination" "example" {
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



