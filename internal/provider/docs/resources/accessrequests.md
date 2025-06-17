---
page_title: "umbrella_accessrequests Resource - accessrequests"
description: |-
  Update the access request for the customer's organization.
---

# umbrella_accessrequests (Resource)

Update the access request for the customer's organization.

## Example Usage


### Basic Usage

Basic usage of the accessrequests resource

```hcl
resource "umbrella_accessrequests" "example" {
  # Add required attributes here
  name = "example-accessrequests"
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
terraform import umbrella_accessrequests.example 12345
```

