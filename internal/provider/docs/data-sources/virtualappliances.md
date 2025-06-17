---
page_title: "umbrella_virtualappliances Data Source - virtualappliances"
description: |-
  List the virtual appliances in the organization.
---

# umbrella_virtualappliances (Data Source)

List the virtual appliances in the organization.

## Example Usage


### Basic Usage

Basic usage of the virtualappliances data source

```hcl
data "umbrella_virtualappliances" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



