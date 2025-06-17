---
page_title: "umbrella_networkdevices Data Source - networkdevices"
description: |-
  List the network devices.
---

# umbrella_networkdevices (Data Source)

List the network devices.

## Example Usage


### Basic Usage

Basic usage of the networkdevices data source

```hcl
data "umbrella_networkdevices" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



