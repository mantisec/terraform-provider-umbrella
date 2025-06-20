---
page_title: "umbrella_getapiusageresponses Data Source - terraform-provider-umbrella"
description: |-
  Get the information about the API responses for the organization within a specific time period, including
the total number of API responses and the HTTP status codes.
---

# umbrella_getapiusageresponses (Data Source)

Get the information about the API responses for the organization within a specific time period, including
the total number of API responses and the HTTP status codes.

## Example Usage


### Basic Usage

Basic usage of the getapiusageresponses data source

```terraform
data "umbrella_getapiusageresponses" "example" {
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
- **`count`** (Number) - The total number of API requests.
- **`items`** (Set of String) - The list of information about API responses for the API requests.
- **`from`** (String) - The date and time where to start reading in the collection.
- **`to`** (String) - The date and time where to stop reading in the collection.



