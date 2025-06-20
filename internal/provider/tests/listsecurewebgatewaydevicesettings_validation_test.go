package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListsecurewebgatewaydevicesettingsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListsecurewebgatewaydevicesettingsConfig_listsecurewebgatewaydevicesettings_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "listsecurewebgatewaydevicesettings", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListsecurewebgatewaydevicesettingsConfig_listsecurewebgatewaydevicesettings_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "listsecurewebgatewaydevicesettings", "another-valid-name"),
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
				Config:      testAccListsecurewebgatewaydevicesettingsConfig_listsecurewebgatewaydevicesettings_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccListsecurewebgatewaydevicesettingsConfig_listsecurewebgatewaydevicesettings_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccListsecurewebgatewaydevicesettingsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_listsecurewebgatewaydevicesettings" "test" {
  name = %s
}

`, value)
}

func testAccListsecurewebgatewaydevicesettingsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_listsecurewebgatewaydevicesettings" "test" {
  name = %s
}

`, value)
}

func testAccListsecurewebgatewaydevicesettingsConfig_originIds_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_listsecurewebgatewaydevicesettings" "test" {
  originIds = %s
}

`, value)
}

func testAccListsecurewebgatewaydevicesettingsConfig_originIds_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_listsecurewebgatewaydevicesettings" "test" {
  originIds = %s
}

`, value)
}
