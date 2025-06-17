---
page_title: "umbrella_deployment_status Data Source - deployment_status"
description: |-
  List the deployment status within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only
---

# umbrella_deployment_status (Data Source)

List the deployment status within the timeframe.

**Access Scope:** Reports > Granular Events > Read-Only

## Example Usage


### Basic Usage

Basic usage of the deployment_status data source

```hcl
data "umbrella_deployment_status" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `meta` (String) 
- `data` (List of String) 



