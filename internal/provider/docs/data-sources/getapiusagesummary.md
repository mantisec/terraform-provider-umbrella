---
page_title: "umbrella_getapiusagesummary Data Source - terraform-provider-umbrella"
description: |-
  Get the total number API requests, and the counts of the successful and failed API requests within a specific time period.
---

# umbrella_getapiusagesummary (Data Source)

Get the total number API requests, and the counts of the successful and failed API requests within a specific time period.

## Example Usage


### Basic Usage

Basic usage of the getapiusagesummary data source

```terraform
data "umbrella_getapiusagesummary" "example" {
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
- **`data`** (String) - The total number of requsts and the counts of the failed and successful requests.
- **`failedRequests`** (Number) - The number of failed API requests.
- **`total`** (Number) - The total number of API requests.
- **`successfulRequests`** (Number) - The number of successful API requests.



