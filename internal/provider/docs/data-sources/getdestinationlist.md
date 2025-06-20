---
page_title: "umbrella_getdestinationlist Data Source - terraform-provider-umbrella"
description: |-
  Get a destination list.
---

# umbrella_getdestinationlist (Data Source)

Get a destination list.

## Example Usage


### Basic Usage

Basic usage of the getdestinationlist data source

```terraform
data "umbrella_getdestinationlist" "example" {
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
- **`data`** (String) - 
- **`organizationId`** (Number) - The organization ID.
- **`isGlobal`** (Boolean) - Specifies whether the destination list is a global destination list. There is only one default `allow` destination list and one default `block` destination list for an organization.
- **`thirdpartyCategoryId`** (Number) - The third-party category ID of the destination list.
- **`createdAt`** (Number) - The date and time when the destination list was created.
- **`markedForDeletion`** (Boolean) - Specifies whether the destination list is marked for deletion.
- **`bundleTypeId`** (Number) - The type of the destination list in the policy. Set `1` for DNS, `2` for web, and `4` for SAML Bypass. If the field is not specified, the default value is `1`.
- **`access`** (String) - The type of access for the destination list.
- **`name`** (String) - The name of the destination list.
- **`modifiedAt`** (Number) - The date and time when the destination list was modified.
- **`isMspDefault`** (Boolean) - Specifies whether MSP is the default.



