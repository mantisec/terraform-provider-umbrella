---
page_title: "umbrella_topmillion Data Source - topmillion"
description: |-
  List the most seen domains in Umbrella. You can download the data in a zip file,
or use the Investigate API to stream the data into a SIEM.
The popularity list contains our most queried domains based on passive DNS usage across
our Umbrella global network of more than 180 billion requests per day with many tens of millions
of unique active users, in more than 165 countries.
The metric does not only consist of browser-based http requests from users but also takes
in to account the number of unique client IPs invoking this domain relative to the sum of all requests
to all domains.
Our popularity ranking reflects the domain's relative internet activity agnostic
to the invocation protocols and applications where as site ranking models (such as Alexa) focus
on the web activity over port 80 (primarily from browsers).
In addition, the Umbrella popularity algorithm also applies data normalization techniques
to smooth potential biases that may occur due to sampling of DNS usage data.

---

# umbrella_topmillion (Data Source)

List the most seen domains in Umbrella. You can download the data in a zip file,
or use the Investigate API to stream the data into a SIEM.
The popularity list contains our most queried domains based on passive DNS usage across
our Umbrella global network of more than 180 billion requests per day with many tens of millions
of unique active users, in more than 165 countries.
The metric does not only consist of browser-based http requests from users but also takes
in to account the number of unique client IPs invoking this domain relative to the sum of all requests
to all domains.
Our popularity ranking reflects the domain's relative internet activity agnostic
to the invocation protocols and applications where as site ranking models (such as Alexa) focus
on the web activity over port 80 (primarily from browsers).
In addition, the Umbrella popularity algorithm also applies data normalization techniques
to smooth potential biases that may occur due to sampling of DNS usage data.


## Example Usage


### Basic Usage

Basic usage of the topmillion data source

```hcl
data "umbrella_topmillion" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



