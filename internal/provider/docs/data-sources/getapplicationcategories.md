---
page_title: "umbrella_getapplicationcategories Data Source - terraform-provider-umbrella"
description: |-
  List all application categories.
---

# umbrella_getapplicationcategories (Data Source)

List all application categories.

## Example Usage


### Basic Usage

Basic usage of the getapplicationcategories data source

```terraform
data "umbrella_getapplicationcategories" "example" {
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



