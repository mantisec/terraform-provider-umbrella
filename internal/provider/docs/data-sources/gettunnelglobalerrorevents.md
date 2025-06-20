---
page_title: "umbrella_gettunnelglobalerrorevents Data Source - terraform-provider-umbrella"
description: |-
  Get the recent global error events.
---

# umbrella_gettunnelglobalerrorevents (Data Source)

Get the recent global error events.

## Example Usage


### Basic Usage

Basic usage of the gettunnelglobalerrorevents data source

```terraform
data "umbrella_gettunnelglobalerrorevents" "example" {
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
- **`data`** (Set of String) - 
- **`peerPort`** (String) - The optional peer port (if available).
- **`time`** (String) - The date and time when the error event was generated (UTC time, with milliseconds).
- **`eventId`** (String) - The unique event ID.
- **`type`** (String) - The event Type: * SSEN - StrongSwan-based Error Notify Events * OTHER - others/future
- **`code`** (String) - The type-specific error codes: * LOCAL_AUTH_FAILED - creating local authentication data failed * PEER_AUTH_FAILED - peer authentication failed * PARSE_ERROR_HEADER - parsing IKE header failed * PARSE_ERROR_BODY - parsing IKE message failed * RETRANSMIT_SEND_TIMEOUT - IKE message retransmission timed out * HALF_OPEN_TIMEOUT - IKE SA timed out before it could be established * PROPOSAL_MISMATCH_IKE - received IKE SA proposals mismatch * PROPOSAL_MISMATCH_CHILD - received CHILD SA proposals mismatch * TS_MISMATCH - received traffic selectors mismatch * INSTALL_CHILD_SA_FAILED - installing IPsec SA failed * INSTALL_CHILD_POLICY_FAILED - installing IPsec policy failed * UNIQUE_REPLACE - replaced old IKE SA due to uniqueness policy * UNIQUE_KEEP - keep existing in favor of rejected new IKE SA due to uniqueness policy * VIP_FAILURE - virtual IP failure * AUTHORIZATION_FAILED - an authorization plugin prevented establishment of an IKE SA * RETRANSMIT_SEND - IKE message retransmission For additional information, see https://datatracker.ietf.org/doc/html/rfc7296.
- **`reason`** (String) - The description of the error.
- **`peerIp`** (String) - The optional peer IP (if available).



