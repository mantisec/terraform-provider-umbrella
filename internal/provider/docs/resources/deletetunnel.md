---
page_title: "umbrella_deletetunnel Resource - terraform-provider-umbrella"
description: |-
  Delete a tunnel in the organization.
---

# umbrella_deletetunnel (Resource)

Delete a tunnel in the organization.

## Example Usage


### Basic Usage

Basic usage of the deletetunnel resource

```terraform
resource "umbrella_deletetunnel" "example" {
  name           = "example-tunnel"
  remote_gateway = "192.168.1.1"
  preshared_key  = "your-preshared-key"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`message`** (String) - 



## Import

umbrella_deletetunnel can be imported using the resource ID:

```shell
terraform import umbrella_deletetunnel.example 12345
```

**Note:** The resource ID can be found in the Cisco Umbrella dashboard or by using the corresponding data source.

