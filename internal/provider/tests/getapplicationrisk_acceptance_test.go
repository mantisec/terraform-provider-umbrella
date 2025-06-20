package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetapplicationriskDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetapplicationriskDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getapplicationrisk.test", "name", "test-getapplicationrisk"),
				),
			},
		},
	})
}

func testAccGetapplicationriskDataSourceConfig_basic() string {
	return `

data "umbrella_getapplicationrisk" "test" {
  id = "test-id-12345"
}

`
}
