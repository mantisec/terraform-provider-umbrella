---
page_title: "umbrella_keys Resource - keys"
description: |-
  Update and rotate the tunnel credentials.
---

# umbrella_keys (Resource)

Update and rotate the tunnel credentials.

## Example Usage


### Basic Usage

Basic usage of the keys resource

```hcl
resource "umbrella_keys" "example" {
  # Add required attributes here
  name = "example-keys"
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
terraform import umbrella_keys.example 12345
```

