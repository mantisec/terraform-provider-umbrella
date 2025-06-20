package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetroamingcomputerDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetroamingcomputerDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getroamingcomputer.test", "name", "test-getroamingcomputer"),
				),
			},
		},
	})
}

func testAccGetroamingcomputerDataSourceConfig_basic() string {
	return `

data "umbrella_getroamingcomputer" "test" {
  id = "test-id-12345"
}

`
}
