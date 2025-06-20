---
page_title: "umbrella_listtunnels Data Source - terraform-provider-umbrella"
description: |-
  List the tunnels for an organization.
---

# umbrella_listtunnels (Data Source)

List the tunnels for an organization.

## Example Usage


### Basic Usage

Basic usage of the listtunnels data source

```terraform
data "umbrella_listtunnels" "example" {
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
- **`createdAt`** (String) - The date and n time (UTC time with milliseconds) when the tunnel was created.
- **`transport`** (String) - 
- **`serviceType`** (String) - The type of service to associate with the tunnel. The default value is `SIG`.
- **`networkCIDRs`** (Set of String) - Enter IPv4 ranges and CIDR addresses. If `serviceType` is SIG, add all public and private address ranges used internally by your organization. Overrides Umbrella's default behavior, which allows traffic that is destined for RFC-1918 addresses to return through the tunnel. If `serviceType` is Private Access, this field is required. The 0.0.0.0/0 address range is not allowed.
- **`uri`** (String) - Resource URI
- **`name`** (String) - The display name of the tunnel. The tunnel name is required, cannot exceed 50 characters in length, and cannot have any special characters other than spaces and hyphens.
- **`siteOriginId`** (Number) - Site Origin ID to associate with the tunnel
- **`client`** (String) - The tunnel client's configuration metadata.
- **`modifiedAt`** (String) - The data and time (timestamp) when the tunnel was updated.



