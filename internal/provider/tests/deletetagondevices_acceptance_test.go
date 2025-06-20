package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletetagondevicesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletetagondevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletetagondevices.test", "name", "test-deletetagondevices"),
				),
			},

			{
				Config: testAccDeletetagondevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletetagondevices.test", "name", "test-deletetagondevices-updated"),
				),
			},
		},
	})
}

func TestDeletetagondevicesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletetagondevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletetagondevices.test", "name", "test-deletetagondevices"),
				),
			},

			{
				Config: testAccDeletetagondevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletetagondevices.test", "name", "test-deletetagondevices-updated"),
				),
			},
		},
	})
}

func TestDeletetagondevicesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletetagondevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletetagondevices.test", "name", "test-deletetagondevices"),
				),
			},

			{
				Config: testAccDeletetagondevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletetagondevices.test", "name", "test-deletetagondevices-updated"),
				),
			},
		},
	})
}

func testAccDeletetagondevicesResourceConfig_basic() string {
	return `

resource "umbrella_deletetagondevices" "test" {
  name = "test-deletetagondevices"
  
  name = "test-resource"
  
}

`
}

func testAccDeletetagondevicesResourceConfig_update() string {
	return `
resource "umbrella_deletetagondevices" "test" {
  name = "test-deletetagondevices-updated"
  
  name = "test-resource-updated"
  
}
`
}
