---
page_title: "umbrella_nameservers Data Source - nameservers"
description: |-
  Get WHOIS information for the nameserver.
As a nameserver can potentially register hundreds or thousands of domains,
the server limits the number of results to 500.
---

# umbrella_nameservers (Data Source)

Get WHOIS information for the nameserver.
As a nameserver can potentially register hundreds or thousands of domains,
the server limits the number of results to 500.

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
- `domains` (List of String) The list of information about the WHOIS emails and nameservers.
- `totalResults` (Number) The total number of WHOIS records found for this query.



