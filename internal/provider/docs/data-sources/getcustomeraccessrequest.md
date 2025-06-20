---
page_title: "umbrella_getcustomeraccessrequest Data Source - terraform-provider-umbrella"
description: |-
  Get the access request details for the customer's organization.
---

# umbrella_getcustomeraccessrequest (Data Source)

Get the access request details for the customer's organization.

## Example Usage


### Basic Usage

Basic usage of the getcustomeraccessrequest data source

```terraform
data "umbrella_getcustomeraccessrequest" "example" {
  # Add filter attributes here
  id = "12345"
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



