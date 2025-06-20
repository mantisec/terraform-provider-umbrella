package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatenetworkResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetwork.test", "name", "test-createnetwork"),
				),
			},

			{
				Config: testAccCreatenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createnetwork.test", "name", "test-createnetwork-updated"),
				),
			},
		},
	})
}

func TestCreatenetworkResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetwork.test", "name", "test-createnetwork"),
				),
			},

			{
				Config: testAccCreatenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createnetwork.test", "name", "test-createnetwork-updated"),
				),
			},
		},
	})
}

func TestCreatenetworkResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetwork.test", "name", "test-createnetwork"),
				),
			},

			{
				Config: testAccCreatenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createnetwork.test", "name", "test-createnetwork-updated"),
				),
			},
		},
	})
}

func testAccCreatenetworkResourceConfig_basic() string {
	return `

resource "umbrella_createnetwork" "test" {
  name = "test-createnetwork"
  
  name = "test-resource"
  
  ipAddress = "test-value"
  
  prefixLength = 123
  
  isDynamic = true
  
  status = "test-value"
  
}

`
}

func testAccCreatenetworkResourceConfig_update() string {
	return `
resource "umbrella_createnetwork" "test" {
  name = "test-createnetwork-updated"
  
  name = "test-resource-updated"
  
  ipAddress = "updated-value"
  
  prefixLength = 456
  
  isDynamic = false
  
  status = "updated-value"
  
}
`
}
