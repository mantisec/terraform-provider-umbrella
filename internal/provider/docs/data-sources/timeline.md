---
page_title: "umbrella_timeline Data Source - timeline"
description: |-
  List the historical tagging timeline for a given IP, domain, or URL.
Investigate sorts the timeline items in descending order using the timestamp field.
Each timeline item includes lists of security category,
attack, or threat type associated with the destination.
Use the Tagging Timeline endpoint to verify when Umbrella assigned or removed
a security category, attack, or threat type.
If the current timeline item contains the security category, type of attack,
or threat type not found in the previous timeline item,
Umbrella updated the current timeline item.
If the current timeline item does not contain the security category, attack,
or threat type found in the previous timeline item,
Umbrella removed the security category, type of attack,
or threat type.

---

# umbrella_timeline (Data Source)

List the historical tagging timeline for a given IP, domain, or URL.
Investigate sorts the timeline items in descending order using the timestamp field.
Each timeline item includes lists of security category,
attack, or threat type associated with the destination.
Use the Tagging Timeline endpoint to verify when Umbrella assigned or removed
a security category, attack, or threat type.
If the current timeline item contains the security category, type of attack,
or threat type not found in the previous timeline item,
Umbrella updated the current timeline item.
If the current timeline item does not contain the security category, attack,
or threat type found in the previous timeline item,
Umbrella removed the security category, type of attack,
or threat type.


## Example Usage


### Basic Usage

Basic usage of the timeline data source

```hcl
data "umbrella_timeline" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



