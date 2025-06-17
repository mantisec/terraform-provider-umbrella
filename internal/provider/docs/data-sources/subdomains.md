---
page_title: "umbrella_subdomains Data Source - subdomains"
description: |-
  Get the subdomains of a given domain. If there is no subdomain for the domain, Investigate returns an empty array.

---

# umbrella_subdomains (Data Source)

Get the subdomains of a given domain. If there is no subdomain for the domain, Investigate returns an empty array.


## Example Usage


### Basic Usage

Basic usage of the subdomains data source

```hcl
data "umbrella_subdomains" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



