package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetorganizationinfoDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetorganizationinfoDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getorganizationinfo.test", "name", "test-getorganizationinfo"),
				),
			},
		},
	})
}

func testAccGetorganizationinfoDataSourceConfig_basic() string {
	return `

data "umbrella_getorganizationinfo" "test" {
  id = "test-id-12345"
}

`
}
