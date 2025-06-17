---
page_title: "umbrella_as_for_ip.json Data Source - as_for_ip.json"
description: |-
  This endpoint provides data about ASN and IP relationships,
showing how IP addresses are related to each other and to the regional registries.
You can find out more about the IP space associated with an AS
and correlate BGP routing information between AS.

When querying an IP to find which AS (Autonomous System), it is helpful
to find associated IP addresses. The AS is part of the BGP routing for that IP.

A valid result returns an array of hash references.
The hash reference contains information about the AS such as the ASN,
the CIDR prefix of the AS, the Internet Registry (RIR) number (0 through 6),
the Description of the AS and the creation date for the AS.
An empty response returns an empty array ([]).
---

# umbrella_as_for_ip.json (Data Source)

This endpoint provides data about ASN and IP relationships,
showing how IP addresses are related to each other and to the regional registries.
You can find out more about the IP space associated with an AS
and correlate BGP routing information between AS.

When querying an IP to find which AS (Autonomous System), it is helpful
to find associated IP addresses. The AS is part of the BGP routing for that IP.

A valid result returns an array of hash references.
The hash reference contains information about the AS such as the ASN,
the CIDR prefix of the AS, the Internet Registry (RIR) number (0 through 6),
the Description of the AS and the creation date for the AS.
An empty response returns an empty array ([]).

## Example Usage


### Basic Usage

Basic usage of the as_for_ip.json data source

```hcl
data "umbrella_as_for_ip.json" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `creation_date` (String) The date when the AS was first created.
- `ir` (Number) The IR number corresponds to one of the 5 Regional Internet Registries (RIR).
| Registry | Number |	Region |
|-----|-----|-----|
| Registry | 1 | AfriNIC: Africa |
| Registry | 2 | APNIC: Asia, Australia, New Zealand, and neighboring countries. |
| Registry | 3 | ARIN: United States, Canada, several parts of the Caribbean region, and Antarctica. |
| Registry | 4 | LACNIC: Latin America and parts of the Caribbean region. |
| Registry | 5 | RIPE NCC: Europe, Russia, the Middle East, and Central Asia. |
| Registry | 0 | Unknown / Not Available |
- `description` (String) Network Owner Description as provided by the network owner.
- `asn` (String) The autonomous system number (ASN) associated with the IP address.
- `cidr` (String) The IP CIDR for the ASN.



