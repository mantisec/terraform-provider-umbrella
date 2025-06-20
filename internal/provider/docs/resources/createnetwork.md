---
page_title: "umbrella_createnetwork Resource - terraform-provider-umbrella"
description: |-
  Create a network. Before you can create a network, contact Support to get your IP range verified.
---

# umbrella_createnetwork (Resource)

Create a network. Before you can create a network, contact Support to get your IP range verified.

## Example Usage


### Basic Usage

Basic usage of the createnetwork resource

```terraform
resource "umbrella_createnetwork" "example" {
  name        = "example-network"
  description = "Example network configuration"
  enabled     = true
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`name`** (String) - The name of the network. Example: `"example-name"`
- **`prefixLength`** (Number) - The length of the prefix. Set a prefix length that is greater than 28 and less than 33. Example: `123`
- **`isDynamic`** (Boolean) - Specifies whether the IP is dynamic. Example: `true`


### Optional

- **`ipAddress`** (String) - The IP address of the network. Example: `"example"`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`isVerified`** (Boolean) - Specifies whether the network is verified.
- **`createdAt`** (String) - The date and time (timestamp) when the network was created.
- **`originId`** (Number) - The origin ID.



## Import

umbrella_createnetwork can be imported using the resource ID:

```shell
terraform import umbrella_createnetwork.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

