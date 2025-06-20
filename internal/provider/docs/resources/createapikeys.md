---
page_title: "umbrella_createapikeys Resource - terraform-provider-umbrella"
description: |-
  Create an API key with certain scopes, name, description, allowed IP addresses and CIDR blocks, and expiration.
The `description` and `allowedIPs` fields are optional.
---

# umbrella_createapikeys (Resource)

Create an API key with certain scopes, name, description, allowed IP addresses and CIDR blocks, and expiration.
The `description` and `allowedIPs` fields are optional.

## Example Usage


### Basic Usage

Basic usage of the createapikeys resource

```terraform
resource "umbrella_createapikeys" "example" {
  name        = "example-createapikeys"
  description = "Example createapikeys resource"
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

umbrella_createapikeys can be imported using the resource ID:

```shell
terraform import umbrella_createapikeys.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

