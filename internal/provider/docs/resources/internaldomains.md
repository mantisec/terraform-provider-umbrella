---
page_title: "umbrella_internaldomains Resource - internaldomains"
description: |-
  Create an internal domain. If you do not assign a list of sites to the internal domain, the internal domain
is associated with all sites in the organization.
---

# umbrella_internaldomains (Resource)

Create an internal domain. If you do not assign a list of sites to the internal domain, the internal domain
is associated with all sites in the organization.

## Example Usage


### Basic Usage

Basic usage of the internaldomains resource

```hcl
resource "umbrella_internaldomains" "example" {
  # Add required attributes here
  name = "example-internaldomains"
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
terraform import umbrella_internaldomains.example 12345
```

