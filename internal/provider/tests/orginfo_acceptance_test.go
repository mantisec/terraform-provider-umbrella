package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrginfoDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOrginfoDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_orginfo.test", "name", "test-orginfo"),
				),
			},
		},
	})
}

func testAccOrginfoDataSourceConfig_basic() string {
	return `

data "umbrella_orginfo" "test" {
  id = "test-id-12345"
}

`
}
