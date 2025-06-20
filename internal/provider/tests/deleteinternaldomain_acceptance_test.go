package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeleteinternaldomainResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternaldomain.test", "name", "test-deleteinternaldomain"),
				),
			},

			{
				Config: testAccDeleteinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteinternaldomain.test", "name", "test-deleteinternaldomain-updated"),
				),
			},
		},
	})
}

func TestDeleteinternaldomainResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternaldomain.test", "name", "test-deleteinternaldomain"),
				),
			},

			{
				Config: testAccDeleteinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteinternaldomain.test", "name", "test-deleteinternaldomain-updated"),
				),
			},
		},
	})
}

func TestDeleteinternaldomainResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteinternaldomain.test", "name", "test-deleteinternaldomain"),
				),
			},

			{
				Config: testAccDeleteinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteinternaldomain.test", "name", "test-deleteinternaldomain-updated"),
				),
			},
		},
	})
}

func testAccDeleteinternaldomainResourceConfig_basic() string {
	return `

resource "umbrella_deleteinternaldomain" "test" {
  name = "test-deleteinternaldomain"
  
  name = "test-resource"
  
}

`
}

func testAccDeleteinternaldomainResourceConfig_update() string {
	return `
resource "umbrella_deleteinternaldomain" "test" {
  name = "test-deleteinternaldomain-updated"
  
  name = "test-resource-updated"
  
}
`
}
