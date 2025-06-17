---
page_title: "umbrella_users Resource - users"
description: |-
  Create an Umbrella user account with a designated role.

Once a user account is deleted, you can recreate the account through the `POST` operation.
When you recreate the user account, only set the `email` and `roleId` fields in the Request Body.

For example: {
  "email": "DEVWKS-22@mailinator.com",
  "roleId": 1
}

If you provide all fields for the user account in the Request Body, Umbrella returns an `HTTP/400` (Bad Request) with the
message: `Email already in use`.
---

# umbrella_users (Resource)

Create an Umbrella user account with a designated role.

Once a user account is deleted, you can recreate the account through the `POST` operation.
When you recreate the user account, only set the `email` and `roleId` fields in the Request Body.

For example: {
  "email": "DEVWKS-22@mailinator.com",
  "roleId": 1
}

If you provide all fields for the user account in the Request Body, Umbrella returns an `HTTP/400` (Bad Request) with the
message: `Email already in use`.

## Example Usage


### Basic Usage

Basic usage of the users resource

```hcl
resource "umbrella_users" "example" {
  # Add required attributes here
  name = "example-users"
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
terraform import umbrella_users.example 12345
```

