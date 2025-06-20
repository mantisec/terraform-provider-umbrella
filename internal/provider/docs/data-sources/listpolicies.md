---
page_title: "umbrella_listpolicies Data Source - terraform-provider-umbrella"
description: |-
  List the Umbrella policies. You can filter by policy type.
If you do not specify a policy type, Umbrella returns the DNS policies.
---

# umbrella_listpolicies (Data Source)

List the Umbrella policies. You can filter by policy type.
If you do not specify a policy type, Umbrella returns the DNS policies.

## Example Usage


### Basic Usage

Basic usage of the listpolicies data source

```terraform
data "umbrella_listpolicies" "example" {
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
- **`policyId`** (Number) - The unique policy ID.
- **`organizationId`** (Number) - 
- **`name`** (String) - The label for the policy.
- **`priority`** (Number) - The priority of the policy.
- **`createdAt`** (String) - The time and date (ISO 8601-formatted timestamp) when the policy was created.
- **`isDefault`** (Boolean) - Specified whether the policy is the default.



