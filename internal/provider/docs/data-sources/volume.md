---
page_title: "umbrella_volume Data Source - volume"
description: |-
  List the query volume for a domain over the last 30 days.
If there is no information about the domain, Investigate returns an empty array.
As the query takes time to generate, the last two hours may be blank.

---

# umbrella_volume (Data Source)

List the query volume for a domain over the last 30 days.
If there is no information about the domain, Investigate returns an empty array.
As the query takes time to generate, the last two hours may be blank.


## Example Usage


### Basic Usage

Basic usage of the volume data source

```hcl
data "umbrella_volume" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `dates` (List of String) The list of dates recorded for the domain.
- `queries` (List of String) The list of the numbers of DNS queries requested for the domain in one hour, listed in ascending order.



