package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetprotocolidentitiesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetprotocolidentitiesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getprotocolidentities.test", "name", "test-getprotocolidentities"),
				),
			},
		},
	})
}

func testAccGetprotocolidentitiesDataSourceConfig_basic() string {
	return `

data "umbrella_getprotocolidentities" "test" {
  id = "test-id-12345"
}

`
}
