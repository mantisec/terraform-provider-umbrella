---
page_title: "umbrella_updateinternalnetwork Resource - terraform-provider-umbrella"
description: |-
  Update an internal network.
---

# umbrella_updateinternalnetwork (Resource)

Update an internal network.

## Example Usage


### Basic Usage

Basic usage of the updateinternalnetwork resource

```terraform
resource "umbrella_updateinternalnetwork" "example" {
  name        = "example-network"
  description = "Example network configuration"
  enabled     = true
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`networkName`** (String) - The name of the network associated with the internal network.
- **`networkId`** (Number) - The ID of the network associated with the internal network.
- **`createdAt`** (String) - The date and time (ISO8601 timestamp) when the internal network was created.
- **`modifiedAt`** (String) - The date and time (ISO8601 timestamp) when the internal network was modified.
- **`originId`** (Number) - The origin ID of the internal network.
- **`prefixLength`** (Number) - The prefix length of the internal network. The prefix length is greater than 8 and no more than 32.
- **`siteName`** (String) - The name of the site associated with the internal network.
- **`siteId`** (Number) - The ID of the site associated with the internal network.
- **`name`** (String) - The name of the internal network.
- **`ipAddress`** (String) - The IPv4 address of the internal network.
- **`tunnelName`** (String) - The name of the tunnel associated with the internal network.
- **`tunnelId`** (Number) - The ID of the tunnel associated with the internal network.



## Import

umbrella_updateinternalnetwork can be imported using the resource ID:

```shell
terraform import umbrella_updateinternalnetwork.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

