package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatedestinationlistResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "name", "test-createdestinationlist"),
				),
			},

			{
				Config: testAccCreatedestinationlistResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "name", "test-createdestinationlist-updated"),
				),
			},
		},
	})
}

func TestCreatedestinationlistResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "name", "test-createdestinationlist"),
				),
			},

			{
				Config: testAccCreatedestinationlistResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "name", "test-createdestinationlist-updated"),
				),
			},
		},
	})
}

func TestCreatedestinationlistResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationlistResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "name", "test-createdestinationlist"),
				),
			},

			{
				Config: testAccCreatedestinationlistResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createdestinationlist.test", "name", "test-createdestinationlist-updated"),
				),
			},
		},
	})
}

func testAccCreatedestinationlistResourceConfig_basic() string {
	return `

resource "umbrella_createdestinationlist" "test" {
  name = "test-createdestinationlist"
  
  name = "test-resource"
  
  bundleTypeId = 123
  
  destinations = ["item1", "item2"]
  
  access = "allow"
  
  isGlobal = true
  
}

`
}

func testAccCreatedestinationlistResourceConfig_update() string {
	return `
resource "umbrella_createdestinationlist" "test" {
  name = "test-createdestinationlist-updated"
  
  name = "test-resource-updated"
  
  bundleTypeId = 456
  
  destinations = ["updated1", "updated2"]
  
  access = "block"
  
  isGlobal = false
  
}
`
}
