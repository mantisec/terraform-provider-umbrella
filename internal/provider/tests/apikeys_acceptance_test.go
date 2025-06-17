package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestApikeysResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApikeysResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_apikeys.test", "name", "test-apikeys"),
				),
			},

			{
				Config: testAccApikeysResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_apikeys.test", "name", "test-apikeys-updated"),
				),
			},
		},
	})
}

func TestApikeysResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApikeysResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_apikeys.test", "name", "test-apikeys"),
				),
			},

			{
				Config: testAccApikeysResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_apikeys.test", "name", "test-apikeys-updated"),
				),
			},
		},
	})
}

func TestApikeysResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApikeysResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_apikeys.test", "name", "test-apikeys"),
				),
			},

			{
				Config: testAccApikeysResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_apikeys.test", "name", "test-apikeys-updated"),
				),
			},
		},
	})
}

func testAccApikeysResourceConfig_basic() string {
	return `

resource "umbrella_apikeys" "test" {
  name = "test-apikeys"
  
  name = "test-resource"
  
}

`
}

func testAccApikeysResourceConfig_update() string {
	return `
resource "umbrella_apikeys" "test" {
  name = "test-apikeys-updated"
  
  name = "test-resource-updated"
  
}
`
}
