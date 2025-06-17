---
page_title: "umbrella_trialextensions Resource - trialextensions"
description: |-
  Create the extension for the customer's trial.
---

# umbrella_trialextensions (Resource)

Create the extension for the customer's trial.

## Example Usage


### Basic Usage

Basic usage of the trialextensions resource

```hcl
resource "umbrella_trialextensions" "example" {
  # Add required attributes here
  name = "example-trialextensions"
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
terraform import umbrella_trialextensions.example 12345
```

