---
page_title: "umbrella_networkdevices Resource - networkdevices"
description: |-
  Create a network device.
---

# umbrella_networkdevices (Resource)

Create a network device.

## Example Usage


### Basic Usage

Basic usage of the networkdevices resource

```hcl
resource "umbrella_networkdevices" "example" {
  # Add required attributes here
  name = "example-networkdevices"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_networkdevices.example 12345
```

