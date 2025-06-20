---
page_title: "umbrella_updatecname Resource - terraform-provider-umbrella"
description: |-
  Update the cname information for the service providers console.
---

# umbrella_updatecname (Resource)

Update the cname information for the service providers console.

## Example Usage


### Basic Usage

Basic usage of the updatecname resource

```terraform
resource "umbrella_updatecname" "example" {
  name        = "example-updatecname"
  description = "Example updatecname resource"
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

umbrella_updatecname can be imported using the resource ID:

```shell
terraform import umbrella_updatecname.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

