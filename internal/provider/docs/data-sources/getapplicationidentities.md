---
page_title: "umbrella_getapplicationidentities Data Source - terraform-provider-umbrella"
description: |-
  List all identities for the application.
---

# umbrella_getapplicationidentities (Data Source)

List all identities for the application.

## Example Usage


### Basic Usage

Basic usage of the getapplicationidentities data source

```terraform
data "umbrella_getapplicationidentities" "example" {
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
- **`items`** (Set of String) - 



