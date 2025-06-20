---
page_title: "umbrella_getapiusagekeys Data Source - terraform-provider-umbrella"
description: |-
  Get the API key usage information, including the total number of API requests within a specific time period.
---

# umbrella_getapiusagekeys (Data Source)

Get the API key usage information, including the total number of API requests within a specific time period.

## Example Usage


### Basic Usage

Basic usage of the getapiusagekeys data source

```terraform
data "umbrella_getapiusagekeys" "example" {
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
- **`items`** (Set of String) - The information about the API key usage.



