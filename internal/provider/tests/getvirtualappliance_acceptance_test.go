package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetvirtualapplianceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetvirtualapplianceDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getvirtualappliance.test", "name", "test-getvirtualappliance"),
				),
			},
		},
	})
}

func testAccGetvirtualapplianceDataSourceConfig_basic() string {
	return `

data "umbrella_getvirtualappliance" "test" {
  id = "test-id-12345"
}

`
}
