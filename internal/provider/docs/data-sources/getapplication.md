---
page_title: "umbrella_getapplication Data Source - terraform-provider-umbrella"
description: |-
  Get an application by ID.
---

# umbrella_getapplication (Data Source)

Get an application by ID.

## Example Usage


### Basic Usage

Basic usage of the getapplication data source

```terraform
data "umbrella_getapplication" "example" {
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
- **`description`** (String) - The description of the app.
- **`label`** (String) - The application label. Valid values are: `unreviewed`, `approved`, `notApproved`, `underAudit`.
- **`category`** (String) - The category applied to the app.
- **`vendor`** (String) - The vendor that owns the app.
- **`sources`** (Set of String) - The list of app sources where the sources are DNS, Web (Secure Web Gateway), and cloud-delivered firewall (CDFW) traffic events. The list can contain one or more of the source types.
- **`name`** (String) - The name of the app.
- **`weightedRisk`** (String) - The risk the app poses to the environment.
- **`appType`** (String) - The type of the app.
- **`url`** (String) - The URL of the app.
- **`identitiesCount`** (Number) - The number of identities.
- **`firstDetected`** (String) - The date and time (ISO 8601 timestamp) when the app was first detected.
- **`lastDetected`** (String) - The date and time (ISO 8601 timestamp) when the app was last detected.



