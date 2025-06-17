---
page_title: "umbrella_applications Resource - applications"
description: |-
  Update the label for the application. Provide an application label in the request body.
---

# umbrella_applications (Resource)

Update the label for the application. Provide an application label in the request body.

## Example Usage


### Basic Usage

Basic usage of the applications resource

```hcl
resource "umbrella_applications" "example" {
  # Add required attributes here
  name = "example-applications"
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
terraform import umbrella_applications.example 12345
```

