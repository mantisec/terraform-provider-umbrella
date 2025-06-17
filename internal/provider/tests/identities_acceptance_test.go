package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestIdentitiesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_identities.test", "name", "test-identities"),
				),
			},

			{
				Config: testAccIdentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_identities.test", "name", "test-identities-updated"),
				),
			},
		},
	})
}

func TestIdentitiesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_identities.test", "name", "test-identities"),
				),
			},

			{
				Config: testAccIdentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_identities.test", "name", "test-identities-updated"),
				),
			},
		},
	})
}

func TestIdentitiesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_identities.test", "name", "test-identities"),
				),
			},

			{
				Config: testAccIdentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_identities.test", "name", "test-identities-updated"),
				),
			},
		},
	})
}

func testAccIdentitiesResourceConfig_basic() string {
	return `

resource "umbrella_identities" "test" {
  name = "test-identities"
  
  name = "test-resource"
  
}

`
}

func testAccIdentitiesResourceConfig_update() string {
	return `
resource "umbrella_identities" "test" {
  name = "test-identities-updated"
  
  name = "test-resource-updated"
  
}
`
}
