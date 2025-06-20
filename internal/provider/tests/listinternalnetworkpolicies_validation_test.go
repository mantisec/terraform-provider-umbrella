package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListinternalnetworkpoliciesSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListinternalnetworkpoliciesConfig_listinternalnetworkpolicies_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listinternalnetworkpolicies.test", "listinternalnetworkpolicies", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListinternalnetworkpoliciesConfig_listinternalnetworkpolicies_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listinternalnetworkpolicies.test", "listinternalnetworkpolicies", "another-valid-name"),
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
				Config:      testAccListinternalnetworkpoliciesConfig_listinternalnetworkpolicies_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccListinternalnetworkpoliciesConfig_listinternalnetworkpolicies_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccListinternalnetworkpoliciesConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listinternalnetworkpolicies" "test" {
  name = %s
}

`, value)
}

func testAccListinternalnetworkpoliciesConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listinternalnetworkpolicies" "test" {
  name = %s
}

`, value)
}
