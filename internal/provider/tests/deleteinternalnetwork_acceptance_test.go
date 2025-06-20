package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeleteinternalnetworkResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "name", "test-deleteinternalnetwork"),
				),
			},

			{
				Config: testAccDeleteinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "name", "test-deleteinternalnetwork-updated"),
				),
			},
		},
	})
}

func TestDeleteinternalnetworkResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "name", "test-deleteinternalnetwork"),
				),
			},

			{
				Config: testAccDeleteinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "name", "test-deleteinternalnetwork-updated"),
				),
			},
		},
	})
}

func TestDeleteinternalnetworkResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternalnetworkResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "name", "test-deleteinternalnetwork"),
				),
			},

			{
				Config: testAccDeleteinternalnetworkResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteinternalnetwork.test", "name", "test-deleteinternalnetwork-updated"),
				),
			},
		},
	})
}

func testAccDeleteinternalnetworkResourceConfig_basic() string {
	return `

resource "umbrella_deleteinternalnetwork" "test" {
  name = "test-deleteinternalnetwork"
  
  name = "test-resource"
  
}

`
}

func testAccDeleteinternalnetworkResourceConfig_update() string {
	return `
resource "umbrella_deleteinternalnetwork" "test" {
  name = "test-deleteinternalnetwork-updated"
  
  name = "test-resource-updated"
  
}
`
}
