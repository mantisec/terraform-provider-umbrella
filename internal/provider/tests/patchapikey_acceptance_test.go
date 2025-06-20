package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPatchapikeyResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPatchapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_patchapikey.test", "name", "test-patchapikey"),
				),
			},

			{
				Config: testAccPatchapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_patchapikey.test", "name", "test-patchapikey-updated"),
				),
			},
		},
	})
}

func TestPatchapikeyResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPatchapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_patchapikey.test", "name", "test-patchapikey"),
				),
			},

			{
				Config: testAccPatchapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_patchapikey.test", "name", "test-patchapikey-updated"),
				),
			},
		},
	})
}

func TestPatchapikeyResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPatchapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_patchapikey.test", "name", "test-patchapikey"),
				),
			},

			{
				Config: testAccPatchapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_patchapikey.test", "name", "test-patchapikey-updated"),
				),
			},
		},
	})
}

func testAccPatchapikeyResourceConfig_basic() string {
	return `

resource "umbrella_patchapikey" "test" {
  name = "test-patchapikey"
  
  name = "test-resource"
  
}

`
}

func testAccPatchapikeyResourceConfig_update() string {
	return `
resource "umbrella_patchapikey" "test" {
  name = "test-patchapikey-updated"
  
  name = "test-resource-updated"
  
}
`
}
