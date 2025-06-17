---
page_title: "umbrella_devices Data Source - devices"
description: |-
  List the roaming computers that have a tag with the specified tag ID.
---

# umbrella_devices (Data Source)

List the roaming computers that have a tag with the specified tag ID.

## Example Usage


### Basic Usage

Basic usage of the devices data source

```hcl
data "umbrella_devices" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



