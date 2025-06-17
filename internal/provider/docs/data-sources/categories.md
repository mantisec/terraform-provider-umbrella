---
page_title: "umbrella_categories Data Source - categories"
description: |-
  List the Provider categories

**Access Scope:** Reports > Customer > Read-Only
---

# umbrella_categories (Data Source)

List the Provider categories

**Access Scope:** Reports > Customer > Read-Only

## Example Usage


### Basic Usage

Basic usage of the categories data source

```hcl
data "umbrella_categories" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `data` (List of String) 
- `meta` (String) 



