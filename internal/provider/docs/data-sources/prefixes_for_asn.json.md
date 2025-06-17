---
page_title: "umbrella_prefixes_for_asn.json Data Source - prefixes_for_asn.json"
description: |-
  A response to a valid ASN returns an array of hash references.
Each hash reference contains two keys: `geo` and `cidr`.
Geo is a hash reference with the country name and country code
(the code corresponds to the country code list for ISO-3166-1 alpha-2).
CIDR contains the IP prefix for this ASN.

---

# umbrella_prefixes_for_asn.json (Data Source)

A response to a valid ASN returns an array of hash references.
Each hash reference contains two keys: `geo` and `cidr`.
Geo is a hash reference with the country name and country code
(the code corresponds to the country code list for ISO-3166-1 alpha-2).
CIDR contains the IP prefix for this ASN.


## Example Usage


### Basic Usage

Basic usage of the prefixes_for_asn.json data source

```hcl
data "umbrella_prefixes_for_asn.json" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `cidr` (List of String) A list of the CIDR range of IP addresses associated with this AS.
The CIDR contains the IP prefix for the ASN.
- `geo` (Object) Geo is a hash reference with the country name and country code
(the code corresponds to the country code list for ISO-3166-1 alpha-2).
For more information, see [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).



