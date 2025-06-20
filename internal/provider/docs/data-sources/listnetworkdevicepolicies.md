---
page_title: "umbrella_listnetworkdevicepolicies Data Source - terraform-provider-umbrella"
description: |-
  List the Umbrella policies associated with a network device.
If no filters are supplied, Umbrella returns the DNS policies.
---

# umbrella_listnetworkdevicepolicies (Data Source)

List the Umbrella policies associated with a network device.
If no filters are supplied, Umbrella returns the DNS policies.

## Example Usage


### Basic Usage

Basic usage of the listnetworkdevicepolicies data source

```terraform
data "umbrella_listnetworkdevicepolicies" "example" {
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
- **`createdAt`** (String) - The time when the policy was created. Specify an ISO 8601-formatted timestamp.
- **`policyId`** (Number) - The unique ID for the policy.
- **`name`** (String) - The name of the policy.
- **`priority`** (Number) - The priority of the policy.
- **`isAppliedDirectly`** (Boolean) - Specify whether the policy is directly applied to this network device.
- **`isDefault`** (Boolean) - Specify whether the policy is the default policy.



