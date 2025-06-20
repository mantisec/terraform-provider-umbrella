package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetmanagedprovidercustomerSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetmanagedprovidercustomerConfig_getmanagedprovidercustomer_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getmanagedprovidercustomer.test", "getmanagedprovidercustomer", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetmanagedprovidercustomerConfig_getmanagedprovidercustomer_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getmanagedprovidercustomer.test", "getmanagedprovidercustomer", "another-valid-name"),
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
				Config:      testAccGetmanagedprovidercustomerConfig_getmanagedprovidercustomer_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccGetmanagedprovidercustomerConfig_getmanagedprovidercustomer_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccGetmanagedprovidercustomerConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_getmanagedprovidercustomer" "test" {
  name = %s
}

`, value)
}

func testAccGetmanagedprovidercustomerConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_getmanagedprovidercustomer" "test" {
  name = %s
}

`, value)
}
