---
page_title: "umbrella_getapplicationsappdiscovery Data Source - terraform-provider-umbrella"
description: |-
  List all discovered applications.
---

# umbrella_getapplicationsappdiscovery (Data Source)

List all discovered applications.

## Example Usage


### Basic Usage

Basic usage of the getapplicationsappdiscovery data source

```terraform
data "umbrella_getapplicationsappdiscovery" "example" {
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
- **`items`** (Set of String) - The list of application properties.
- **`currentPage`** (Number) - The index of the current page in the collection.
- **`totalPages`** (Number) - The total number of pages in the collection.
- **`itemsCount`** (Number) - The number of items in the collection.



