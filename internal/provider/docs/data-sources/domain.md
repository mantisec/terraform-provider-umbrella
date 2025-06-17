---
page_title: "umbrella_domain Data Source - domain"
description: |-
  Get the Resource Record (RR) data for DNS responses, and categorization data, where the answer (or rdata) is the domain(s).
---

# umbrella_domain (Data Source)

Get the Resource Record (RR) data for DNS responses, and categorization data, where the answer (or rdata) is the domain(s).

## Example Usage


### Basic Usage

Basic usage of the domain data source

```hcl
data "umbrella_domain" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



