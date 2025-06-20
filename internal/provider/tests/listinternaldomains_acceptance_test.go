package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListinternaldomainsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListinternaldomainsDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listinternaldomains.test", "name", "test-listinternaldomains"),
				),
			},
		},
	})
}

func testAccListinternaldomainsDataSourceConfig_basic() string {
	return `

data "umbrella_listinternaldomains" "test" {
  id = "test-id-12345"
}

`
}
