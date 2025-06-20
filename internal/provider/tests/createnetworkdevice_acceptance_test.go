package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatenetworkdeviceResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "name", "test-createnetworkdevice"),
				),
			},

			{
				Config: testAccCreatenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "name", "test-createnetworkdevice-updated"),
				),
			},
		},
	})
}

func TestCreatenetworkdeviceResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "name", "test-createnetworkdevice"),
				),
			},

			{
				Config: testAccCreatenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "name", "test-createnetworkdevice-updated"),
				),
			},
		},
	})
}

func TestCreatenetworkdeviceResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "name", "test-createnetworkdevice"),
				),
			},

			{
				Config: testAccCreatenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createnetworkdevice.test", "name", "test-createnetworkdevice-updated"),
				),
			},
		},
	})
}

func testAccCreatenetworkdeviceResourceConfig_basic() string {
	return `

resource "umbrella_createnetworkdevice" "test" {
  name = "test-createnetworkdevice"
  
  name = "test-resource"
  
  serialNumber = "test-value"
  
  tag = "test-value"
  
  model = "test-value"
  
  macAddress = "test-value"
  
}

`
}

func testAccCreatenetworkdeviceResourceConfig_update() string {
	return `
resource "umbrella_createnetworkdevice" "test" {
  name = "test-createnetworkdevice-updated"
  
  name = "test-resource-updated"
  
  serialNumber = "updated-value"
  
  tag = "updated-value"
  
  model = "updated-value"
  
  macAddress = "updated-value"
  
}
`
}
