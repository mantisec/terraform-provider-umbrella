---
page_title: "umbrella_getprotocol Data Source - terraform-provider-umbrella"
description: |-
  Get protocol by ID.
---

# umbrella_getprotocol (Data Source)

Get protocol by ID.

## Example Usage


### Basic Usage

Basic usage of the getprotocol data source

```terraform
data "umbrella_getprotocol" "example" {
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
- **`blockedEvents`** (Number) - The number of blocked events for the identities.
- **`firstDetected`** (String) - The date and time (ISO 8601 timestamp) when the protocol was first detected for the identities.
- **`lastDetected`** (String) - The date and time (ISO 8601 timestamp) when the protocol was last detected for the identities.
- **`name`** (String) - The name of the protocol.
- **`description`** (String) - The description of the protocol.
- **`identitiesCount`** (Number) - The number of identities.
- **`events`** (Number) - The number of identity events.



