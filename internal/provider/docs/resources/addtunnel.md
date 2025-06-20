---
page_title: "umbrella_addtunnel Resource - terraform-provider-umbrella"
description: |-
  Add a new tunnel to the organization.
---

# umbrella_addtunnel (Resource)

Add a new tunnel to the organization.

## Example Usage


### Basic Usage

Basic usage of the addtunnel resource

```terraform
resource "umbrella_addtunnel" "example" {
  name           = "example-tunnel"
  remote_gateway = "192.168.1.1"
  preshared_key  = "your-preshared-key"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`name`** (String) - The name of the tunnel. Example: `"example-name"`


### Optional

- **`siteOriginId`** (Number) - The site origin ID to associate with the tunnel. Example: `123`
- **`deviceType`** (String) - The type of device where the tunnel originates. The default value is `other`. Example: `"example"`
- **`serviceType`** (String) - The type of service to associate with the tunnel. The default value is `SIG`. Example: `"example"`
- **`networkCIDRs`** (Set of String) - Enter IPv4 ranges and CIDR addresses. If `serviceType` is SIG, add all public and private address ranges used internally by your organization. Overrides Umbrella's default behavior, which allows traffic that is destined for RFC-1918 addresses to return through the tunnel. If `serviceType` is Private Access, this field is required. The 0.0.0.0/0 address range is not allowed. Example: `["item1", "item2"]`
- **`transport`** (String) -  Example: `"example"`
- **`authentication`** (String) -  Example: `"example"`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`uri`** (String) - Resource URI
- **`createdAt`** (String) - The date and time (timestamp) when the tunnel was created.
- **`client`** (String) - The tunnel client's configuration metadata including the client secret.
- **`modifiedAt`** (String) - The data and time (timestamp) when the tunnel was updated.



## Import

umbrella_addtunnel can be imported using the resource ID:

```shell
terraform import umbrella_addtunnel.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

