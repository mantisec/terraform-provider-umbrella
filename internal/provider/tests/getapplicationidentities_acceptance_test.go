package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapplicationidentitiesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapplicationidentitiesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapplicationidentities.test", "name", "test-getapplicationidentities"),
				),
			},
		},
	})
}

func testAccGetapplicationidentitiesDataSourceConfig_basic() string {
	return `

data "umbrella_getapplicationidentities" "test" {
  id = "test-id-12345"
}

`
}
