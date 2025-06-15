---
page_title: "umbrella_destination_list Resource - terraform-provider-umbrella"
subcategory: ""
description: |-
  Manages Umbrella destination lists for policy enforcement.
---

# umbrella_destination_list (Resource)

Manages Umbrella destination lists for policy enforcement. Destination lists are collections of URLs, domains, or CIDR blocks that can be referenced in security policies.

## Example Usage

### Domain Destination List

```terraform
resource "umbrella_destination_list" "blocked_domains" {
  name = "Blocked Domains"
  type = "DOMAIN"
  destinations = [
    "malicious-site.com",
    "phishing-domain.net",
    "suspicious-website.org"
  ]
}
```

### URL Destination List

```terraform
resource "umbrella_destination_list" "allowed_urls" {
  name = "Allowed URLs"
  type = "URL"
  destinations = [
    "https://trusted-site.com/api",
    "https://corporate-portal.example.com",
    "https://secure-service.net/endpoint"
  ]
}
```

### CIDR Destination List

```terraform
resource "umbrella_destination_list" "internal_networks" {
  name = "Internal Network Ranges"
  type = "CIDR"
  destinations = [
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.0.0/16"
  ]
}
```

## Schema

### Required

- `name` (String) - Name of the destination list
- `type` (String) - Type of destinations in the list. Must be one of: `URL`, `DOMAIN`, or `CIDR`

### Optional

- `destinations` (Set of String) - Set of destination entries. The format depends on the type:
  - For `DOMAIN`: Domain names (e.g., `example.com`)
  - For `URL`: Full URLs (e.g., `https://example.com/path`)
  - For `CIDR`: IP address ranges in CIDR notation (e.g., `192.168.1.0/24`)

### Read-Only

- `id` (String) - Unique identifier of the destination list

## Import

Destination lists can be imported using their ID:

```bash
terraform import umbrella_destination_list.example 12345678
```

## Notes

- Destination lists are referenced by name in policy rules
- Changes to destinations within a list will trigger updates to the remote list
- Empty destination lists are allowed and can be populated later
- The maximum number of destinations per list depends on your Umbrella subscription
- Destination validation is performed based on the list type:
  - `DOMAIN` entries must be valid domain names
  - `URL` entries must be valid URLs with protocol
  - `CIDR` entries must be valid IP address ranges in CIDR notation