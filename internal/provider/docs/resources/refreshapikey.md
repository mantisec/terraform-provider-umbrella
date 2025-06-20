---
page_title: "umbrella_refreshapikey Resource - terraform-provider-umbrella"
description: |-
  Refresh the API key and secret.
---

# umbrella_refreshapikey (Resource)

Refresh the API key and secret.

## Example Usage


### Basic Usage

Basic usage of the refreshapikey resource

```terraform
resource "umbrella_refreshapikey" "example" {
  name        = "example-refreshapikey"
  description = "Example refreshapikey resource"
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

umbrella_refreshapikey can be imported using the resource ID:

```shell
terraform import umbrella_refreshapikey.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

