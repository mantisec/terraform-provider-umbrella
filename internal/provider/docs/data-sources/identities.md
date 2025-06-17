---
page_title: "umbrella_identities Data Source - identities"
description: |-
  List identities for a specific protocol.
---

# umbrella_identities (Data Source)

List identities for a specific protocol.

## Example Usage


### Basic Usage

Basic usage of the identities data source

```hcl
data "umbrella_identities" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



