---
page_title: "umbrella_trialconversions Resource - trialconversions"
description: |-
  Convert your Managed Services License Agreement (MSLA) trial to an MSLA customer.
---

# umbrella_trialconversions (Resource)

Convert your Managed Services License Agreement (MSLA) trial to an MSLA customer.

## Example Usage


### Basic Usage

Basic usage of the trialconversions resource

```hcl
resource "umbrella_trialconversions" "example" {
  # Add required attributes here
  name = "example-trialconversions"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `conversionStatus` (String) The status of the trial conversion.



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_trialconversions.example 12345
```

