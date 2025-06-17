---
page_title: "umbrella_threat_types Data Source - threat_types"
description: |-
  List the threat types by threat ID.

**Access Scope:** Reports > Utilities > Read-Only
---

# umbrella_threat_types (Data Source)

List the threat types by threat ID.

**Access Scope:** Reports > Utilities > Read-Only

## Example Usage


### Basic Usage

Basic usage of the threat_types data source

```hcl
data "umbrella_threat_types" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `meta` (String) 
- `data` (String) 



