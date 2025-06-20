---
page_title: "umbrella_createcname Resource - terraform-provider-umbrella"
description: |-
  Add the cname information for the service providers console.
---

# umbrella_createcname (Resource)

Add the cname information for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the createcname resource

```terraform
resource "umbrella_createcname" "example" {
  name        = "example-createcname"
  description = "Example createcname resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`cnameId`** (Number) - The cname ID.
- **`cname`** (String) - The cname for the service providers console.
- **`organizationId`** (Number) - The organization ID.



## Import

umbrella_createcname can be imported using the resource ID:

```shell
terraform import umbrella_createcname.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

