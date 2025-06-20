---
page_title: "umbrella_listinternalnetworkpolicies Data Source - terraform-provider-umbrella"
description: |-
  List the policies for an internal network.
---

# umbrella_listinternalnetworkpolicies (Data Source)

List the policies for an internal network.

## Example Usage


### Basic Usage

Basic usage of the listinternalnetworkpolicies data source

```terraform
data "umbrella_listinternalnetworkpolicies" "example" {
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
- **`modifiedAt`** (String) - The date and time (ISO8601 timestamp) when the policy was modified.
- **`uri`** (String) - The resource URI.
- **`name`** (String) - The name that is given to the policy.
- **`type`** (String) - The type of the policy.
- **`organizationId`** (Number) - The Umbrella organization ID.
- **`priority`** (Number) - A number that represents the priority of the policy in the policy list.
- **`isAppliedDirectly`** (Boolean) - Indicates if policy is directly applied to this identity.
- **`isDefault`** (Boolean) - Specifies whether the policy is the default policy.
- **`createdAt`** (String) - The date and time (ISO8601 timestamp) when the policy was created.



