---
page_title: "umbrella_remove Resource - remove"
description: |-
  Remove destinations from the destination list.

**Note:** Accepts a list that contains no more than 500 destination IDs.

You can retrieve a list of the destinations in the destination list through the GET `/destinationlists/{destinationListId}/destinations` operation.
Then, to remove destinations in a destination list, provide a list of destination IDs in the request body of the
DELETE `/destinationlists/{destinationListId}/destinations/remove` operation.
---

# umbrella_remove (Resource)

Remove destinations from the destination list.

**Note:** Accepts a list that contains no more than 500 destination IDs.

You can retrieve a list of the destinations in the destination list through the GET `/destinationlists/{destinationListId}/destinations` operation.
Then, to remove destinations in a destination list, provide a list of destination IDs in the request body of the
DELETE `/destinationlists/{destinationListId}/destinations/remove` operation.

## Example Usage


### Basic Usage

Basic usage of the remove resource

```hcl
resource "umbrella_remove" "example" {
  # Add required attributes here
  name = "example-remove"
}
```



## Schema

### Required



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



## Import

Import is supported using the following syntax:

```shell
terraform import umbrella_remove.example 12345
```

