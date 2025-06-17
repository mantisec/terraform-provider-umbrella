---
page_title: "umbrella_categorization Resource - categorization"
description: |-
  Provide a list of domains and look up the status, and security and content category IDs for each domain.
The domain status is a numerical value determined by the Cisco Security Labs team.
Valid status values are: '-1' (malicious), '1' (safe), or '0' (undetermined status).

---

# umbrella_categorization (Resource)

Provide a list of domains and look up the status, and security and content category IDs for each domain.
The domain status is a numerical value determined by the Cisco Security Labs team.
Valid status values are: '-1' (malicious), '1' (safe), or '0' (undetermined status).


## Example Usage


### Basic Usage

Basic usage of the categorization resource

```hcl
resource "umbrella_categorization" "example" {
  # Add required attributes here
  name = "example-categorization"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_categorization.example 12345
```

