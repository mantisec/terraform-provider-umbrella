---
page_title: "umbrella_applications Data Source - applications"
description: |-
  List all discovered applications.
---

# umbrella_applications (Data Source)

List all discovered applications.

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



