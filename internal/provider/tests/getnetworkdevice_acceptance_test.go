package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetnetworkdeviceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetnetworkdeviceDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getnetworkdevice.test", "name", "test-getnetworkdevice"),
				),
			},
		},
	})
}

func testAccGetnetworkdeviceDataSourceConfig_basic() string {
	return `

data "umbrella_getnetworkdevice" "test" {
  id = "test-id-12345"
}

`
}
