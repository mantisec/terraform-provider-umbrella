---
page_title: "umbrella_tunnels Data Source - tunnels"
description: |-
  List the tunnels for an organization.
---

# umbrella_tunnels (Data Source)

List the tunnels for an organization.

## Example Usage


### Basic Usage

Basic usage of the tunnels data source

```hcl
data "umbrella_tunnels" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



