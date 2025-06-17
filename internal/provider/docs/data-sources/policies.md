---
page_title: "umbrella_policies Data Source - policies"
description: |-
  List the policies for a network.
---

# umbrella_policies (Data Source)

List the policies for a network.

## Example Usage


### Basic Usage

Basic usage of the policies data source

```hcl
data "umbrella_policies" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



