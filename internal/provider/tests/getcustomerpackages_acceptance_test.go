package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcustomerpackagesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomerpackagesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomerpackages.test", "name", "test-getcustomerpackages"),
				),
			},
		},
	})
}

func testAccGetcustomerpackagesDataSourceConfig_basic() string {
	return `

data "umbrella_getcustomerpackages" "test" {
  id = "test-id-12345"
}

`
}
