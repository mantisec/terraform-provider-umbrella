package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDevicesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_devices.test", "name", "test-devices"),
				),
			},

			{
				Config: testAccDevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_devices.test", "name", "test-devices-updated"),
				),
			},
		},
	})
}

func TestDevicesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_devices.test", "name", "test-devices"),
				),
			},

			{
				Config: testAccDevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_devices.test", "name", "test-devices-updated"),
				),
			},
		},
	})
}

func TestDevicesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDevicesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_devices.test", "name", "test-devices"),
				),
			},

			{
				Config: testAccDevicesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_devices.test", "name", "test-devices-updated"),
				),
			},
		},
	})
}

func testAccDevicesResourceConfig_basic() string {
	return `

resource "umbrella_devices" "test" {
  name = "test-devices"
  
  name = "test-resource"
  
}

`
}

func testAccDevicesResourceConfig_update() string {
	return `
resource "umbrella_devices" "test" {
  name = "test-devices-updated"
  
  name = "test-resource-updated"
  
}
`
}
