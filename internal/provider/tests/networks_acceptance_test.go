package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworksResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworksResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-networks"),
				),
			},

			{
				Config: testAccNetworksResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-networks-updated"),
				),
			},
		},
	})
}

func TestNetworksResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworksResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-networks"),
				),
			},

			{
				Config: testAccNetworksResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-networks-updated"),
				),
			},
		},
	})
}

func TestNetworksResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworksResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-networks"),
				),
			},

			{
				Config: testAccNetworksResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-networks-updated"),
				),
			},
		},
	})
}

func testAccNetworksResourceConfig_basic() string {
	return `

resource "umbrella_networks" "test" {
  name = "test-networks"
  
  name = "test-resource"
  
}

`
}

func testAccNetworksResourceConfig_update() string {
	return `
resource "umbrella_networks" "test" {
  name = "test-networks-updated"
  
  name = "test-resource-updated"
  
}
`
}
