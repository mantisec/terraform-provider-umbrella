package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcnamesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcnamesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcnames.test", "name", "test-getcnames"),
				),
			},
		},
	})
}

func testAccGetcnamesDataSourceConfig_basic() string {
	return `

data "umbrella_getcnames" "test" {
  id = "test-id-12345"
}

`
}
