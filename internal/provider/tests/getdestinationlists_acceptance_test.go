package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetdestinationlistsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetdestinationlistsDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getdestinationlists.test", "name", "test-getdestinationlists"),
				),
			},
		},
	})
}

func testAccGetdestinationlistsDataSourceConfig_basic() string {
	return `

data "umbrella_getdestinationlists" "test" {
  id = "test-id-12345"
}

`
}
