---
page_title: "umbrella_getcnames Data Source - terraform-provider-umbrella"
description: |-
  List the cname information for the service providers console.
---

# umbrella_getcnames (Data Source)

List the cname information for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the getcnames data source

```terraform
data "umbrella_getcnames" "example" {
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
- **`cnameId`** (Number) - The cname ID.
- **`cname`** (String) - The cname for the service providers console.
- **`organizationId`** (Number) - The organization ID.



