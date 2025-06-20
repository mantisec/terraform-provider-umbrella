package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListinternalnetworkpoliciesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListinternalnetworkpoliciesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listinternalnetworkpolicies.test", "name", "test-listinternalnetworkpolicies"),
				),
			},
		},
	})
}

func testAccListinternalnetworkpoliciesDataSourceConfig_basic() string {
	return `

data "umbrella_listinternalnetworkpolicies" "test" {
  id = "test-id-12345"
}

`
}
