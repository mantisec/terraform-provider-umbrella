---
page_title: "umbrella_identities Resource - identities"
description: |-
  Get the identities information by providing a list of identity IDs in the request body.

**Access Scope:** Reports > Utilities > Read-Only
---

# umbrella_identities (Resource)

Get the identities information by providing a list of identity IDs in the request body.

**Access Scope:** Reports > Utilities > Read-Only

## Example Usage


### Basic Usage

Basic usage of the identities resource

```hcl
resource "umbrella_identities" "example" {
  # Add required attributes here
  name = "example-identities"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `meta` (String) 
- `data` (List of String) 



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_identities.example 12345
```

