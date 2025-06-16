package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/mantisec/terraform-provider-umbrella/internal/provider"
)

func TestDestinationListResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationListResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destination_list.test", "name", "test-destination-list"),
					resource.TestCheckResourceAttr("umbrella_destination_list.test", "access", "block"),
					resource.TestCheckResourceAttrSet("umbrella_destination_list.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_destination_list.test", "created_at"),
				),
			},
			{
				Config: testAccDestinationListResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destination_list.test", "name", "test-destination-list-updated"),
					resource.TestCheckResourceAttr("umbrella_destination_list.test", "access", "allow"),
				),
			},
		},
	})
}

func TestDestinationListResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationListResourceConfig_basic(),
			},
			{
				ResourceName:      "umbrella_destination_list.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestDestinationListResource_withDestinations(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationListResourceConfig_withDestinations(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destination_list.test", "name", "test-with-destinations"),
					resource.TestCheckResourceAttr("umbrella_destination_list.test", "destinations.#", "3"),
					resource.TestCheckTypeSetElemAttr("umbrella_destination_list.test", "destinations.*", "example.com"),
					resource.TestCheckTypeSetElemAttr("umbrella_destination_list.test", "destinations.*", "malicious-site.com"),
					resource.TestCheckTypeSetElemAttr("umbrella_destination_list.test", "destinations.*", "192.168.1.100"),
				),
			},
		},
	})
}

func testAccDestinationListResourceConfig_basic() string {
	return `
resource "umbrella_destination_list" "test" {
  name   = "test-destination-list"
  access = "block"
}
`
}

func testAccDestinationListResourceConfig_update() string {
	return `
resource "umbrella_destination_list" "test" {
  name   = "test-destination-list-updated"
  access = "allow"
}
`
}

func testAccDestinationListResourceConfig_withDestinations() string {
	return `
resource "umbrella_destination_list" "test" {
  name   = "test-with-destinations"
  access = "block"
  destinations = [
    "example.com",
    "malicious-site.com",
    "192.168.1.100"
  ]
}
`
}

// Schema validation tests
func TestDestinationListSchema_validation(t *testing.T) {
	// Test valid access values
	validAccessValues := []string{"allow", "block"}
	for _, value := range validAccessValues {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccDestinationListConfig_access_valid(value),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("umbrella_destination_list.test", "access", value),
					),
				},
			},
		})
	}

	// Test invalid access values
	invalidAccessValues := []string{"invalid-access-type", "permit", "deny"}
	for _, value := range invalidAccessValues {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      testAccDestinationListConfig_access_invalid(value),
					ExpectError: regexp.MustCompile(".*"),
				},
			},
		})
	}
}

func testAccDestinationListConfig_access_valid(value string) string {
	return fmt.Sprintf(`
resource "umbrella_destination_list" "test" {
  name   = "test-validation"
  access = "%s"
}
`, value)
}

func testAccDestinationListConfig_access_invalid(value string) string {
	return fmt.Sprintf(`
resource "umbrella_destination_list" "test" {
  name   = "test-validation"
  access = "%s"
}
`, value)
}

// Unit tests for schema
func TestDestinationListResource_Schema(t *testing.T) {
	r := provider.NewGeneratedDestinationListResource()

	// Test that the resource implements the correct interface
	var _ tfresource.Resource = r
}

// Mock response tests
func TestDestinationListResource_MockResponses(t *testing.T) {
	// Test with mock server responses
	// This would be implemented with httptest.Server for unit testing
	t.Skip("Mock response tests not implemented yet")
}
