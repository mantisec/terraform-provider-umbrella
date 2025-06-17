---
page_title: "umbrella_whois Data Source - whois"
description: |-
  Get the WHOIS information for the specified email addresses,
nameservers, and domains. You can search by multiple email addresses or multiple nameservers.
This documentation outlines the following API endpoints:
email (single and multiple), domain record (current and historical), and nameserver (single and multiple).
In some instances, WHOIS information can be irregular as there are no standards
between domain registrars and large volumes of information can be returned from a query.
As such, both the email and nameserver WHOIS endpoints have a limit of 500 results,
which you can reduce to a smaller set of results.
There is an `offset` parameter that can be leveraged to retrieve the entire set
of domain entries for a given email without any limitation. Only the email parameter
supports this.
You can sort the email parameter by filtering the entries based on the timestamp field.
If a domain, email, or nameserver has no known WHOIS information, Investigate returns `HTTP 404`.
If a domain, email or nameserver does not exist, Investigate returns `HTTP 404`.

---

# umbrella_whois (Data Source)

Get the WHOIS information for the specified email addresses,
nameservers, and domains. You can search by multiple email addresses or multiple nameservers.
This documentation outlines the following API endpoints:
email (single and multiple), domain record (current and historical), and nameserver (single and multiple).
In some instances, WHOIS information can be irregular as there are no standards
between domain registrars and large volumes of information can be returned from a query.
As such, both the email and nameserver WHOIS endpoints have a limit of 500 results,
which you can reduce to a smaller set of results.
There is an `offset` parameter that can be leveraged to retrieve the entire set
of domain entries for a given email without any limitation. Only the email parameter
supports this.
You can sort the email parameter by filtering the entries based on the timestamp field.
If a domain, email, or nameserver has no known WHOIS information, Investigate returns `HTTP 404`.
If a domain, email or nameserver does not exist, Investigate returns `HTTP 404`.


## Example Usage


### Basic Usage

Basic usage of the whois data source

```hcl
data "umbrella_whois" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



