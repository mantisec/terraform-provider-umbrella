---
page_title: "umbrella_virtualappliances Resource - virtualappliances"
description: |-
  Update a virtual appliance in the organization.
---

# umbrella_virtualappliances (Resource)

Update a virtual appliance in the organization.

## Example Usage


### Basic Usage

Basic usage of the virtualappliances resource

```hcl
resource "umbrella_virtualappliances" "example" {
  # Add required attributes here
  name = "example-virtualappliances"
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
terraform import umbrella_virtualappliances.example 12345
```

