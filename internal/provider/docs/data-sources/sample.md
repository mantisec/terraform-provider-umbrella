---
page_title: "umbrella_sample Data Source - sample"
description: |-
  Gather the information from the /samples endpoint, then pivot using the checksums
of the samples revealed in your initial query.
This pivot can reveal large chunks of new data about the malware being researched.
Returns a variety of data as nested JSON arrays.
The initial results array contains the information about the original sample.
These results are described first and are, in effect, the samples of the sample.

---

# umbrella_sample (Data Source)

Gather the information from the /samples endpoint, then pivot using the checksums
of the samples revealed in your initial query.
This pivot can reveal large chunks of new data about the malware being researched.
Returns a variety of data as nested JSON arrays.
The initial results array contains the information about the original sample.
These results are described first and are, in effect, the samples of the sample.


## Example Usage


### Basic Usage

Basic usage of the sample data source

```hcl
data "umbrella_sample" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



