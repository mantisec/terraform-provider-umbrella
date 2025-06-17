---
page_title: "umbrella_tunnels Resource - tunnels"
description: |-
  Add a new tunnel to the organization.
---

# umbrella_tunnels (Resource)

Add a new tunnel to the organization.

## Example Usage


### Basic Usage

Basic usage of the tunnels resource

```hcl
resource "umbrella_tunnels" "example" {
  # Add required attributes here
  name = "example-tunnels"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `message` (String) 



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_tunnels.example 12345
```

