package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListinternalnetworksDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListinternalnetworksDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_listinternalnetworks.test", "name", "test-listinternalnetworks"),
				),
			},
		},
	})
}

func testAccListinternalnetworksDataSourceConfig_basic() string {
	return `

data "umbrella_listinternalnetworks" "test" {
  id = "test-id-12345"
}

`
}
