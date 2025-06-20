package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetallnetworkdevicesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetallnetworkdevicesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getallnetworkdevices.test", "name", "test-getallnetworkdevices"),
				),
			},
		},
	})
}

func testAccGetallnetworkdevicesDataSourceConfig_basic() string {
	return `

data "umbrella_getallnetworkdevices" "test" {
  id = "test-id-12345"
}

`
}
