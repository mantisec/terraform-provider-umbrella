package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTopUrlsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTopUrlsDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_top_urls.test", "name", "test-top_urls"),
				),
			},
		},
	})
}

func testAccTopUrlsDataSourceConfig_basic() string {
	return `

data "umbrella_top_urls" "test" {
  id = "test-id-12345"
}

`
}
