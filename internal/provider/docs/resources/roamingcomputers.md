---
page_title: "umbrella_roamingcomputers Resource - roamingcomputers"
description: |-
  Update a roaming computer in the organization.
---

# umbrella_roamingcomputers (Resource)

Update a roaming computer in the organization.

## Example Usage


### Basic Usage

Basic usage of the roamingcomputers resource

```hcl
resource "umbrella_roamingcomputers" "example" {
  # Add required attributes here
  name = "example-roamingcomputers"
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
terraform import umbrella_roamingcomputers.example 12345
```

