package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListprovidercustomersSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListprovidercustomersConfig_listprovidercustomers_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listprovidercustomers.test", "listprovidercustomers", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListprovidercustomersConfig_listprovidercustomers_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listprovidercustomers.test", "listprovidercustomers", "another-valid-name"),
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
				Config:      testAccListprovidercustomersConfig_listprovidercustomers_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccListprovidercustomersConfig_listprovidercustomers_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccListprovidercustomersConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listprovidercustomers" "test" {
  name = %s
}

`, value)
}

func testAccListprovidercustomersConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listprovidercustomers" "test" {
  name = %s
}

`, value)
}
