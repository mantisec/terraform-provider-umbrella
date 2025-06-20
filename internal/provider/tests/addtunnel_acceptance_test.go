package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAddtunnelResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "name", "test-addtunnel"),
				),
			},

			{
				Config: testAccAddtunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "name", "test-addtunnel-updated"),
				),
			},
		},
	})
}

func TestAddtunnelResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "name", "test-addtunnel"),
				),
			},

			{
				Config: testAccAddtunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "name", "test-addtunnel-updated"),
				),
			},
		},
	})
}

func TestAddtunnelResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtunnelResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "name", "test-addtunnel"),
				),
			},

			{
				Config: testAccAddtunnelResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_addtunnel.test", "name", "test-addtunnel-updated"),
				),
			},
		},
	})
}

func testAccAddtunnelResourceConfig_basic() string {
	return `

resource "umbrella_addtunnel" "test" {
  name = "test-addtunnel"
  
  name = "test-resource"
  
  siteOriginId = 123
  
  deviceType = "ASA"
  
  serviceType = "SIG"
  
  networkCIDRs = ["item1", "item2"]
  
  transport = {}
  
  authentication = {}
  
}

`
}

func testAccAddtunnelResourceConfig_update() string {
	return `
resource "umbrella_addtunnel" "test" {
  name = "test-addtunnel-updated"
  
  name = "test-resource-updated"
  
  siteOriginId = 456
  
  deviceType = "FTD"
  
  serviceType = "Private Access"
  
  networkCIDRs = ["updated1", "updated2"]
  
  transport = {}
  
  authentication = {}
  
}
`
}
