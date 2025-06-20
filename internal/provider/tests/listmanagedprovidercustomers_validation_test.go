package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListmanagedprovidercustomersSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListmanagedprovidercustomersConfig_listmanagedprovidercustomers_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listmanagedprovidercustomers.test", "listmanagedprovidercustomers", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListmanagedprovidercustomersConfig_listmanagedprovidercustomers_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listmanagedprovidercustomers.test", "listmanagedprovidercustomers", "another-valid-name"),
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
				Config:      testAccListmanagedprovidercustomersConfig_listmanagedprovidercustomers_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccListmanagedprovidercustomersConfig_listmanagedprovidercustomers_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccListmanagedprovidercustomersConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listmanagedprovidercustomers" "test" {
  name = %s
}

`, value)
}

func testAccListmanagedprovidercustomersConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listmanagedprovidercustomers" "test" {
  name = %s
}

`, value)
}
