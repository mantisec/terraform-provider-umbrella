---
page_title: "umbrella_puttrialconversions Resource - terraform-provider-umbrella"
description: |-
  Convert your Managed Services License Agreement (MSLA) trial to an MSLA customer.
---

# umbrella_puttrialconversions (Resource)

Convert your Managed Services License Agreement (MSLA) trial to an MSLA customer.

## Example Usage


### Basic Usage

Basic usage of the puttrialconversions resource

```terraform
resource "umbrella_puttrialconversions" "example" {
  name        = "example-puttrialconversions"
  description = "Example puttrialconversions resource"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`conversionStatus`** (String) - The status of the trial conversion.



## Import

umbrella_puttrialconversions can be imported using the resource ID:

```shell
terraform import umbrella_puttrialconversions.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

