package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetinternalnetworkDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetinternalnetworkDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getinternalnetwork.test", "name", "test-getinternalnetwork"),
				),
			},
		},
	})
}

func testAccGetinternalnetworkDataSourceConfig_basic() string {
	return `

data "umbrella_getinternalnetwork" "test" {
  id = "test-id-12345"
}

`
}
