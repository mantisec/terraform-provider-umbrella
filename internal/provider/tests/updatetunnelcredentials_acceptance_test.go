package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatetunnelcredentialsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "name", "test-updatetunnelcredentials"),
				),
			},

			{
				Config: testAccUpdatetunnelcredentialsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "name", "test-updatetunnelcredentials-updated"),
				),
			},
		},
	})
}

func TestUpdatetunnelcredentialsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "name", "test-updatetunnelcredentials"),
				),
			},

			{
				Config: testAccUpdatetunnelcredentialsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "name", "test-updatetunnelcredentials-updated"),
				),
			},
		},
	})
}

func TestUpdatetunnelcredentialsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatetunnelcredentialsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "name", "test-updatetunnelcredentials"),
				),
			},

			{
				Config: testAccUpdatetunnelcredentialsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatetunnelcredentials.test", "name", "test-updatetunnelcredentials-updated"),
				),
			},
		},
	})
}

func testAccUpdatetunnelcredentialsResourceConfig_basic() string {
	return `

resource "umbrella_updatetunnelcredentials" "test" {
  name = "test-updatetunnelcredentials"
  
  name = "test-resource"
  
  deprecateCurrentKeys = true
  
  autoRotate = true
  
  psk = {}
  
}

`
}

func testAccUpdatetunnelcredentialsResourceConfig_update() string {
	return `
resource "umbrella_updatetunnelcredentials" "test" {
  name = "test-updatetunnelcredentials-updated"
  
  name = "test-resource-updated"
  
  deprecateCurrentKeys = false
  
  autoRotate = false
  
  psk = {}
  
}
`
}
