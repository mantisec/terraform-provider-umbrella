---
page_title: "umbrella_getprotocols Data Source - terraform-provider-umbrella"
description: |-
  List all protocols.
---

# umbrella_getprotocols (Data Source)

List all protocols.

## Example Usage


### Basic Usage

Basic usage of the getprotocols data source

```terraform
data "umbrella_getprotocols" "example" {
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



