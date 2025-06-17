package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRefreshResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRefreshResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_refresh.test", "name", "test-refresh"),
				),
			},

			{
				Config: testAccRefreshResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_refresh.test", "name", "test-refresh-updated"),
				),
			},
		},
	})
}

func TestRefreshResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRefreshResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_refresh.test", "name", "test-refresh"),
				),
			},

			{
				Config: testAccRefreshResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_refresh.test", "name", "test-refresh-updated"),
				),
			},
		},
	})
}

func TestRefreshResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRefreshResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_refresh.test", "name", "test-refresh"),
				),
			},

			{
				Config: testAccRefreshResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_refresh.test", "name", "test-refresh-updated"),
				),
			},
		},
	})
}

func testAccRefreshResourceConfig_basic() string {
	return `

resource "umbrella_refresh" "test" {
  name = "test-refresh"
  
  name = "test-resource"
  
}

`
}

func testAccRefreshResourceConfig_update() string {
	return `
resource "umbrella_refresh" "test" {
  name = "test-refresh-updated"
  
  name = "test-resource-updated"
  
}
`
}
