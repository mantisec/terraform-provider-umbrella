---
page_title: "umbrella_networks Resource - networks"
description: |-
  Create a network. Before you can create a network, contact Support to get your IP range verified.
---

# umbrella_networks (Resource)

Create a network. Before you can create a network, contact Support to get your IP range verified.

## Example Usage


### Basic Usage

Basic usage of the networks resource

```hcl
resource "umbrella_networks" "example" {
  # Add required attributes here
  name = "example-networks"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `message` (String) Detailed error message
- `statusCode` (Number) HTTP status code
- `error` (String) Shorthand error message



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_networks.example 12345
```

