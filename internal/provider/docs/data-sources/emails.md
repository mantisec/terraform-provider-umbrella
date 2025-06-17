---
page_title: "umbrella_emails Data Source - emails"
description: |-
  Get the email address or addresses of the registrar for the domain or domains. The results include
the total number of results for domains registered by this email address and a list
of the first 500 domains associated with this email.
You can pivot on the email address to find other malicious domains registered by the same email.
This endpoint is limited to a maximum of 500 results, which are the first 500 gathered from the database.
Reduce the number of results by setting the `limit` query parameter.
**Note:** Due to the sample length, Investigate may truncate a sample.
---

# umbrella_emails (Data Source)

Get the email address or addresses of the registrar for the domain or domains. The results include
the total number of results for domains registered by this email address and a list
of the first 500 domains associated with this email.
You can pivot on the email address to find other malicious domains registered by the same email.
This endpoint is limited to a maximum of 500 results, which are the first 500 gathered from the database.
Reduce the number of results by setting the `limit` query parameter.
**Note:** Due to the sample length, Investigate may truncate a sample.

## Example Usage


### Basic Usage

Basic usage of the emails data source

```hcl
data "umbrella_emails" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource
- `offset` (String) 
- `moreDataAvailable` (Boolean) Specifies whether there is more than 500 results for this email.
- `limit` (Number) The number of results returned in the response. The default
limit is 500.
- `sortField` (String) The field that is used to sort the collection.
- `domains` (List of String) The list of domains registered by this email and if the domain is currently registered
by this email address.
- `totalResults` (Number) The total number of results for this email address.



