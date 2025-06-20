package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletenetworkdeviceResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletenetworkdevice.test", "name", "test-deletenetworkdevice"),
				),
			},

			{
				Config: testAccDeletenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletenetworkdevice.test", "name", "test-deletenetworkdevice-updated"),
				),
			},
		},
	})
}

func TestDeletenetworkdeviceResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletenetworkdevice.test", "name", "test-deletenetworkdevice"),
				),
			},

			{
				Config: testAccDeletenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletenetworkdevice.test", "name", "test-deletenetworkdevice-updated"),
				),
			},
		},
	})
}

func TestDeletenetworkdeviceResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletenetworkdeviceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletenetworkdevice.test", "name", "test-deletenetworkdevice"),
				),
			},

			{
				Config: testAccDeletenetworkdeviceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletenetworkdevice.test", "name", "test-deletenetworkdevice-updated"),
				),
			},
		},
	})
}

func testAccDeletenetworkdeviceResourceConfig_basic() string {
	return `

resource "umbrella_deletenetworkdevice" "test" {
  name = "test-deletenetworkdevice"
  
  name = "test-resource"
  
}

`
}

func testAccDeletenetworkdeviceResourceConfig_update() string {
	return `
resource "umbrella_deletenetworkdevice" "test" {
  name = "test-deletenetworkdevice-updated"
  
  name = "test-resource-updated"
  
}
`
}
