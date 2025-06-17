---
page_title: "umbrella_responses Data Source - responses"
description: |-
  Get the information about the API responses for the organization within a specific time period, including
the total number of API responses and the HTTP status codes.
---

# umbrella_responses (Data Source)

Get the information about the API responses for the organization within a specific time period, including
the total number of API responses and the HTTP status codes.

## Example Usage


### Basic Usage

Basic usage of the responses data source

```hcl
data "umbrella_responses" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



