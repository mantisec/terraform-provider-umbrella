---
page_title: "umbrella_listnetworkpolicies Data Source - terraform-provider-umbrella"
description: |-
  List the policies for a network.
---

# umbrella_listnetworkpolicies (Data Source)

List the policies for a network.

## Example Usage


### Basic Usage

Basic usage of the listnetworkpolicies data source

```terraform
data "umbrella_listnetworkpolicies" "example" {
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
- **`name`** (String) - The name of the policy.
- **`type`** (String) - The type of the policy.
- **`organizationId`** (Number) - The organization ID.
- **`isAppliedDirectly`** (Boolean) - Indicates if policy is applied directly to this identity.
- **`uri`** (String) - The resource URI.
- **`priority`** (Number) - A number that represents the priority of the policy in the policy list.
- **`isDefault`** (Boolean) - Specifies whether the policy is the default policy.
- **`createdAt`** (String) - The date and time (ISO 8601 timestamp) when the policy was created.
- **`modifiedAt`** (String) - The date and time (ISO 8601 timestamp) when the policy was modified.



