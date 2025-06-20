package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatesiteResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsite.test", "name", "test-createsite"),
				),
			},

			{
				Config: testAccCreatesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createsite.test", "name", "test-createsite-updated"),
				),
			},
		},
	})
}

func TestCreatesiteResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsite.test", "name", "test-createsite"),
				),
			},

			{
				Config: testAccCreatesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createsite.test", "name", "test-createsite-updated"),
				),
			},
		},
	})
}

func TestCreatesiteResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsite.test", "name", "test-createsite"),
				),
			},

			{
				Config: testAccCreatesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createsite.test", "name", "test-createsite-updated"),
				),
			},
		},
	})
}

func testAccCreatesiteResourceConfig_basic() string {
	return `

resource "umbrella_createsite" "test" {
  name = "test-createsite"
  
  name = "test-resource"
  
}

`
}

func testAccCreatesiteResourceConfig_update() string {
	return `
resource "umbrella_createsite" "test" {
  name = "test-createsite-updated"
  
  name = "test-resource-updated"
  
}
`
}
