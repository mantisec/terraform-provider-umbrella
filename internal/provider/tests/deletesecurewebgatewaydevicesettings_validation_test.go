package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletesecurewebgatewaydevicesettingsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesecurewebgatewaydevicesettingsConfig_deletesecurewebgatewaydevicesettings_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "deletesecurewebgatewaydevicesettings", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesecurewebgatewaydevicesettingsConfig_deletesecurewebgatewaydevicesettings_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "deletesecurewebgatewaydevicesettings", "another-valid-name"),
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
				Config:      testAccDeletesecurewebgatewaydevicesettingsConfig_deletesecurewebgatewaydevicesettings_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDeletesecurewebgatewaydevicesettingsConfig_deletesecurewebgatewaydevicesettings_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccDeletesecurewebgatewaydevicesettingsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_deletesecurewebgatewaydevicesettings" "test" {
  name = %s
}

`, value)
}

func testAccDeletesecurewebgatewaydevicesettingsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_deletesecurewebgatewaydevicesettings" "test" {
  name = %s
}

`, value)
}

func testAccDeletesecurewebgatewaydevicesettingsConfig_originIds_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_deletesecurewebgatewaydevicesettings" "test" {
  originIds = %s
}

`, value)
}

func testAccDeletesecurewebgatewaydevicesettingsConfig_originIds_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_deletesecurewebgatewaydevicesettings" "test" {
  originIds = %s
}

`, value)
}
