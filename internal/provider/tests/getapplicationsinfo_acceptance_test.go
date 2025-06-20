package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapplicationsinfoDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapplicationsinfoDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapplicationsinfo.test", "name", "test-getapplicationsinfo"),
				),
			},
		},
	})
}

func testAccGetapplicationsinfoDataSourceConfig_basic() string {
	return `

data "umbrella_getapplicationsinfo" "test" {
  id = "test-id-12345"
}

`
}
