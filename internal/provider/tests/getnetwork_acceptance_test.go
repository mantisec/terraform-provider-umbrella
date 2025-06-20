package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetnetworkDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetnetworkDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getnetwork.test", "name", "test-getnetwork"),
				),
			},
		},
	})
}

func testAccGetnetworkDataSourceConfig_basic() string {
	return `

data "umbrella_getnetwork" "test" {
  id = "test-id-12345"
}

`
}
