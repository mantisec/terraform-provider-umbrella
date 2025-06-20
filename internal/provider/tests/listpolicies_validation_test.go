package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListpoliciesSchema_validation(t *testing.T) {
	// Test valid configurations

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListpoliciesConfig_listpolicies_valid("valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listpolicies.test", "listpolicies", "valid-name"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListpoliciesConfig_listpolicies_valid("another-valid-name"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listpolicies.test", "listpolicies", "another-valid-name"),
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
				Config:      testAccListpoliciesConfig_listpolicies_invalid(""),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccListpoliciesConfig_listpolicies_invalid("invalid name with spaces"),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})

}

func testAccListpoliciesConfig_name_valid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listpolicies" "test" {
  name = %s
}

`, value)
}

func testAccListpoliciesConfig_name_invalid(value string) string {
	return fmt.Sprintf(`

data "umbrella_listpolicies" "test" {
  name = %s
}

`, value)
}
