---
page_title: "umbrella_getcustomertrialstrength Data Source - terraform-provider-umbrella"
description: |-
  Get the strength of a customer trial for a provider.
---

# umbrella_getcustomertrialstrength (Data Source)

Get the strength of a customer trial for a provider.

## Example Usage


### Basic Usage

Basic usage of the getcustomertrialstrength data source

```terraform
data "umbrella_getcustomertrialstrength" "example" {
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
- **`customerLoggedIn`** (Boolean) - Specify whether the trial customer logged into their organization.
- **`lastLoginDate`** (String) - The date when the customer last logged into their organization.
- **`identitiesCreated`** (Number) - The number of identities that were created.
- **`hasTraffic`** (Boolean) - Specify whether the organization has traffic.
- **`trialStrength`** (String) - The number of features used by the trial. The hyphen (`-`) character indicates no features.



