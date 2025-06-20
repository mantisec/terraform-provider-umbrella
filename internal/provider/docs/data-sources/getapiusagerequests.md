---
page_title: "umbrella_getapiusagerequests Data Source - terraform-provider-umbrella"
description: |-
  Get the information about the API requests for the organization within a specific time period, including
the total number of API requests for the type of client program.
---

# umbrella_getapiusagerequests (Data Source)

Get the information about the API requests for the organization within a specific time period, including
the total number of API requests for the type of client program.

## Example Usage


### Basic Usage

Basic usage of the getapiusagerequests data source

```terraform
data "umbrella_getapiusagerequests" "example" {
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
- **`from`** (String) - The date and time where to start reading in the collection.
- **`to`** (String) - The date and time where to stop reading in the collection.
- **`count`** (Number) - The total number of API requests.
- **`items`** (Set of String) - The information about the API requests.



