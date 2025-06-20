package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeleteroamingcomputerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteroamingcomputerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteroamingcomputer.test", "name", "test-deleteroamingcomputer"),
				),
			},

			{
				Config: testAccDeleteroamingcomputerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteroamingcomputer.test", "name", "test-deleteroamingcomputer-updated"),
				),
			},
		},
	})
}

func TestDeleteroamingcomputerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteroamingcomputerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteroamingcomputer.test", "name", "test-deleteroamingcomputer"),
				),
			},

			{
				Config: testAccDeleteroamingcomputerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteroamingcomputer.test", "name", "test-deleteroamingcomputer-updated"),
				),
			},
		},
	})
}

func TestDeleteroamingcomputerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteroamingcomputerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteroamingcomputer.test", "name", "test-deleteroamingcomputer"),
				),
			},

			{
				Config: testAccDeleteroamingcomputerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteroamingcomputer.test", "name", "test-deleteroamingcomputer-updated"),
				),
			},
		},
	})
}

func testAccDeleteroamingcomputerResourceConfig_basic() string {
	return `

resource "umbrella_deleteroamingcomputer" "test" {
  name = "test-deleteroamingcomputer"
  
  name = "test-resource"
  
}

`
}

func testAccDeleteroamingcomputerResourceConfig_update() string {
	return `
resource "umbrella_deleteroamingcomputer" "test" {
  name = "test-deleteroamingcomputer-updated"
  
  name = "test-resource-updated"
  
}
`
}
