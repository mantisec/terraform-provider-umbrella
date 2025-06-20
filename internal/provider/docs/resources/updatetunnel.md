---
page_title: "umbrella_updatetunnel Resource - terraform-provider-umbrella"
description: |-
  Update the `name`, `siteOriginId`, `networkCIDRs`, and client `deviceType` properties for a tunnel.
Updates to read-only attributes are ignored.
---

# umbrella_updatetunnel (Resource)

Update the `name`, `siteOriginId`, `networkCIDRs`, and client `deviceType` properties for a tunnel.
Updates to read-only attributes are ignored.

## Example Usage


### Basic Usage

Basic usage of the updatetunnel resource

```terraform
resource "umbrella_updatetunnel" "example" {
  name           = "example-tunnel"
  remote_gateway = "192.168.1.1"
  preshared_key  = "your-preshared-key"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`uri`** (String) - Resource URI
- **`createdAt`** (String) - The time when the tunnel was created.
- **`modifiedAt`** (String) - The data and time (timestamp) when the tunnel was updated.
- **`client`** (String) - The tunnel client's configuration metadata.
- **`transport`** (String) - 
- **`serviceType`** (String) - The type of service to associate with the tunnel. The default value is `SIG`.
- **`networkCIDRs`** (Set of String) - Enter IPv4 ranges and CIDR addresses. If `serviceType` is SIG, add all public and private address ranges used internally by your organization. Overrides Umbrella's default behavior, which allows traffic that is destined for RFC-1918 addresses to return through the tunnel. If `serviceType` is Private Access, this field is required. The 0.0.0.0/0 address range is not allowed.
- **`name`** (String) - The display name of the tunnel. The tunnel name is required, cannot exceed 50 characters in length, and can't have any special characters other than spaces and hyphens.
- **`siteOriginId`** (Number) - The site origin ID, which is associated with the tunnel.



## Import

umbrella_updatetunnel can be imported using the resource ID:

```shell
terraform import umbrella_updatetunnel.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

