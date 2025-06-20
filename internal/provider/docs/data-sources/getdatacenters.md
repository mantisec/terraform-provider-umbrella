---
page_title: "umbrella_getdatacenters Data Source - terraform-provider-umbrella"
description: |-
  List the information about the IPsec-enabled data centers.
The data center information includes the IP address and location details.
---

# umbrella_getdatacenters (Data Source)

List the information about the IPsec-enabled data centers.
The data center information includes the IP address and location details.

## Example Usage


### Basic Usage

Basic usage of the getdatacenters data source

```terraform
data "umbrella_getdatacenters" "example" {
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
- **`continents`** (Set of String) - The list of continents.



