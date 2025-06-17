---
page_title: "umbrella_orginfo Data Source - orginfo"
description: |-
  Get the OrgInfo.json properties for deploying the Cisco Secure Client on user devices in the organization.
The Cisco Secure Client with the Internet Security module requires the OrgInfo.json properties.
---

# umbrella_orginfo (Data Source)

Get the OrgInfo.json properties for deploying the Cisco Secure Client on user devices in the organization.
The Cisco Secure Client with the Internet Security module requires the OrgInfo.json properties.

## Example Usage


### Basic Usage

Basic usage of the orginfo data source

```hcl
data "umbrella_orginfo" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



