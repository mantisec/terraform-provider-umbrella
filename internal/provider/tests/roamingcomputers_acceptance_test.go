package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRoamingcomputersResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoamingcomputersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_roamingcomputers.test", "name", "test-roamingcomputers"),
				),
			},

			{
				Config: testAccRoamingcomputersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_roamingcomputers.test", "name", "test-roamingcomputers-updated"),
				),
			},
		},
	})
}

func TestRoamingcomputersResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoamingcomputersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_roamingcomputers.test", "name", "test-roamingcomputers"),
				),
			},

			{
				Config: testAccRoamingcomputersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_roamingcomputers.test", "name", "test-roamingcomputers-updated"),
				),
			},
		},
	})
}

func TestRoamingcomputersResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoamingcomputersResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_roamingcomputers.test", "name", "test-roamingcomputers"),
				),
			},

			{
				Config: testAccRoamingcomputersResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_roamingcomputers.test", "name", "test-roamingcomputers-updated"),
				),
			},
		},
	})
}

func testAccRoamingcomputersResourceConfig_basic() string {
	return `

resource "umbrella_roamingcomputers" "test" {
  name = "test-roamingcomputers"
  
  name = "test-resource"
  
}

`
}

func testAccRoamingcomputersResourceConfig_update() string {
	return `
resource "umbrella_roamingcomputers" "test" {
  name = "test-roamingcomputers-updated"
  
  name = "test-resource-updated"
  
}
`
}
