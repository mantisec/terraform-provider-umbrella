---
page_title: "umbrella_samples Data Source - samples"
description: |-
  Specify a domain, IP, or URL. Use the destination to search for all samples
associated with the destination. The default number of items in a response is 10. You can extend the limit.
You must have a license for Cisco Secure Malware Analytics to receive the samples data.

Cisco Secure Malware Analytics retains checksum samples for one year.
You may find that Investigate previously listed a sample related
to a destination. If Cisco Secure Malware Analytics no longer contains a sample
related to the destination, Investigate does not display the sample in the list
of associated samples.

An error may occur when the requested destination is not in a valid format,
if the requested host is not found in our database, or if there is no data available
for the destination that you have requested. CIDR subnets (for example: 10.10.10.0/24) and pattern search is not supported.

---

# umbrella_samples (Data Source)

Specify a domain, IP, or URL. Use the destination to search for all samples
associated with the destination. The default number of items in a response is 10. You can extend the limit.
You must have a license for Cisco Secure Malware Analytics to receive the samples data.

Cisco Secure Malware Analytics retains checksum samples for one year.
You may find that Investigate previously listed a sample related
to a destination. If Cisco Secure Malware Analytics no longer contains a sample
related to the destination, Investigate does not display the sample in the list
of associated samples.

An error may occur when the requested destination is not in a valid format,
if the requested host is not found in our database, or if there is no data available
for the destination that you have requested. CIDR subnets (for example: 10.10.10.0/24) and pattern search is not supported.


## Example Usage


### Basic Usage

Basic usage of the samples data source

```hcl
data "umbrella_samples" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



