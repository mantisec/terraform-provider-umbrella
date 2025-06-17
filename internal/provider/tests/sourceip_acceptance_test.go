package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSourceipDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceipDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_sourceip.test", "name", "test-sourceip"),
				),
			},
		},
	})
}

func testAccSourceipDataSourceConfig_basic() string {
	return `

data "umbrella_sourceip" "test" {
  id = "test-id-12345"
}

`
}
