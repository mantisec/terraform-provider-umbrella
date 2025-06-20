package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetcustomeraddressesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetcustomeraddressesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getcustomeraddresses.test", "name", "test-getcustomeraddresses"),
				),
			},
		},
	})
}

func testAccGetcustomeraddressesDataSourceConfig_basic() string {
	return `

data "umbrella_getcustomeraddresses" "test" {
  id = "test-id-12345"
}

`
}
