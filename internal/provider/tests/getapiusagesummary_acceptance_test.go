package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapiusagesummaryDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapiusagesummaryDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapiusagesummary.test", "name", "test-getapiusagesummary"),
				),
			},
		},
	})
}

func testAccGetapiusagesummaryDataSourceConfig_basic() string {
	return `

data "umbrella_getapiusagesummary" "test" {
  id = "test-id-12345"
}

`
}
