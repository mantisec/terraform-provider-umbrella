---
page_title: "umbrella_summary Data Source - summary"
description: |-
  Get the total number API requests, and the counts of the successful and failed API requests within a specific time period.
---

# umbrella_summary (Data Source)

Get the total number API requests, and the counts of the successful and failed API requests within a specific time period.

## Example Usage


### Basic Usage

Basic usage of the summary data source

```hcl
data "umbrella_summary" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



