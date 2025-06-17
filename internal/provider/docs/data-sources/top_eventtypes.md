---
page_title: "umbrella_top_eventtypes Data Source - top_eventtypes"
description: |-
  List the top event types by the number of requests made for each type of event.
Order the number of requests in descending order.
The valid event types are: `domain_security`, `domain_integration`,
`url_security`, `url_integration`, `cisco_amp`, and `antivirus`.

**Access Scope:** Reports > Aggregations > Read-Only
---

# umbrella_top_eventtypes (Data Source)

List the top event types by the number of requests made for each type of event.
Order the number of requests in descending order.
The valid event types are: `domain_security`, `domain_integration`,
`url_security`, `url_integration`, `cisco_amp`, and `antivirus`.

**Access Scope:** Reports > Aggregations > Read-Only

## Example Usage


### Basic Usage

Basic usage of the top_eventtypes data source

```hcl
data "umbrella_top_eventtypes" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `data` (List of String) 
- `meta` (String) 



