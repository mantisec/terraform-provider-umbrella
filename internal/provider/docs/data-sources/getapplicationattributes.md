---
page_title: "umbrella_getapplicationattributes Data Source - terraform-provider-umbrella"
description: |-
  List all attributes for the application.
---

# umbrella_getapplicationattributes (Data Source)

List all attributes for the application.

## Example Usage


### Basic Usage

Basic usage of the getapplicationattributes data source

```terraform
data "umbrella_getapplicationattributes" "example" {
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
- **`attributesCategories`** (Set of String) - 



