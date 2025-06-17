---
page_title: "umbrella_sites Resource - sites"
description: |-
  Create a site.
---

# umbrella_sites (Resource)

Create a site.

## Example Usage


### Basic Usage

Basic usage of the sites resource

```hcl
resource "umbrella_sites" "example" {
  # Add required attributes here
  name = "example-sites"
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
terraform import umbrella_sites.example 12345
```

