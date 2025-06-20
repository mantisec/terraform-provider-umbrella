---
page_title: "umbrella_createdestinations Resource - terraform-provider-umbrella"
description: |-
  Add destinations to a destination list.
---

# umbrella_createdestinations (Resource)

Add destinations to a destination list.

## Example Usage


### Basic Usage

Basic usage of the createdestinations resource

```terraform
resource "umbrella_createdestinations" "example" {
  name        = "example-createdestinations"
  description = "Example createdestinations resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`destination`** (String) - A domain, URL, or IP. Example: `"example"`


### Optional

- **`comment`** (String) - A comment about the destination. Example: `"example"`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`data`** (String) - 
- **`createdAt`** (Number) - The date and time when the destination list was created.
- **`markedForDeletion`** (Boolean) - Specifies whether the destination list is marked for deletion.
- **`bundleTypeId`** (Number) - The type of the destination list in the policy. Set `1` for DNS, `2` for web, and `4` for SAML Bypass. If the field is not specified, the default value is `1`.
- **`organizationId`** (Number) - The organization ID.
- **`isGlobal`** (Boolean) - Specifies whether the destination list is a global destination list. There is only one default `allow` destination list and one default `block` destination list for an organization.
- **`thirdpartyCategoryId`** (Number) - The third-party category ID of the destination list.
- **`modifiedAt`** (Number) - The date and time when the destination list was modified.
- **`isMspDefault`** (Boolean) - Specifies whether MSP is the default.
- **`access`** (String) - The type of access for the destination list.
- **`name`** (String) - The name of the destination list.



## Import

umbrella_createdestinations can be imported using the resource ID:

```shell
terraform import umbrella_createdestinations.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

