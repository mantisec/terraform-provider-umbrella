package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListnetworkpoliciesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListnetworkpoliciesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listnetworkpolicies.test", "name", "test-listnetworkpolicies"),
				),
			},
		},
	})
}

func testAccListnetworkpoliciesDataSourceConfig_basic() string {
	return `

data "umbrella_listnetworkpolicies" "test" {
  id = "test-id-12345"
}

`
}
