---
page_title: "umbrella_rotatekey Resource - rotatekey"
description: |-
  Rotate the Cisco-managed S3 bucket key for the organization.
---

# umbrella_rotatekey (Resource)

Rotate the Cisco-managed S3 bucket key for the organization.

## Example Usage


### Basic Usage

Basic usage of the rotatekey resource

```hcl
resource "umbrella_rotatekey" "example" {
  # Add required attributes here
  name = "example-rotatekey"
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
terraform import umbrella_rotatekey.example 12345
```

