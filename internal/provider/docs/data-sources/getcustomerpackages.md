---
page_title: "umbrella_getcustomerpackages Data Source - terraform-provider-umbrella"
description: |-
  Get the packages for the trial customer.
---

# umbrella_getcustomerpackages (Data Source)

Get the packages for the trial customer.

## Example Usage


### Basic Usage

Basic usage of the getcustomerpackages data source

```terraform
data "umbrella_getcustomerpackages" "example" {
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
- **`name`** (String) - The name of the package.
- **`pkgSeatMin`** (Number) - The minimum number of seats for the package.
- **`ppovSeatMin`** (Number) - The minimum number of seats for the customer trial.



