package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGetprotocolDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGetprotocolDataSourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckDataSourceAttr("umbrella_getprotocol.test", "name", "test-getprotocol"),
				),
			},
		},
	})
}

func testAccGetprotocolDataSourceConfig_basic() string {
	return `

data "umbrella_getprotocol" "test" {
  id = "test-id-12345"
}

`
}
