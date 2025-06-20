package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreateinternalnetworkResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createinternalnetwork.test", "name", "test-createinternalnetwork"),
				),
			},

			{
				Config: testAccCreateinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createinternalnetwork.test", "name", "test-createinternalnetwork-updated"),
				),
			},
		},
	})
}

func TestCreateinternalnetworkResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createinternalnetwork.test", "name", "test-createinternalnetwork"),
				),
			},

			{
				Config: testAccCreateinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createinternalnetwork.test", "name", "test-createinternalnetwork-updated"),
				),
			},
		},
	})
}

func TestCreateinternalnetworkResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createinternalnetwork.test", "name", "test-createinternalnetwork"),
				),
			},

			{
				Config: testAccCreateinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createinternalnetwork.test", "name", "test-createinternalnetwork-updated"),
				),
			},
		},
	})
}

func testAccCreateinternalnetworkResourceConfig_basic() string {
	return `

resource "umbrella_createinternalnetwork" "test" {
  name = "test-createinternalnetwork"
  
  name = "test-resource"
  
  ipAddress = "test-value"
  
  prefixLength = 123
  
  siteId = 123
  
  networkId = 123
  
  tunnelId = 123
  
}

`
}

func testAccCreateinternalnetworkResourceConfig_update() string {
	return `
resource "umbrella_createinternalnetwork" "test" {
  name = "test-createinternalnetwork-updated"
  
  name = "test-resource-updated"
  
  ipAddress = "updated-value"
  
  prefixLength = 456
  
  siteId = 456
  
  networkId = 456
  
  tunnelId = 456
  
}
`
}
