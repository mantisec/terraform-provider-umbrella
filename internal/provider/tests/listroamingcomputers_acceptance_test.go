package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListroamingcomputersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListroamingcomputersDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listroamingcomputers.test", "name", "test-listroamingcomputers"),
				),
			},
		},
	})
}

func testAccListroamingcomputersDataSourceConfig_basic() string {
	return `

data "umbrella_listroamingcomputers" "test" {
  id = "test-id-12345"
}

`
}
