package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRefreshapikeyResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRefreshapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_refreshapikey.test", "name", "test-refreshapikey"),
				),
			},

			{
				Config: testAccRefreshapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_refreshapikey.test", "name", "test-refreshapikey-updated"),
				),
			},
		},
	})
}

func TestRefreshapikeyResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRefreshapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_refreshapikey.test", "name", "test-refreshapikey"),
				),
			},

			{
				Config: testAccRefreshapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_refreshapikey.test", "name", "test-refreshapikey-updated"),
				),
			},
		},
	})
}

func TestRefreshapikeyResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRefreshapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_refreshapikey.test", "name", "test-refreshapikey"),
				),
			},

			{
				Config: testAccRefreshapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_refreshapikey.test", "name", "test-refreshapikey-updated"),
				),
			},
		},
	})
}

func testAccRefreshapikeyResourceConfig_basic() string {
	return `

resource "umbrella_refreshapikey" "test" {
  name = "test-refreshapikey"
  
  name = "test-resource"
  
}

`
}

func testAccRefreshapikeyResourceConfig_update() string {
	return `
resource "umbrella_refreshapikey" "test" {
  name = "test-refreshapikey-updated"
  
  name = "test-resource-updated"
  
}
`
}
