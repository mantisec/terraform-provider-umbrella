package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDatacentersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDatacentersDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_datacenters.test", "name", "test-datacenters"),
				),
			},
		},
	})
}

func testAccDatacentersDataSourceConfig_basic() string {
	return `

data "umbrella_datacenters" "test" {
  id = "test-id-12345"
}

`
}
