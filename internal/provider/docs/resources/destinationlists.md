---
page_title: "umbrella_destinationlists Resource - destinationlists"
description: |-
  Create a destination list in your organization.
---

# umbrella_destinationlists (Resource)

Create a destination list in your organization.

## Example Usage


### Basic Usage

Basic usage of the destinationlists resource

```hcl
resource "umbrella_destinationlists" "example" {
  # Add required attributes here
  name = "example-destinationlists"
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
terraform import umbrella_destinationlists.example 12345
```

