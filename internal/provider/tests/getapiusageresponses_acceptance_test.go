package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapiusageresponsesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapiusageresponsesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapiusageresponses.test", "name", "test-getapiusageresponses"),
				),
			},
		},
	})
}

func testAccGetapiusageresponsesDataSourceConfig_basic() string {
	return `

data "umbrella_getapiusageresponses" "test" {
  id = "test-id-12345"
}

`
}
