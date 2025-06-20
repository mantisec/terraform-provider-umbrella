package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAddtagtodevicesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtagtodevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtagtodevices.test", "name", "test-addtagtodevices"),
				),
			},

			{
				Config: testAccAddtagtodevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_addtagtodevices.test", "name", "test-addtagtodevices-updated"),
				),
			},
		},
	})
}

func TestAddtagtodevicesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtagtodevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtagtodevices.test", "name", "test-addtagtodevices"),
				),
			},

			{
				Config: testAccAddtagtodevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_addtagtodevices.test", "name", "test-addtagtodevices-updated"),
				),
			},
		},
	})
}

func TestAddtagtodevicesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAddtagtodevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_addtagtodevices.test", "name", "test-addtagtodevices"),
				),
			},

			{
				Config: testAccAddtagtodevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_addtagtodevices.test", "name", "test-addtagtodevices-updated"),
				),
			},
		},
	})
}

func testAccAddtagtodevicesResourceConfig_basic() string {
	return `

resource "umbrella_addtagtodevices" "test" {
  name = "test-addtagtodevices"
  
  name = "test-resource"
  
}

`
}

func testAccAddtagtodevicesResourceConfig_update() string {
	return `
resource "umbrella_addtagtodevices" "test" {
  name = "test-addtagtodevices-updated"
  
  name = "test-resource-updated"
  
}
`
}
