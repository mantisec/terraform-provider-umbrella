---
page_title: "umbrella_categorization Data Source - categorization"
description: |-
  Look up the status, and security and content category IDs for the domain.
The domain status is a numerical value determined by the Cisco Security Labs team.
Valid status values are: '-1' (malicious), '1' (safe), or '0' (undetermined status).

---

# umbrella_categorization (Data Source)

Look up the status, and security and content category IDs for the domain.
The domain status is a numerical value determined by the Cisco Security Labs team.
Valid status values are: '-1' (malicious), '1' (safe), or '0' (undetermined status).


## Example Usage


### Basic Usage

Basic usage of the categorization data source

```hcl
data "umbrella_categorization" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



