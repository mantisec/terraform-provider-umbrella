---
page_title: "umbrella_patchapikey Resource - terraform-provider-umbrella"
description: |-
  Update the name, description, scopes, and allowed IPs for an API key.
---

# umbrella_patchapikey (Resource)

Update the name, description, scopes, and allowed IPs for an API key.

## Example Usage


### Basic Usage

Basic usage of the patchapikey resource

```terraform
resource "umbrella_patchapikey" "example" {
  name        = "example-patchapikey"
  description = "Example patchapikey resource"
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

umbrella_patchapikey can be imported using the resource ID:

```shell
terraform import umbrella_patchapikey.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

