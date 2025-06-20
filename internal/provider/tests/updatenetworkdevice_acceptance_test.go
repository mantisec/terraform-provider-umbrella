package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatenetworkdeviceResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatenetworkdevice.test", "name", "test-updatenetworkdevice"),
				),
			},

			{
				Config: testAccUpdatenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatenetworkdevice.test", "name", "test-updatenetworkdevice-updated"),
				),
			},
		},
	})
}

func TestUpdatenetworkdeviceResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatenetworkdevice.test", "name", "test-updatenetworkdevice"),
				),
			},

			{
				Config: testAccUpdatenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatenetworkdevice.test", "name", "test-updatenetworkdevice-updated"),
				),
			},
		},
	})
}

func TestUpdatenetworkdeviceResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatenetworkdevice.test", "name", "test-updatenetworkdevice"),
				),
			},

			{
				Config: testAccUpdatenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatenetworkdevice.test", "name", "test-updatenetworkdevice-updated"),
				),
			},
		},
	})
}

func testAccUpdatenetworkdeviceResourceConfig_basic() string {
	return `

resource "umbrella_updatenetworkdevice" "test" {
  name = "test-updatenetworkdevice"
  
  name = "test-resource"
  
}

`
}

func testAccUpdatenetworkdeviceResourceConfig_update() string {
	return `
resource "umbrella_updatenetworkdevice" "test" {
  name = "test-updatenetworkdevice-updated"
  
  name = "test-resource-updated"
  
}
`
}
