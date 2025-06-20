---
page_title: "umbrella_getcname Data Source - terraform-provider-umbrella"
description: |-
  Get the cname information for the service providers console.
---

# umbrella_getcname (Data Source)

Get the cname information for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the getcname data source

```terraform
data "umbrella_getcname" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`organizationId`** (Number) - The organization ID.
- **`cnameId`** (Number) - The cname ID.
- **`cname`** (String) - The cname for the service providers console.



