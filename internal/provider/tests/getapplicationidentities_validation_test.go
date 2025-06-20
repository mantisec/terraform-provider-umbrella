package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapplicationidentitiesSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapplicationidentitiesConfig_getapplicationidentities_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapplicationidentities.test", "getapplicationidentities", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapplicationidentitiesConfig_getapplicationidentities_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapplicationidentities.test", "getapplicationidentities", "another-valid-name"),
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
				Config:      testAccGetapplicationidentitiesConfig_getapplicationidentities_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGetapplicationidentitiesConfig_getapplicationidentities_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccGetapplicationidentitiesConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_getapplicationidentities" "test" {
  name = %s
}

`, value)
}

func testAccGetapplicationidentitiesConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_getapplicationidentities" "test" {
  name = %s
}

`, value)
}
