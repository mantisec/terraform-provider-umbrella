---
page_title: "umbrella_info Data Source - info"
description: |-
  List the information about the applications.
---

# umbrella_info (Data Source)

List the information about the applications.

## Example Usage


### Basic Usage

Basic usage of the info data source

```hcl
data "umbrella_info" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



