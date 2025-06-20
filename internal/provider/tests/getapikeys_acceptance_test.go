package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapikeysDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapikeysDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapikeys.test", "name", "test-getapikeys"),
				),
			},
		},
	})
}

func testAccGetapikeysDataSourceConfig_basic() string {
	return `

data "umbrella_getapikeys" "test" {
  id = "test-id-12345"
}

`
}
