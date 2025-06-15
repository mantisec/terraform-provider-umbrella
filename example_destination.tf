# Example usage of the umbrella_destination resource
# This demonstrates how to manage individual destinations within destination lists

# First, create a destination list
resource "umbrella_destination_list" "example_list" {
  name = "Example Destination List"
  type = "DOMAIN"
}

# Then add individual destinations to the list
resource "umbrella_destination" "example_domain1" {
  destination_list_id = umbrella_destination_list.example_list.id
  destination         = "example.com"
  comment            = "Primary example domain"
}

resource "umbrella_destination" "example_domain2" {
  destination_list_id = umbrella_destination_list.example_list.id
  destination         = "test.example.com"
  comment            = "Test subdomain"
}

resource "umbrella_destination" "example_domain3" {
  destination_list_id = umbrella_destination_list.example_list.id
  destination         = "api.example.com"
  # comment is optional
}

# Example with URL type destination list
resource "umbrella_destination_list" "url_list" {
  name = "URL Destination List"
  type = "URL"
}

resource "umbrella_destination" "example_url" {
  destination_list_id = umbrella_destination_list.url_list.id
  destination         = "https://example.com/api/v1"
  comment            = "API endpoint"
}

# Example with CIDR type destination list
resource "umbrella_destination_list" "cidr_list" {
  name = "CIDR Destination List"
  type = "CIDR"
}

resource "umbrella_destination" "example_cidr" {
  destination_list_id = umbrella_destination_list.cidr_list.id
  destination         = "192.168.1.0/24"
  comment            = "Internal network range"
}