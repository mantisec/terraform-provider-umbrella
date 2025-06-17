---
page_title: "umbrella_internalnetworks Resource - internalnetworks"
description: |-
  Create an internal network.
---

# umbrella_internalnetworks (Resource)

Create an internal network.

## Example Usage


### Basic Usage

Basic usage of the internalnetworks resource

```hcl
resource "umbrella_internalnetworks" "example" {
  # Add required attributes here
  name = "example-internalnetworks"
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
terraform import umbrella_internalnetworks.example 12345
```

