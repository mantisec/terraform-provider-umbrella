---
page_title: "umbrella_updatecustomeraccessrequest Resource - terraform-provider-umbrella"
description: |-
  Update the access request for the customer's organization.
---

# umbrella_updatecustomeraccessrequest (Resource)

Update the access request for the customer's organization.

## Example Usage


### Basic Usage

Basic usage of the updatecustomeraccessrequest resource

```terraform
resource "umbrella_updatecustomeraccessrequest" "example" {
  name        = "example-updatecustomeraccessrequest"
  description = "Example updatecustomeraccessrequest resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`state`** (String) - The current state of the access request.
- **`displayAt`** (Number) - The time when the logo was created. Specify the time in milliseconds.
- **`createdAt`** (Number) - The time when the logo was created. Specify the time in milliseconds.
- **`modifiedAt`** (Number) - The time when the logo was last modified. Specify the time in milliseconds.
- **`organizationName`** (String) - The name of the organization that created the access request.
- **`organizationId`** (Number) - The ID of the organization that created the access request.



## Import

umbrella_updatecustomeraccessrequest can be imported using the resource ID:

```shell
terraform import umbrella_updatecustomeraccessrequest.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

