---
page_title: "umbrella_creates3bucketkey Resource - terraform-provider-umbrella"
description: |-
  Rotate the Cisco-managed S3 bucket key for the organization.
---

# umbrella_creates3bucketkey (Resource)

Rotate the Cisco-managed S3 bucket key for the organization.

## Example Usage


### Basic Usage

Basic usage of the creates3bucketkey resource

```terraform
resource "umbrella_creates3bucketkey" "example" {
  name        = "example-creates3bucketkey"
  description = "Example creates3bucketkey resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`oldKeyId`** (String) - The previous ID of the Cisco-managed S3 bucket key. Example: `"example"`
- **`currentKeyId`** (String) - The ID of the Cisco-managed S3 bucket key. Example: `"example"`
- **`secretAccessKey`** (String) - The secret for the Cisco-managed S3 bucket key. Example: `"example"`
- **`keyCreationDate`** (String) - The date and time (ISO 8601-formatted timestamp) when the system created the Cisco-managed S3 bucket key. Example: `"example"`


### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier



## Import

umbrella_creates3bucketkey can be imported using the resource ID:

```shell
terraform import umbrella_creates3bucketkey.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

