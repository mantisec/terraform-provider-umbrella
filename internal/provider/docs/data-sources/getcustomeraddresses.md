---
page_title: "umbrella_getcustomeraddresses Data Source - terraform-provider-umbrella"
description: |-
  Get the customer addresses for provider.
---

# umbrella_getcustomeraddresses (Data Source)

Get the customer addresses for provider.

## Example Usage


### Basic Usage

Basic usage of the getcustomeraddresses data source

```terraform
data "umbrella_getcustomeraddresses" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Argument Reference

The following arguments are supported:

### Required



### Optional



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- **`id`** (String) - Resource identifier
- **`organizationName`** (String) - The name of the customer's organization.
- **`streetAddress`** (String) - The street address for the customer.
- **`city`** (String) - The name of the city where the customer's organization is located.
- **`state`** (String) - The name of the customer's state.
- **`countryCode`** (String) - The country code of the customer's organization.
- **`accountId`** (Number) - The ID of the account.
- **`accountSiteId`** (Number) - The ID of the site for the account.
- **`mappedCrPartyId`** (Number) - The ID of the mapped CR party.
- **`zipCode`** (String) - The zip code of the customer's organization.
- **`streetAddress2`** (String) - The second street address for the customer.



