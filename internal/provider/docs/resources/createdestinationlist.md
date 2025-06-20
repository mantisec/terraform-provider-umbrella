---
page_title: "umbrella_createdestinationlist Resource - terraform-provider-umbrella"
description: |-
  Create a destination list in your organization.
---

# umbrella_createdestinationlist (Resource)

Create a destination list in your organization.

## Example Usage


### Basic Usage

Basic usage of the createdestinationlist resource

```terraform
resource "umbrella_createdestinationlist" "example" {
  name      = "example-destination-list"
  access    = "allow"
  is_global = false
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`access`** (String) - The type of access for the destination list. Example: `"example"`
- **`isGlobal`** (Boolean) - Specifies whether the destination list is a global destination list. There is only one default `allow` destination list and one default `block` destination list for an organization. Example: `true`
- **`name`** (String) - The name of the destination list. Example: `"example-name"`


### Optional

- **`bundleTypeId`** (Number) - The type of the destination list in the policy. Set `1` for DNS, `2` for web, and `4` for SAML Bypass. If the field is not specified, the default value is `1`. Example: `123`
- **`destinations`** (Set of String) - The list of destinations. Example: `["item1", "item2"]`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`data`** (String) - 
- **`modifiedAt`** (Number) - The date and time when the destination list was modified.
- **`isMspDefault`** (Boolean) - Specifies whether MSP is the default.
- **`organizationId`** (Number) - The organization ID.
- **`thirdpartyCategoryId`** (Number) - The third-party category ID of the destination list.
- **`createdAt`** (Number) - The date and time when the destination list was created.
- **`markedForDeletion`** (Boolean) - Specifies whether the destination list is marked for deletion.



## Import

umbrella_createdestinationlist can be imported using the resource ID:

```shell
terraform import umbrella_createdestinationlist.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

