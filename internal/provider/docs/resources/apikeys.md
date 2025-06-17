---
page_title: "umbrella_apikeys Resource - apikeys"
description: |-
  Create an API key with certain scopes, name, description, allowed IP addresses and CIDR blocks, and expiration.
The `description` and `allowedIPs` fields are optional.
---

# umbrella_apikeys (Resource)

Create an API key with certain scopes, name, description, allowed IP addresses and CIDR blocks, and expiration.
The `description` and `allowedIPs` fields are optional.

## Example Usage


### Basic Usage

Basic usage of the apikeys resource

```hcl
resource "umbrella_apikeys" "example" {
  # Add required attributes here
  name = "example-apikeys"
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
terraform import umbrella_apikeys.example 12345
```

