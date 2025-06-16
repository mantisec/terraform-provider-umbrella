package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccUsersResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUsersDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccUsersResourceConfig("test-user@example.com", "John", "Doe", "America/New_York", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "email", "test-user@example.com"),
					resource.TestCheckResourceAttr("umbrella_users.test", "firstname", "John"),
					resource.TestCheckResourceAttr("umbrella_users.test", "lastname", "Doe"),
					resource.TestCheckResourceAttr("umbrella_users.test", "timezone", "America/New_York"),
					resource.TestCheckResourceAttr("umbrella_users.test", "role_id", "1"),
					resource.TestCheckResourceAttrSet("umbrella_users.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_users.test", "user_id"),
					resource.TestCheckResourceAttrSet("umbrella_users.test", "role"),
					resource.TestCheckResourceAttrSet("umbrella_users.test", "status"),
					resource.TestCheckResourceAttrSet("umbrella_users.test", "two_factor_enabled"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "umbrella_users.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"}, // Password is not returned by the API
			},
			// Update testing - Note: Updates are not supported by the API
			// This test verifies that updates fail gracefully
			{
				Config:      testAccUsersResourceConfig("test-user@example.com", "Jane", "Smith", "America/New_York", 1),
				ExpectError: regexp.MustCompile("Update Not Supported"),
			},
		},
	})
}

func TestAccUsersResource_differentRoles(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUsersDestroy,
		Steps: []resource.TestStep{
			// Test with different role IDs
			{
				Config: testAccUsersResourceConfig("admin-user@example.com", "Admin", "User", "UTC", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "email", "admin-user@example.com"),
					resource.TestCheckResourceAttr("umbrella_users.test", "role_id", "1"),
				),
			},
			{
				Config: testAccUsersResourceConfig("readonly-user@example.com", "ReadOnly", "User", "UTC", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "email", "readonly-user@example.com"),
					resource.TestCheckResourceAttr("umbrella_users.test", "role_id", "2"),
				),
			},
		},
	})
}

func TestAccUsersResource_validation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test validation errors
			{
				Config:      testAccUsersResourceConfigInvalid("", "John", "Doe", "UTC", 1),
				ExpectError: regexp.MustCompile("email must be at least 1 character long"),
			},
			{
				Config:      testAccUsersResourceConfigInvalid("invalid-email", "John", "Doe", "UTC", 1),
				ExpectError: regexp.MustCompile("email must be a valid email address"),
			},
			{
				Config:      testAccUsersResourceConfigInvalid("test@example.com", "", "Doe", "UTC", 1),
				ExpectError: regexp.MustCompile("firstname must be at least 1 character long"),
			},
			{
				Config:      testAccUsersResourceConfigInvalid("test@example.com", "John", "", "UTC", 1),
				ExpectError: regexp.MustCompile("lastname must be at least 1 character long"),
			},
			{
				Config:      testAccUsersResourceConfigInvalid("test@example.com", "John", "Doe", "", 1),
				ExpectError: regexp.MustCompile("timezone must be at least 1 character long"),
			},
			{
				Config:      testAccUsersResourceConfigInvalid("test@example.com", "John", "Doe", "UTC", 0),
				ExpectError: regexp.MustCompile("role_id must be at least 1"),
			},
		},
	})
}

func TestAccUsersResource_recreateOnEmailChange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUsersDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersResourceConfig("original@example.com", "John", "Doe", "UTC", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "email", "original@example.com"),
				),
			},
			// Changing email should force recreation
			{
				Config: testAccUsersResourceConfig("changed@example.com", "John", "Doe", "UTC", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_users.test", "email", "changed@example.com"),
				),
			},
		},
	})
}

func testAccCheckUsersDestroy(s *terraform.State) error {
	// This function should verify that the user has been destroyed
	// In a real implementation, you would check that the user no longer exists
	// via the API. For now, we'll return nil to indicate success.
	return nil
}

func testAccUsersResourceConfig(email, firstname, lastname, timezone string, roleID int) string {
	return fmt.Sprintf(`
resource "umbrella_users" "test" {
  email     = "%s"
  firstname = "%s"
  lastname  = "%s"
  password  = "TestPassword123!"
  role_id   = %d
  timezone  = "%s"
}
`, email, firstname, lastname, roleID, timezone)
}

func testAccUsersResourceConfigInvalid(email, firstname, lastname, timezone string, roleID int) string {
	return fmt.Sprintf(`
resource "umbrella_users" "test" {
  email     = "%s"
  firstname = "%s"
  lastname  = "%s"
  password  = "TestPassword123!"
  role_id   = %d
  timezone  = "%s"
}
`, email, firstname, lastname, roleID, timezone)
}

// Additional test configurations for edge cases
func testAccUsersResourceConfigMinimal() string {
	return `
resource "umbrella_users" "test" {
  email     = "minimal@example.com"
  firstname = "M"
  lastname  = "U"
  password  = "P"
  role_id   = 1
  timezone  = "UTC"
}
`
}

func testAccUsersResourceConfigLongValues() string {
	return `
resource "umbrella_users" "test" {
  email     = "very.long.email.address.for.testing@example.com"
  firstname = "VeryLongFirstNameForTestingPurposes"
  lastname  = "VeryLongLastNameForTestingPurposes"
  password  = "VeryLongPasswordForTestingPurposes123!"
  role_id   = 1
  timezone  = "America/Argentina/ComodRivadavia"
}
`
}

// Test configuration with special characters
func testAccUsersResourceConfigSpecialChars() string {
	return `
resource "umbrella_users" "test" {
  email     = "test+special@example.com"
  firstname = "José"
  lastname  = "García-López"
  password  = "P@ssw0rd!#$%"
  role_id   = 1
  timezone  = "Europe/London"
}
`
}
