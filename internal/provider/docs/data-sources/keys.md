---
page_title: "umbrella_keys Data Source - keys"
description: |-
  Get the API key usage information, including the total number of API requests within a specific time period.
---

# umbrella_keys (Data Source)

Get the API key usage information, including the total number of API requests within a specific time period.

## Example Usage


### Basic Usage

Basic usage of the keys data source

```hcl
data "umbrella_keys" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



