---
page_title: "umbrella_nameservers Data Source - nameservers"
description: |-
  Get WHOIS information for the nameservers. To search by multiple nameservers, provide
a comma-delimited list of domain names for the `nameServerList` query parameter.
For example: `ns1.google.com,ns2.google.com`.
---

# umbrella_nameservers (Data Source)

Get WHOIS information for the nameservers. To search by multiple nameservers, provide
a comma-delimited list of domain names for the `nameServerList` query parameter.
For example: `ns1.google.com,ns2.google.com`.

## Example Usage


### Basic Usage

Basic usage of the nameservers data source

```hcl
data "umbrella_nameservers" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `moreDataAvailable` (String) 
- `limit` (String) 
- `sortField` (String) The field that is used to sort the collection.
- `domains` (List of String) The list of WHOIS nameserver domain information.
- `totalResults` (String) 



