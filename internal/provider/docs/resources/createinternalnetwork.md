---
page_title: "umbrella_createinternalnetwork Resource - terraform-provider-umbrella"
description: |-
  Create an internal network.
---

# umbrella_createinternalnetwork (Resource)

Create an internal network.

## Example Usage


### Basic Usage

Basic usage of the createinternalnetwork resource

```terraform
resource "umbrella_createinternalnetwork" "example" {
  name        = "example-network"
  description = "Example network configuration"
  enabled     = true
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`name`** (String) - The name of the internal network. Example: `"example-name"`
- **`ipAddress`** (String) - The IPv4 address of the internal network. Example: `"example"`
- **`prefixLength`** (Number) - The length of the prefix. The prefix length is from 8 through 32. Example: `123`


### Optional

- **`siteId`** (Number) - The site ID. For DNS policies, specify the ID of the site that is associated with internal network. Provide the value of either the `siteId`, `networkId`, or `tunnelId`. Example: `123`
- **`networkId`** (Number) - The network ID. For Web policies that use proxy chaining, specify the ID of the network, which is associated with the internal network. Provide the value of either the `siteId`, `networkId`, or `tunnelId`. Example: `123`
- **`tunnelId`** (Number) - The ID of the tunnel. For Web policies that use an IPsec tunnel, specify the ID of tunnel, which is associated with the internal network. Provide the value of either the `siteId`, `networkId`, or `tunnelId`. Example: `123`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`networkName`** (String) - The name of the network associated with the internal network.
- **`createdAt`** (String) - The date and time (ISO8601 timestamp) when the internal network was created.
- **`modifiedAt`** (String) - The date and time (ISO8601 timestamp) when the internal network was modified.
- **`originId`** (Number) - The origin ID of the internal network.
- **`siteName`** (String) - The name of the site associated with the internal network.
- **`tunnelName`** (String) - The name of the tunnel associated with the internal network.



## Import

umbrella_createinternalnetwork can be imported using the resource ID:

```shell
terraform import umbrella_createinternalnetwork.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

