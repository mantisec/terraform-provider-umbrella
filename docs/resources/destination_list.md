---
page_title: "umbrella_destination_list Resource - destination_list"
description: |-
  Manages destination lists in Cisco Umbrella
---

# umbrella_destination_list (Resource)

Manages destination lists in Cisco Umbrella. Destination lists allow you to create custom lists of domains, IPs, or URLs that can be used in policies to allow or block access.

## Example Usage

### Basic Usage

Basic usage of the destination_list resource

```hcl
resource "umbrella_destination_list" "example" {
  name = "example-destination-list"
  access = "block"
  destinations = [
    "example.com",
    "malicious-site.com",
    "192.168.1.100"
  ]
}
```

### Advanced Usage

```hcl
resource "umbrella_destination_list" "advanced" {
  name = "advanced-destination-list"
  access = "allow"
  is_global = false
  destinations = [
    "trusted-domain.com",
    "internal-app.company.com",
    "10.0.0.0/8"
  ]
}
```

## Schema

### Required

- `name` (String) The name of the destination list
- `access` (String) Access type for the destination list (allow/block)

### Optional

- `destinations` (Set of String) List of destinations (domains, IPs, URLs)
- `is_global` (Boolean) Whether this is a global destination list

### Read-Only

- `id` (String) The unique identifier for this destination list
- `created_at` (String) Timestamp when the destination list was created
- `modified_at` (String) Timestamp when the destination list was last modified

## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_destination_list.example 12345
```

## API Endpoints

This resource uses the following Cisco Umbrella API endpoints:

- `POST /policies/v2/organizations/{orgId}/destinationlists` - Create destination list
- `GET /policies/v2/organizations/{orgId}/destinationlists/{id}` - Read destination list
- `PUT /policies/v2/organizations/{orgId}/destinationlists/{id}` - Update destination list
- `DELETE /policies/v2/organizations/{orgId}/destinationlists/{id}` - Delete destination list

## Notes

- Destination lists can contain domains, IP addresses, CIDR blocks, or URLs
- The `access` field determines whether destinations in the list are allowed or blocked
- Global destination lists are shared across all policies in the organization
- Changes to destination lists may take a few minutes to propagate across the Umbrella network