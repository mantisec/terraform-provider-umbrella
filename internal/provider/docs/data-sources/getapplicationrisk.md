---
page_title: "umbrella_getapplicationrisk Data Source - terraform-provider-umbrella"
description: |-
  Get the risk for the application.
---

# umbrella_getapplicationrisk (Data Source)

Get the risk for the application.

## Example Usage


### Basic Usage

Basic usage of the getapplicationrisk data source

```terraform
data "umbrella_getapplicationrisk" "example" {
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
- **`weightedRisk`** (String) - The risk the app poses to the environment.
- **`usageType`** (String) - The type of usage. Valid values are: `personal`, `corporate` (higher risk), or `indirect` (lower risk, e.g. content delivery network).
- **`vendorCompliance`** (Set of String) - The list of compliance information for the vendor.
- **`financialViability`** (String) - Financial risk to the service provider, based on Dun & Bradstreet's Dynamic Risk Score.
- **`dataStorage`** (String) - The form of the stored data. Valid values are: `noStorage`, `structured`, `unstructured`, or `na`.
- **`name`** (String) - The name of the app.
- **`businessRisk`** (String) - The business risk of the app.
- **`webReputation`** (Number) - Provides accurate conclusions about a given host by tracking a broad set of attributes and using sophisticated security modeling. Powered by Talos Security Intelligence.



