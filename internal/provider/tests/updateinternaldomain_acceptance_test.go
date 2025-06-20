package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdateinternaldomainResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateinternaldomain.test", "name", "test-updateinternaldomain"),
				),
			},

			{
				Config: testAccUpdateinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateinternaldomain.test", "name", "test-updateinternaldomain-updated"),
				),
			},
		},
	})
}

func TestUpdateinternaldomainResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateinternaldomain.test", "name", "test-updateinternaldomain"),
				),
			},

			{
				Config: testAccUpdateinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateinternaldomain.test", "name", "test-updateinternaldomain-updated"),
				),
			},
		},
	})
}

func TestUpdateinternaldomainResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdateinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updateinternaldomain.test", "name", "test-updateinternaldomain"),
				),
			},

			{
				Config: testAccUpdateinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updateinternaldomain.test", "name", "test-updateinternaldomain-updated"),
				),
			},
		},
	})
}

func testAccUpdateinternaldomainResourceConfig_basic() string {
	return `

resource "umbrella_updateinternaldomain" "test" {
  name = "test-updateinternaldomain"
  
  name = "test-resource"
  
}

`
}

func testAccUpdateinternaldomainResourceConfig_update() string {
	return `
resource "umbrella_updateinternaldomain" "test" {
  name = "test-updateinternaldomain-updated"
  
  name = "test-resource-updated"
  
}
`
}
