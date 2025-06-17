---
page_title: "umbrella_logos Resource - logos"
description: |-
  Create a logo for the service providers console.
---

# umbrella_logos (Resource)

Create a logo for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the logos resource

```hcl
resource "umbrella_logos" "example" {
  # Add required attributes here
  name = "example-logos"
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
terraform import umbrella_logos.example 12345
```

