---
page_title: "umbrella_listroamingcomputers Data Source - terraform-provider-umbrella"
description: |-
  List the roaming computers.
---

# umbrella_listroamingcomputers (Data Source)

List the roaming computers.

## Example Usage


### Basic Usage

Basic usage of the listroamingcomputers data source

```terraform
data "umbrella_listroamingcomputers" "example" {
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
- **`name`** (String) - The name of the roaming computer. `name` is a sequence of 1–50 characters.
- **`lastSyncStatus`** (String) - The status of the last sync on the roaming computer with DNS-layer security.
- **`lastSync`** (String) - The date and time (timestamp) of the last sync.
- **`hasIpBlocking`** (Boolean) - Specifies whether the roaming computer has IP blocking.
- **`version`** (String) - The version of the Cisco Secure Client with the Internet Security module deployed on the roaming computer.
- **`osVersionName`** (String) - The OS version name of the roaming computer.
- **`anyconnectDeviceId`** (String) - The ID of the device that has the Cisco Secure Client deployed with the Internet Security module.
- **`deviceId`** (String) - The hex ID of the roaming computer.
- **`type`** (String) - The type of the roaming computer.
- **`swgStatus`** (String) - The status of the roaming computer with Internet security (Secure Web Gateway).
- **`osVersion`** (String) - The OS version of the roaming computer.
- **`originId`** (Number) - The origin ID for the roaming computer.
- **`lastSyncSwgStatus`** (String) - The status of the last sync on the roaming computer with Internet security (Secure Web Gateway).
- **`appliedBundle`** (Number) - The policy ID.



