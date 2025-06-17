---
page_title: "umbrella_identities Resource - identities"
description: |-
  Add an identity to your policy.
Policy changes may require up to 20 minutes to take effect globally.
For DNS policies, TTLs, caching, and session reuse may cause some devices
and domains to appear to take longer to update.
---

# umbrella_identities (Resource)

Add an identity to your policy.
Policy changes may require up to 20 minutes to take effect globally.
For DNS policies, TTLs, caching, and session reuse may cause some devices
and domains to appear to take longer to update.

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



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_identities.example 12345
```

