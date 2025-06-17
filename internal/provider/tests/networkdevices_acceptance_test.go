package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworkdevicesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkdevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_networkdevices.test", "name", "test-networkdevices"),
				),
			},

			{
				Config: testAccNetworkdevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networkdevices.test", "name", "test-networkdevices-updated"),
				),
			},
		},
	})
}

func TestNetworkdevicesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkdevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_networkdevices.test", "name", "test-networkdevices"),
				),
			},

			{
				Config: testAccNetworkdevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networkdevices.test", "name", "test-networkdevices-updated"),
				),
			},
		},
	})
}

func TestNetworkdevicesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkdevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_networkdevices.test", "name", "test-networkdevices"),
				),
			},

			{
				Config: testAccNetworkdevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networkdevices.test", "name", "test-networkdevices-updated"),
				),
			},
		},
	})
}

func testAccNetworkdevicesResourceConfig_basic() string {
	return `

resource "umbrella_networkdevices" "test" {
  name = "test-networkdevices"
  
  name = "test-resource"
  
}

`
}

func testAccNetworkdevicesResourceConfig_update() string {
	return `
resource "umbrella_networkdevices" "test" {
  name = "test-networkdevices-updated"
  
  name = "test-resource-updated"
  
}
`
}
