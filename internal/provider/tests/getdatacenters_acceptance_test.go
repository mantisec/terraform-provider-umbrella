package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetdatacentersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetdatacentersDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getdatacenters.test", "name", "test-getdatacenters"),
				),
			},
		},
	})
}

func testAccGetdatacentersDataSourceConfig_basic() string {
	return `

data "umbrella_getdatacenters" "test" {
  id = "test-id-12345"
}

`
}
