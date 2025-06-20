package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletenetworkResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletenetwork.test", "name", "test-deletenetwork"),
				),
			},

			{
				Config: testAccDeletenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletenetwork.test", "name", "test-deletenetwork-updated"),
				),
			},
		},
	})
}

func TestDeletenetworkResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletenetwork.test", "name", "test-deletenetwork"),
				),
			},

			{
				Config: testAccDeletenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletenetwork.test", "name", "test-deletenetwork-updated"),
				),
			},
		},
	})
}

func TestDeletenetworkResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletenetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletenetwork.test", "name", "test-deletenetwork"),
				),
			},

			{
				Config: testAccDeletenetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletenetwork.test", "name", "test-deletenetwork-updated"),
				),
			},
		},
	})
}

func testAccDeletenetworkResourceConfig_basic() string {
	return `

resource "umbrella_deletenetwork" "test" {
  name = "test-deletenetwork"
  
  name = "test-resource"
  
}

`
}

func testAccDeletenetworkResourceConfig_update() string {
	return `
resource "umbrella_deletenetwork" "test" {
  name = "test-deletenetwork-updated"
  
  name = "test-resource-updated"
  
}
`
}
