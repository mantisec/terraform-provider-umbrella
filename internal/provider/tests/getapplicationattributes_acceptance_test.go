package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapplicationattributesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapplicationattributesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapplicationattributes.test", "name", "test-getapplicationattributes"),
				),
			},
		},
	})
}

func testAccGetapplicationattributesDataSourceConfig_basic() string {
	return `

data "umbrella_getapplicationattributes" "test" {
  id = "test-id-12345"
}

`
}
