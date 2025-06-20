package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdateinternalnetworkResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateinternalnetwork.test", "name", "test-updateinternalnetwork"),
				),
			},

			{
				Config: testAccUpdateinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateinternalnetwork.test", "name", "test-updateinternalnetwork-updated"),
				),
			},
		},
	})
}

func TestUpdateinternalnetworkResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateinternalnetwork.test", "name", "test-updateinternalnetwork"),
				),
			},

			{
				Config: testAccUpdateinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateinternalnetwork.test", "name", "test-updateinternalnetwork-updated"),
				),
			},
		},
	})
}

func TestUpdateinternalnetworkResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateinternalnetwork.test", "name", "test-updateinternalnetwork"),
				),
			},

			{
				Config: testAccUpdateinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateinternalnetwork.test", "name", "test-updateinternalnetwork-updated"),
				),
			},
		},
	})
}

func testAccUpdateinternalnetworkResourceConfig_basic() string {
	return `

resource "umbrella_updateinternalnetwork" "test" {
  name = "test-updateinternalnetwork"
  
  name = "test-resource"
  
  ipAddress = "test-value"
  
  prefixLength = 123
  
  siteId = 123
  
  networkId = 123
  
  tunnelId = 123
  
}

`
}

func testAccUpdateinternalnetworkResourceConfig_update() string {
	return `
resource "umbrella_updateinternalnetwork" "test" {
  name = "test-updateinternalnetwork-updated"
  
  name = "test-resource-updated"
  
  ipAddress = "updated-value"
  
  prefixLength = 456
  
  siteId = 456
  
  networkId = 456
  
  tunnelId = 456
  
}
`
}
