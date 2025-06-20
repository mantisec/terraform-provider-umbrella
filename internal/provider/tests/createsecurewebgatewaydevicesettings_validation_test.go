package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatesecurewebgatewaydevicesettingsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "createsecurewebgatewaydevicesettings", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "createsecurewebgatewaydevicesettings", "another-valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_valid("0"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "createsecurewebgatewaydevicesettings", "0"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_valid("1"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "createsecurewebgatewaydevicesettings", "1"),
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
				Config:      testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccCreatesecurewebgatewaydevicesettingsConfig_createsecurewebgatewaydevicesettings_invalid("invalid-enum-value"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccCreatesecurewebgatewaydevicesettingsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  name = %s
}

`, value)
}

func testAccCreatesecurewebgatewaydevicesettingsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  name = %s
}

`, value)
}

func testAccCreatesecurewebgatewaydevicesettingsConfig_originIds_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  originIds = %s
}

`, value)
}

func testAccCreatesecurewebgatewaydevicesettingsConfig_originIds_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  originIds = %s
}

`, value)
}

func testAccCreatesecurewebgatewaydevicesettingsConfig_value_valid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  value = %s
}

`, value)
}

func testAccCreatesecurewebgatewaydevicesettingsConfig_value_invalid(value string) string {
	return fmt.Sprintf(`

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  value = %s
}

`, value)
}
