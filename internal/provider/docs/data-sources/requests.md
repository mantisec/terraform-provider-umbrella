---
page_title: "umbrella_requests Data Source - requests"
description: |-
  Get the information about the API requests for the organization within a specific time period, including
the total number of API requests for the type of client program.
---

# umbrella_requests (Data Source)

Get the information about the API requests for the organization within a specific time period, including
the total number of API requests for the type of client program.

## Example Usage


### Basic Usage

Basic usage of the requests data source

```hcl
data "umbrella_requests" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



