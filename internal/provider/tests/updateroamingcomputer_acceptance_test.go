package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdateroamingcomputerResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateroamingcomputerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateroamingcomputer.test", "name", "test-updateroamingcomputer"),
				),
			},

			{
				Config: testAccUpdateroamingcomputerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateroamingcomputer.test", "name", "test-updateroamingcomputer-updated"),
				),
			},
		},
	})
}

func TestUpdateroamingcomputerResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateroamingcomputerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateroamingcomputer.test", "name", "test-updateroamingcomputer"),
				),
			},

			{
				Config: testAccUpdateroamingcomputerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateroamingcomputer.test", "name", "test-updateroamingcomputer-updated"),
				),
			},
		},
	})
}

func TestUpdateroamingcomputerResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateroamingcomputerResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateroamingcomputer.test", "name", "test-updateroamingcomputer"),
				),
			},

			{
				Config: testAccUpdateroamingcomputerResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateroamingcomputer.test", "name", "test-updateroamingcomputer-updated"),
				),
			},
		},
	})
}

func testAccUpdateroamingcomputerResourceConfig_basic() string {
	return `

resource "umbrella_updateroamingcomputer" "test" {
  name = "test-updateroamingcomputer"
  
  name = "test-resource"
  
}

`
}

func testAccUpdateroamingcomputerResourceConfig_update() string {
	return `
resource "umbrella_updateroamingcomputer" "test" {
  name = "test-updateroamingcomputer-updated"
  
  name = "test-resource-updated"
  
}
`
}
