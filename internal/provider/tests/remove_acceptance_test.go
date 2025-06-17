package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRemoveResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRemoveResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_remove.test", "name", "test-remove"),
				),
			},

			{
				Config: testAccRemoveResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_remove.test", "name", "test-remove-updated"),
				),
			},
		},
	})
}

func TestRemoveResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRemoveResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_remove.test", "name", "test-remove"),
				),
			},

			{
				Config: testAccRemoveResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_remove.test", "name", "test-remove-updated"),
				),
			},
		},
	})
}

func TestRemoveResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRemoveResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_remove.test", "name", "test-remove"),
				),
			},

			{
				Config: testAccRemoveResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_remove.test", "name", "test-remove-updated"),
				),
			},
		},
	})
}

func testAccRemoveResourceConfig_basic() string {
	return `

resource "umbrella_remove" "test" {
  name = "test-remove"
  
  name = "test-resource"
  
}

`
}

func testAccRemoveResourceConfig_update() string {
	return `
resource "umbrella_remove" "test" {
  name = "test-remove-updated"
  
  name = "test-resource-updated"
  
}
`
}
