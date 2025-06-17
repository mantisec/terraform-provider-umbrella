---
page_title: "umbrella_artifacts Data Source - artifacts"
description: |-
  Other samples associated with this sample. The sample data does not include a threat score.
Artifacts are only available for Cisco Secure Malware Analytics customers.
---

# umbrella_artifacts (Data Source)

Other samples associated with this sample. The sample data does not include a threat score.
Artifacts are only available for Cisco Secure Malware Analytics customers.

## Example Usage


### Basic Usage

Basic usage of the artifacts data source

```hcl
data "umbrella_artifacts" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



