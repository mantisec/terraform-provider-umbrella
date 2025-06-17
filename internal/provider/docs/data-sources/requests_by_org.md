---
page_title: "umbrella_requests_by_org Data Source - requests_by_org"
description: |-
  List the summary counts of all requests within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
---

# umbrella_requests_by_org (Data Source)

List the summary counts of all requests within the timeframe.

**Access Scope:** Reports > Customer > Read-Only

## Example Usage


### Basic Usage

Basic usage of the requests_by_org data source

```hcl
data "umbrella_requests_by_org" "example" {
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



