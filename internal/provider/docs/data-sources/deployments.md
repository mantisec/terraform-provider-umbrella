---
page_title: "umbrella_deployments Data Source - deployments"
description: |-
  List the summary counts of deployment status for the organization within the timeframe.

**Access Scope:** Reports > Customer > Read-Only
---

# umbrella_deployments (Data Source)

List the summary counts of deployment status for the organization within the timeframe.

**Access Scope:** Reports > Customer > Read-Only

## Example Usage


### Basic Usage

Basic usage of the deployments data source

```hcl
data "umbrella_deployments" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `data` (List of String) 
- `meta` (String) 



