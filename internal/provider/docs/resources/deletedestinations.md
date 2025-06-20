---
page_title: "umbrella_deletedestinations Resource - terraform-provider-umbrella"
description: |-
  Remove destinations from the destination list.

**Note:** Accepts a list that contains no more than 500 destination IDs.

You can retrieve a list of the destinations in the destination list through the GET `/destinationlists/{destinationListId}/destinations` operation.
Then, to remove destinations in a destination list, provide a list of destination IDs in the request body of the
DELETE `/destinationlists/{destinationListId}/destinations/remove` operation.
---

# umbrella_deletedestinations (Resource)

Remove destinations from the destination list.

**Note:** Accepts a list that contains no more than 500 destination IDs.

You can retrieve a list of the destinations in the destination list through the GET `/destinationlists/{destinationListId}/destinations` operation.
Then, to remove destinations in a destination list, provide a list of destination IDs in the request body of the
DELETE `/destinationlists/{destinationListId}/destinations/remove` operation.

## Example Usage


### Basic Usage

Basic usage of the deletedestinations resource

```terraform
resource "umbrella_deletedestinations" "example" {
  name        = "example-deletedestinations"
  description = "Example deletedestinations resource"
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



## Import

umbrella_deletedestinations can be imported using the resource ID:

```shell
terraform import umbrella_deletedestinations.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

