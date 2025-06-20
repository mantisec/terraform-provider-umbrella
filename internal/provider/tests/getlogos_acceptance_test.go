package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetlogosDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetlogosDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getlogos.test", "name", "test-getlogos"),
				),
			},
		},
	})
}

func testAccGetlogosDataSourceConfig_basic() string {
	return `

data "umbrella_getlogos" "test" {
  id = "test-id-12345"
}

`
}
