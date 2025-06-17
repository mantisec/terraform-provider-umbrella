---
page_title: "umbrella_search Data Source - search"
description: |-
  Performs a regular expression (RegEx) search on the WHOIS data (domain, nameserver, and email fields)
that was updated or created in the specified time range.
Returns a list of ten WHOIS records that match the specified RegEx expression.
Use the `offset` query parameter to paginate the collection.
By default, Investigate sorts by the `updated` field.

---

# umbrella_search (Data Source)

Performs a regular expression (RegEx) search on the WHOIS data (domain, nameserver, and email fields)
that was updated or created in the specified time range.
Returns a list of ten WHOIS records that match the specified RegEx expression.
Use the `offset` query parameter to paginate the collection.
By default, Investigate sorts by the `updated` field.


## Example Usage


### Basic Usage

Basic usage of the search data source

```hcl
data "umbrella_search" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `limit` (Number) The total number of results for this page. Default limit is 10.
- `sortField` (String) The field that is used to sort the collection.
- `records` (List of String) The list of WHOIS records.
- `totalResults` (Number) The total number of results for this search.
- `offset` (String) 
- `moreDataAvailable` (Boolean) Specifies whether there is more than 10 results for this search.
- `matches` (List of String) The list of matching records.
- `expression` (String) Specifies the regular expression used in the search.



