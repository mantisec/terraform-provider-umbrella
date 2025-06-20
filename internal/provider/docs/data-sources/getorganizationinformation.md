---
page_title: "umbrella_getorganizationinformation Data Source - terraform-provider-umbrella"
description: |-
  Get information about your provider organizations.
To query the collection, provide an email address for a member of an Umbrella provider organization.
---

# umbrella_getorganizationinformation (Data Source)

Get information about your provider organizations.
To query the collection, provide an email address for a member of an Umbrella provider organization.

## Example Usage


### Basic Usage

Basic usage of the getorganizationinformation data source

```terraform
data "umbrella_getorganizationinformation" "example" {
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
- **`website`** (String) - The URL for the organization.
- **`modifiedAt`** (Number) - The date when the organization was last modified.
- **`hasDelegatedAdmin`** (Boolean) - Specifies whether the organization has assigned an administrator.
- **`organizationName`** (String) - The name of the organization.
- **`mspOrganizationId`** (Number) - The managed service provider (MSP) ID.
- **`organizationTypeId`** (Number) - The type ID of the organization.
- **`createdAt`** (Number) - The date when the organization was created.
- **`originId`** (Number) - The origin ID.
- **`resellerId`** (Number) - The reseller ID.
- **`creatorUserId`** (Number) - The user ID of the creator.
- **`accountManagerUserId`** (Number) - The user ID of the account manager.
- **`salesforceAccountId`** (String) - The Salesforce account ID.



