---
page_title: "umbrella_listinternalnetworks Data Source - terraform-provider-umbrella"
description: |-
  List the internal networks.
---

# umbrella_listinternalnetworks (Data Source)

List the internal networks.

## Example Usage


### Basic Usage

Basic usage of the listinternalnetworks data source

```terraform
data "umbrella_listinternalnetworks" "example" {
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
- **`modifiedAt`** (String) - The date and time (ISO8601 timestamp) when the internal network was modified.
- **`originId`** (Number) - The origin ID of the internal network.
- **`prefixLength`** (Number) - The prefix length of the internal network. The prefix length is greater than 8 and no more than 32.
- **`siteName`** (String) - The name of the site associated with the internal network.
- **`siteId`** (Number) - The ID of the site associated with the internal network.
- **`networkName`** (String) - The name of the network associated with the internal network.
- **`networkId`** (Number) - The ID of the network associated with the internal network.
- **`createdAt`** (String) - The date and time (ISO8601 timestamp) when the internal network was created.
- **`name`** (String) - The name of the internal network.
- **`ipAddress`** (String) - The IPv4 address of the internal network.
- **`tunnelName`** (String) - The name of the tunnel associated with the internal network.
- **`tunnelId`** (Number) - The ID of the tunnel associated with the internal network.



