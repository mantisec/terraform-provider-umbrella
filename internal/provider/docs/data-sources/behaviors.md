---
page_title: "umbrella_behaviors Data Source - behaviors"
description: |-
  Get the information about specific actions or unique properties of this sample,
especially local to your network or the computer where the sample is run.
---

# umbrella_behaviors (Data Source)

Get the information about specific actions or unique properties of this sample,
especially local to your network or the computer where the sample is run.

## Example Usage


### Basic Usage

Basic usage of the behaviors data source

```hcl
data "umbrella_behaviors" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



