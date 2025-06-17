---
page_title: "umbrella_organizations Data Source - organizations"
description: |-
  Get information about your provider organizations.
To query the collection, provide an email address for a member of an Umbrella provider organization.
---

# umbrella_organizations (Data Source)

Get information about your provider organizations.
To query the collection, provide an email address for a member of an Umbrella provider organization.

## Example Usage


### Basic Usage

Basic usage of the organizations data source

```hcl
data "umbrella_organizations" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



