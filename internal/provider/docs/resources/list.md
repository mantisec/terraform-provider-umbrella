---
page_title: "umbrella_list Resource - list"
description: |-
  List the Secure Web Gateway (SWG) override setting for devices in the organization.
---

# umbrella_list (Resource)

List the Secure Web Gateway (SWG) override setting for devices in the organization.

## Example Usage


### Basic Usage

Basic usage of the list resource

```hcl
resource "umbrella_list" "example" {
  # Add required attributes here
  name = "example-list"
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
terraform import umbrella_list.example 12345
```

