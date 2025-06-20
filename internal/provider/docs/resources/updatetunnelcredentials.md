---
page_title: "umbrella_updatetunnelcredentials Resource - terraform-provider-umbrella"
description: |-
  Update and rotate the tunnel credentials.
---

# umbrella_updatetunnelcredentials (Resource)

Update and rotate the tunnel credentials.

## Example Usage


### Basic Usage

Basic usage of the updatetunnelcredentials resource

```terraform
resource "umbrella_updatetunnelcredentials" "example" {
  name           = "example-tunnel"
  remote_gateway = "192.168.1.1"
  preshared_key  = "your-preshared-key"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`autoRotate`** (Boolean) - Specifies whether to autogenerate keys. The 24-hour window applies. Umbrella ignores any passed in credentials. Example: `true`


### Optional

- **`deprecateCurrentKeys`** (Boolean) - Specifies whether to deprecate any existing credentials. The 24-hour window does not apply. Umbrella deletes the existing keys immediately. Set either both `idPrefix` and `secret` or `autoRotate` to true. The default value is false. Example: `true`
- **`psk`** (String) -  Example: `"example"`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`uri`** (String) - Resource URI
- **`siteOriginId`** (Number) - The Site origin ID that is associated with the tunnel.
- **`createdAt`** (String) - The date and time (timestamp) when the tunnel was created.
- **`networkCIDRs`** (Set of String) - Enter IPv4 ranges and CIDR addresses. If `serviceType` is SIG, add all public and private address ranges used internally by your organization. Overrides Umbrella's default behavior, which allows traffic that is destined for RFC-1918 addresses to return through the tunnel. If `serviceType` is Private Access, this field is required. The 0.0.0.0/0 address range is not allowed.
- **`modifiedAt`** (String) - The data and time (timestamp) when the tunnel was updated.
- **`name`** (String) - Display the name of the tunnel. The tunnel name is required, cannot exceed 50 characters in length, and can't have any special characters other than spaces and hyphens.
- **`client`** (String) - The tunnel client's configuration metadata including the client secret.
- **`transport`** (String) - 
- **`serviceType`** (String) - The type of service to associate with the tunnel. The default value is `SIG`.



## Import

umbrella_updatetunnelcredentials can be imported using the resource ID:

```shell
terraform import umbrella_updatetunnelcredentials.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

