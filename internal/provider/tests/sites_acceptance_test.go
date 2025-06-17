package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSitesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSitesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_sites.test", "name", "test-sites"),
				),
			},

			{
				Config: testAccSitesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_sites.test", "name", "test-sites-updated"),
				),
			},
		},
	})
}

func TestSitesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSitesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_sites.test", "name", "test-sites"),
				),
			},

			{
				Config: testAccSitesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_sites.test", "name", "test-sites-updated"),
				),
			},
		},
	})
}

func TestSitesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSitesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_sites.test", "name", "test-sites"),
				),
			},

			{
				Config: testAccSitesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_sites.test", "name", "test-sites-updated"),
				),
			},
		},
	})
}

func testAccSitesResourceConfig_basic() string {
	return `

resource "umbrella_sites" "test" {
  name = "test-sites"
  
  name = "test-resource"
  
}

`
}

func testAccSitesResourceConfig_update() string {
	return `
resource "umbrella_sites" "test" {
  name = "test-sites-updated"
  
  name = "test-resource-updated"
  
}
`
}
