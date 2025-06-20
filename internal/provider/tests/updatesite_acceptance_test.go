package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatesiteResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatesite.test", "name", "test-updatesite"),
				),
			},

			{
				Config: testAccUpdatesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatesite.test", "name", "test-updatesite-updated"),
				),
			},
		},
	})
}

func TestUpdatesiteResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatesite.test", "name", "test-updatesite"),
				),
			},

			{
				Config: testAccUpdatesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatesite.test", "name", "test-updatesite-updated"),
				),
			},
		},
	})
}

func TestUpdatesiteResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatesite.test", "name", "test-updatesite"),
				),
			},

			{
				Config: testAccUpdatesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatesite.test", "name", "test-updatesite-updated"),
				),
			},
		},
	})
}

func testAccUpdatesiteResourceConfig_basic() string {
	return `

resource "umbrella_updatesite" "test" {
  name = "test-updatesite"
  
  name = "test-resource"
  
}

`
}

func testAccUpdatesiteResourceConfig_update() string {
	return `
resource "umbrella_updatesite" "test" {
  name = "test-updatesite-updated"
  
  name = "test-resource-updated"
  
}
`
}
