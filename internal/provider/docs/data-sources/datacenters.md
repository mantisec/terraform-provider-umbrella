---
page_title: "umbrella_datacenters Data Source - datacenters"
description: |-
  List the information about the IPsec-enabled data centers.
The data center information includes the IP address and location details.
---

# umbrella_datacenters (Data Source)

List the information about the IPsec-enabled data centers.
The data center information includes the IP address and location details.

## Example Usage


### Basic Usage

Basic usage of the datacenters data source

```hcl
data "umbrella_datacenters" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `continents` (List of Object) The list of continents.



