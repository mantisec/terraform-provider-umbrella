package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatetunnelResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "name", "test-updatetunnel"),
				),
			},

			{
				Config: testAccUpdatetunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "name", "test-updatetunnel-updated"),
				),
			},
		},
	})
}

func TestUpdatetunnelResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "name", "test-updatetunnel"),
				),
			},

			{
				Config: testAccUpdatetunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "name", "test-updatetunnel-updated"),
				),
			},
		},
	})
}

func TestUpdatetunnelResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "name", "test-updatetunnel"),
				),
			},

			{
				Config: testAccUpdatetunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatetunnel.test", "name", "test-updatetunnel-updated"),
				),
			},
		},
	})
}

func testAccUpdatetunnelResourceConfig_basic() string {
	return `

resource "umbrella_updatetunnel" "test" {
  name = "test-updatetunnel"
  
  name = "test-resource"
  
  siteOriginId = 123
  
  networkCIDRs = ["item1", "item2"]
  
  client = {}
  
}

`
}

func testAccUpdatetunnelResourceConfig_update() string {
	return `
resource "umbrella_updatetunnel" "test" {
  name = "test-updatetunnel-updated"
  
  name = "test-resource-updated"
  
  siteOriginId = 456
  
  networkCIDRs = ["updated1", "updated2"]
  
  client = {}
  
}
`
}
