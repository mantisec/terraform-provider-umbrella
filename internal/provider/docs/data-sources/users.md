---
page_title: "umbrella_users Data Source - users"
description: |-
  List the user accounts in the organization.
---

# umbrella_users (Data Source)

List the user accounts in the organization.

## Example Usage


### Basic Usage

Basic usage of the users data source

```hcl
data "umbrella_users" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



