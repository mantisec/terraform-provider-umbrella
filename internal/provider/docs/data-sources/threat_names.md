---
page_title: "umbrella_threat_names Data Source - threat_names"
description: |-
  List the threat names.

**Access Scope:** Reports > Utilities > Read-Only
---

# umbrella_threat_names (Data Source)

List the threat names.

**Access Scope:** Reports > Utilities > Read-Only

## Example Usage


### Basic Usage

Basic usage of the threat_names data source

```hcl
data "umbrella_threat_names" "example" {
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



