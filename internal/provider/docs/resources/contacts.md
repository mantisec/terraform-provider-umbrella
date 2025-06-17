---
page_title: "umbrella_contacts Resource - contacts"
description: |-
  Create a contact for the service providers console.
---

# umbrella_contacts (Resource)

Create a contact for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the contacts resource

```hcl
resource "umbrella_contacts" "example" {
  # Add required attributes here
  name = "example-contacts"
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
terraform import umbrella_contacts.example 12345
```

