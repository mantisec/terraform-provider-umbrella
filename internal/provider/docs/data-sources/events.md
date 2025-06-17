---
page_title: "umbrella_events Data Source - events"
description: |-
  Get the recent tunnel error events.
---

# umbrella_events (Data Source)

Get the recent tunnel error events.

## Example Usage


### Basic Usage

Basic usage of the events data source

```hcl
data "umbrella_events" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `meta` (Object) The tunnel error metadata.
- `data` (List of Object) 



