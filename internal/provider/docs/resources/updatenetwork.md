---
page_title: "umbrella_updatenetwork Resource - terraform-provider-umbrella"
description: |-
  Update a Network. Before you can update the network's IP address, contact Support to get your IP range verified.
---

# umbrella_updatenetwork (Resource)

Update a Network. Before you can update the network's IP address, contact Support to get your IP range verified.

## Example Usage


### Basic Usage

Basic usage of the updatenetwork resource

```terraform
resource "umbrella_updatenetwork" "example" {
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
- **`originId`** (Number) - The origin ID.
- **`name`** (String) - The name of the network.
- **`ipAddress`** (String) - The IP address of the network.
- **`prefixLength`** (Number) - The length of the prefix. Set a prefix length that is greater than 28 and less than 33.
- **`isDynamic`** (Boolean) - Specifies whether the network has a dynamic IP.
- **`isVerified`** (Boolean) - Specifies whether the network is verified.
- **`createdAt`** (String) - The date and time (timestamp) when the network was created.



## Import

umbrella_updatenetwork can be imported using the resource ID:

```shell
terraform import umbrella_updatenetwork.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

