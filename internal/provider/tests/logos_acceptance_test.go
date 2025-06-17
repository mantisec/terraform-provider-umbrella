package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestLogosResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogosResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_logos.test", "name", "test-logos"),
				),
			},

			{
				Config: testAccLogosResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_logos.test", "name", "test-logos-updated"),
				),
			},
		},
	})
}

func TestLogosResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogosResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_logos.test", "name", "test-logos"),
				),
			},

			{
				Config: testAccLogosResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_logos.test", "name", "test-logos-updated"),
				),
			},
		},
	})
}

func TestLogosResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLogosResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_logos.test", "name", "test-logos"),
				),
			},

			{
				Config: testAccLogosResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_logos.test", "name", "test-logos-updated"),
				),
			},
		},
	})
}

func testAccLogosResourceConfig_basic() string {
	return `

resource "umbrella_logos" "test" {
  name = "test-logos"
  
  name = "test-resource"
  
}

`
}

func testAccLogosResourceConfig_update() string {
	return `
resource "umbrella_logos" "test" {
  name = "test-logos-updated"
  
  name = "test-resource-updated"
  
}
`
}
