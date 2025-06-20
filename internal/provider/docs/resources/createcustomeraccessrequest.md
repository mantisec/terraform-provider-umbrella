---
page_title: "umbrella_createcustomeraccessrequest Resource - terraform-provider-umbrella"
description: |-
  Create the access request for the customer's organization.
---

# umbrella_createcustomeraccessrequest (Resource)

Create the access request for the customer's organization.

## Example Usage


### Basic Usage

Basic usage of the createcustomeraccessrequest resource

```terraform
resource "umbrella_createcustomeraccessrequest" "example" {
  name        = "example-createcustomeraccessrequest"
  description = "Example createcustomeraccessrequest resource"
}
```



## Argument Reference

The following arguments are supported:

### Required

- **`state`** (String) - The current state of the access request. Example: `"example"`
- **`organizationName`** (String) - The name of the organization that created the access request. Example: `"example-name"`


### Optional

- **`displayAt`** (Number) - The time when the logo was created. Specify the time in milliseconds. Example: `123`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`organizationId`** (Number) - The ID of the organization that created the access request.
- **`createdAt`** (Number) - The time when the logo was created. Specify the time in milliseconds.
- **`modifiedAt`** (Number) - The time when the logo was last modified. Specify the time in milliseconds.



## Import

umbrella_createcustomeraccessrequest can be imported using the resource ID:

```shell
terraform import umbrella_createcustomeraccessrequest.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

