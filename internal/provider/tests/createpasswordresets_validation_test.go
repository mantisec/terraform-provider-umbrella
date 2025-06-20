package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatepasswordresetsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatepasswordresetsConfig_createpasswordresets_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "createpasswordresets", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatepasswordresetsConfig_createpasswordresets_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "createpasswordresets", "another-valid-name"),
				),
			},
		},
	})

	// Test invalid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatepasswordresetsConfig_createpasswordresets_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatepasswordresetsConfig_createpasswordresets_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccCreatepasswordresetsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createpasswordresets" "test" {
  name = %s
}

`, value)
}

func testAccCreatepasswordresetsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createpasswordresets" "test" {
  name = %s
}

`, value)
}

func testAccCreatepasswordresetsConfig_adminEmails_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createpasswordresets" "test" {
  adminEmails = %s
}

`, value)
}

func testAccCreatepasswordresetsConfig_adminEmails_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createpasswordresets" "test" {
  adminEmails = %s
}

`, value)
}
