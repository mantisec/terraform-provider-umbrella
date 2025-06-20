---
page_title: "umbrella_getapplicationsinfo Data Source - terraform-provider-umbrella"
description: |-
  List the information about the applications.
---

# umbrella_getapplicationsinfo (Data Source)

List the information about the applications.

## Example Usage


### Basic Usage

Basic usage of the getapplicationsinfo data source

```terraform
data "umbrella_getapplicationsinfo" "example" {
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
- **`currentPage`** (Number) - The index of the current page in the collection.
- **`totalPages`** (Number) - The total number of pages in the collection.
- **`itemsCount`** (Number) - The number of items in the collection.
- **`items`** (Set of String) - The list of information about the applications.



