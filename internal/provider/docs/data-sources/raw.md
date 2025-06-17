---
page_title: "umbrella_raw Data Source - raw"
description: |-
  Get the Resource Record (RR) data for DNS responses, and categorization data, where the answer (or rdata) could be anything.
---

# umbrella_raw (Data Source)

Get the Resource Record (RR) data for DNS responses, and categorization data, where the answer (or rdata) could be anything.

## Example Usage


### Basic Usage

Basic usage of the raw data source

```hcl
data "umbrella_raw" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



