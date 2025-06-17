---
page_title: "umbrella_identities Data Source - identities"
description: |-
  List the identities.

**Access Scope:** Reports > Utilities > Read-Only
---

# umbrella_identities (Data Source)

List the identities.

**Access Scope:** Reports > Utilities > Read-Only

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
- `data` (List of String) 
- `meta` (String) 



