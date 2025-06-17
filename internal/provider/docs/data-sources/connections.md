---
page_title: "umbrella_connections Data Source - connections"
description: |-
  Information about network activity associated with this sample, such as connections to other domains or IPs.
---

# umbrella_connections (Data Source)

Information about network activity associated with this sample, such as connections to other domains or IPs.

## Example Usage


### Basic Usage

Basic usage of the connections data source

```hcl
data "umbrella_connections" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



