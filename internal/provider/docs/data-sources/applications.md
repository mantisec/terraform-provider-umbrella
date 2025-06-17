---
page_title: "umbrella_applications Data Source - applications"
description: |-
  Get an application by ID.
---

# umbrella_applications (Data Source)

Get an application by ID.

## Example Usage


### Basic Usage

Basic usage of the applications data source

```hcl
data "umbrella_applications" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



