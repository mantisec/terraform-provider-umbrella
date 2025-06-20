package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetlogoDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetlogoDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getlogo.test", "name", "test-getlogo"),
				),
			},
		},
	})
}

func testAccGetlogoDataSourceConfig_basic() string {
	return `

data "umbrella_getlogo" "test" {
  id = "test-id-12345"
}

`
}
