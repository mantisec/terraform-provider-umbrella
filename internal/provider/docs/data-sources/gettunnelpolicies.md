---
page_title: "umbrella_gettunnelpolicies Data Source - terraform-provider-umbrella"
description: |-
  List the policies that include the network tunnel.
---

# umbrella_gettunnelpolicies (Data Source)

List the policies that include the network tunnel.

## Example Usage


### Basic Usage

Basic usage of the gettunnelpolicies data source

```terraform
data "umbrella_gettunnelpolicies" "example" {
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
- **`isAppliedDirectly`** (Boolean) - True if the tunnel is applied directly to this policy, false if the policy is configured to use all tunnels.
- **`modifiedAt`** (String) - The data and time (timestamp) when the tunnel was updated.
- **`uri`** (String) - Resource URI
- **`isDefault`** (Boolean) - Indicates whether the policy is the default policy.
- **`type`** (String) - The type of policy.
- **`name`** (String) - The name of the policy.
- **`organizationId`** (Number) - The organization ID.
- **`priority`** (Number) - An integer that represents the position of the policy in the policy list.
- **`createdAt`** (String) - The date and time (timestamp) when the tunnel was created.



