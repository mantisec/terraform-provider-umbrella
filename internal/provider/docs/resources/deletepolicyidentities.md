---
page_title: "umbrella_deletepolicyidentities Resource - terraform-provider-umbrella"
description: |-
  Remove an identity from an Umbrella policy.
Policy changes may require up to 20 minutes to take effect globally.
For DNS policies, TTLs, caching, and session reuse may cause some devices and domains to appear to take longer to update.
---

# umbrella_deletepolicyidentities (Resource)

Remove an identity from an Umbrella policy.
Policy changes may require up to 20 minutes to take effect globally.
For DNS policies, TTLs, caching, and session reuse may cause some devices and domains to appear to take longer to update.

## Example Usage


### Basic Usage

Basic usage of the deletepolicyidentities resource

```terraform
resource "umbrella_deletepolicyidentities" "example" {
  name        = "example-deletepolicyidentities"
  description = "Example deletepolicyidentities resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`name`** (String) - The name of the resource
- **`description`** (String) - The description of the resource
- **`enabled`** (Boolean) - Whether the resource is enabled
- **`active`** (Boolean) - Whether the resource is active
- **`status`** (String) - The status of the resource
- **`organization_id`** (Number) - The organization ID
- **`created_at`** (Number) - The date and time when the resource was created
- **`modified_at`** (Number) - The date and time when the resource was modified
- **`updated_at`** (Number) - The date and time when the resource was updated
- **`created_by`** (String) - The user who created the resource
- **`modified_by`** (String) - The user who modified the resource



## Import

umbrella_deletepolicyidentities can be imported using the resource ID:

```shell
terraform import umbrella_deletepolicyidentities.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

