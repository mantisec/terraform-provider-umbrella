---
page_title: "umbrella_getnetwork Data Source - terraform-provider-umbrella"
description: |-
  Get a Network.
---

# umbrella_getnetwork (Data Source)

Get a Network.

## Example Usage


### Basic Usage

Basic usage of the getnetwork data source

```terraform
data "umbrella_getnetwork" "example" {
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
- **`originId`** (Number) - The origin ID.
- **`name`** (String) - The name of the network.
- **`ipAddress`** (String) - The IP address of the network.
- **`prefixLength`** (Number) - The length of the prefix. Set a prefix length that is greater than 28 and less than 33.
- **`isDynamic`** (Boolean) - Specifies whether the network has a dynamic IP.
- **`isVerified`** (Boolean) - Specifies whether the network is verified.
- **`createdAt`** (String) - The date and time (timestamp) when the network was created.



