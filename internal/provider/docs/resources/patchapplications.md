---
page_title: "umbrella_patchapplications Resource - terraform-provider-umbrella"
description: |-
  Update the labels for the applications.
---

# umbrella_patchapplications (Resource)

Update the labels for the applications.

## Example Usage


### Basic Usage

Basic usage of the patchapplications resource

```terraform
resource "umbrella_patchapplications" "example" {
  name        = "example-patchapplications"
  description = "Example patchapplications resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`number_of_apps`** (Number) - The number of apps updated
- **`timestamp`** (String) - The date and time (ISO 8601 timestamp) of the update.



## Import

umbrella_patchapplications can be imported using the resource ID:

```shell
terraform import umbrella_patchapplications.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

