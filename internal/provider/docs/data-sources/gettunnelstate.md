---
page_title: "umbrella_gettunnelstate Data Source - terraform-provider-umbrella"
description: |-
  Get the tunnel state information.
---

# umbrella_gettunnelstate (Data Source)

Get the tunnel state information.

## Example Usage


### Basic Usage

Basic usage of the gettunnelstate data source

```terraform
data "umbrella_gettunnelstate" "example" {
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
- **`dcDesc`** (String) - The city and country of region of the data center.
- **`ikeState`** (String) - IKE SA State: * CREATED * CONNECTING * ESTABLISHED * PASSIVE * REKEYING * REKEYED * DELETING * DESTROYING
- **`peerPort`** (String) - The port of the remote peer.
- **`ipsec`** (String) - The tunnel IPSec session state.
- **`dc`** (String) - The domain name of the data center.
- **`dcName`** (String) - The name of the data center.
- **`ike`** (String) - The tunnel IKE session state.
- **`localIp`** (String) - The public IP address, which is assigned to an endpoint device (ISR, Viptela).
- **`ipsecState`** (String) - IPSec/Child SA State: * CREATED * ROUTED * INSTALLING * INSTALLED * UPDATING * REKEYING * REKEYED * RETRYING * DELETING * DELETED * DESTROYING
- **`data`** (String) - The state of the network tunnel data plane.
- **`peerIp`** (String) - The remote peer IP.
- **`modifiedAt`** (String) - The date and time (UTC time with milliseconds) when the tunnel's state was last updated.
- **`tunnelId`** (String) - The tunnel ID
- **`peerId`** (String) - The remote peer IKE ID.
- **`bytesIn`** (String) - The number of processed input bytes (tunnel ingress).
- **`idleTimeIn`** (String) - The idle time (seconds since last inbound packet).
- **`packetsOut`** (String) - The number of processed output packets (tunnel egress).
- **`bytesOut`** (String) - The number of processed output bytes (tunnel egress).
- **`idleTimeOut`** (String) - The idle time (seconds since last outbound packet).
- **`initialized`** (String) - The time when the packet and byte counters were initialized to 0.
- **`packetsIn`** (String) - The number of processed input packets (tunnel ingress).



