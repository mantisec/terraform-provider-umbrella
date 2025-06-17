package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestInternaldomainsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInternaldomainsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_internaldomains.test", "name", "test-internaldomains"),
				),
			},

			{
				Config: testAccInternaldomainsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internaldomains.test", "name", "test-internaldomains-updated"),
				),
			},
		},
	})
}

func TestInternaldomainsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInternaldomainsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_internaldomains.test", "name", "test-internaldomains"),
				),
			},

			{
				Config: testAccInternaldomainsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internaldomains.test", "name", "test-internaldomains-updated"),
				),
			},
		},
	})
}

func TestInternaldomainsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInternaldomainsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_internaldomains.test", "name", "test-internaldomains"),
				),
			},

			{
				Config: testAccInternaldomainsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internaldomains.test", "name", "test-internaldomains-updated"),
				),
			},
		},
	})
}

func testAccInternaldomainsResourceConfig_basic() string {
	return `

resource "umbrella_internaldomains" "test" {
  name = "test-internaldomains"
  
  name = "test-resource"
  
}

`
}

func testAccInternaldomainsResourceConfig_update() string {
	return `
resource "umbrella_internaldomains" "test" {
  name = "test-internaldomains-updated"
  
  name = "test-resource-updated"
  
}
`
}
