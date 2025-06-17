---
page_title: "umbrella_cnames Data Source - cnames"
description: |-
  Get the cname information for the service providers console.
---

# umbrella_cnames (Data Source)

Get the cname information for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the cnames data source

```hcl
data "umbrella_cnames" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



