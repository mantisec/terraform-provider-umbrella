package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapiusagekeysDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapiusagekeysDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapiusagekeys.test", "name", "test-getapiusagekeys"),
				),
			},
		},
	})
}

func testAccGetapiusagekeysDataSourceConfig_basic() string {
	return `

data "umbrella_getapiusagekeys" "test" {
  id = "test-id-12345"
}

`
}
