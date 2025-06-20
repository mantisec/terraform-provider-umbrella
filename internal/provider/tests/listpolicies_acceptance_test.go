package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListpoliciesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListpoliciesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listpolicies.test", "name", "test-listpolicies"),
				),
			},
		},
	})
}

func testAccListpoliciesDataSourceConfig_basic() string {
	return `

data "umbrella_listpolicies" "test" {
  id = "test-id-12345"
}

`
}
