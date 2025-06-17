---
page_title: "umbrella_history Data Source - history"
description: |-
  Get a standard WHOIS response record for a single domain with available historical
WHOIS data returned in an object. The information displayed varies by registrant.
The default limit for history is 10. You can set another value with the `limit` query parameter.

---

# umbrella_history (Data Source)

Get a standard WHOIS response record for a single domain with available historical
WHOIS data returned in an object. The information displayed varies by registrant.
The default limit for history is 10. You can set another value with the `limit` query parameter.


## Example Usage


### Basic Usage

Basic usage of the history data source

```hcl
data "umbrella_history" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



