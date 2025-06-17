---
page_title: "umbrella_apikeys Data Source - apikeys"
description: |-
  List the API keys created by your organization.
---

# umbrella_apikeys (Data Source)

List the API keys created by your organization.

## Example Usage


### Basic Usage

Basic usage of the apikeys data source

```hcl
data "umbrella_apikeys" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



