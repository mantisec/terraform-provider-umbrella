---
page_title: "umbrella_getprotocolidentities Data Source - terraform-provider-umbrella"
description: |-
  List identities for a specific protocol.
---

# umbrella_getprotocolidentities (Data Source)

List identities for a specific protocol.

## Example Usage


### Basic Usage

Basic usage of the getprotocolidentities data source

```terraform
data "umbrella_getprotocolidentities" "example" {
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



