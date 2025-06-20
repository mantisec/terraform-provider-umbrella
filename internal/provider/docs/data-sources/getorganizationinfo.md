---
page_title: "umbrella_getorganizationinfo Data Source - terraform-provider-umbrella"
description: |-
  Get the OrgInfo.json properties for deploying the Cisco Secure Client on user devices in the organization.
The Cisco Secure Client with the Internet Security module requires the OrgInfo.json properties.
---

# umbrella_getorganizationinfo (Data Source)

Get the OrgInfo.json properties for deploying the Cisco Secure Client on user devices in the organization.
The Cisco Secure Client with the Internet Security module requires the OrgInfo.json properties.

## Example Usage


### Basic Usage

Basic usage of the getorganizationinfo data source

```terraform
data "umbrella_getorganizationinfo" "example" {
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
- **`organizationId`** (Number) - The organization ID.
- **`fingerprint`** (String) - A hash that is used to register the Cisco Secure Client on users devices in the organization.
- **`userId`** (Number) - The first 32 bits of the API key ID.



