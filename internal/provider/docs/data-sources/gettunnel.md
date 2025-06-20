---
page_title: "umbrella_gettunnel Data Source - terraform-provider-umbrella"
description: |-
  Get a specific tunnel.
---

# umbrella_gettunnel (Data Source)

Get a specific tunnel.

## Example Usage


### Basic Usage

Basic usage of the gettunnel data source

```terraform
data "umbrella_gettunnel" "example" {
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
- **`modifiedAt`** (String) - The data and time (timestamp) when the tunnel was updated.
- **`uri`** (String) - Resource URI
- **`createdAt`** (String) - The time when the tunnel was created.
- **`siteOriginId`** (Number) - The site origin ID, which is associated with the tunnel.
- **`client`** (String) - The tunnel client's configuration metadata.
- **`transport`** (String) - 
- **`serviceType`** (String) - The type of service to associate with the tunnel. The default value is `SIG`.
- **`networkCIDRs`** (Set of String) - Enter IPv4 ranges and CIDR addresses. If `serviceType` is SIG, add all public and private address ranges used internally by your organization. Overrides Umbrella's default behavior, which allows traffic that is destined for RFC-1918 addresses to return through the tunnel. If `serviceType` is Private Access, this field is required. The 0.0.0.0/0 address range is not allowed.
- **`name`** (String) - The display name of the tunnel. The tunnel name is required, cannot exceed 50 characters in length, and can't have any special characters other than spaces and hyphens.



