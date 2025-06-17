---
page_title: "umbrella_refresh Resource - refresh"
description: |-
  Refresh the API key and secret.
---

# umbrella_refresh (Resource)

Refresh the API key and secret.

## Example Usage


### Basic Usage

Basic usage of the refresh resource

```hcl
resource "umbrella_refresh" "example" {
  # Add required attributes here
  name = "example-refresh"
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
terraform import umbrella_refresh.example 12345
```

