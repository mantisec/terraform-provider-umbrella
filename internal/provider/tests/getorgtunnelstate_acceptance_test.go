package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetorgtunnelstateDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetorgtunnelstateDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getorgtunnelstate.test", "name", "test-getorgtunnelstate"),
				),
			},
		},
	})
}

func testAccGetorgtunnelstateDataSourceConfig_basic() string {
	return `

data "umbrella_getorgtunnelstate" "test" {
  id = "test-id-12345"
}

`
}
