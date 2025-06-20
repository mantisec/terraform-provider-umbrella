package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGettunnelglobalerroreventsSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGettunnelglobalerroreventsConfig_gettunnelglobalerrorevents_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_gettunnelglobalerrorevents.test", "gettunnelglobalerrorevents", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGettunnelglobalerroreventsConfig_gettunnelglobalerrorevents_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_gettunnelglobalerrorevents.test", "gettunnelglobalerrorevents", "another-valid-name"),
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
				Config:      testAccGettunnelglobalerroreventsConfig_gettunnelglobalerrorevents_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGettunnelglobalerroreventsConfig_gettunnelglobalerrorevents_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccGettunnelglobalerroreventsConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_gettunnelglobalerrorevents" "test" {
  name = %s
}

`, value)
}

func testAccGettunnelglobalerroreventsConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_gettunnelglobalerrorevents" "test" {
  name = %s
}

`, value)
}
