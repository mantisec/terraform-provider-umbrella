package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCustomeraddressesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomeraddressesDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_customeraddresses.test", "name", "test-customeraddresses"),
				),
			},
		},
	})
}

func testAccCustomeraddressesDataSourceConfig_basic() string {
	return `

data "umbrella_customeraddresses" "test" {
  id = "test-id-12345"
}

`
}
