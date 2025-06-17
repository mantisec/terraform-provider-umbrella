package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSampleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSampleDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_sample.test", "name", "test-sample"),
				),
			},
		},
	})
}

func testAccSampleDataSourceConfig_basic() string {
	return `

data "umbrella_sample" "test" {
  id = "test-id-12345"
}

`
}
