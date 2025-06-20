package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeleteapikeyResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteapikey.test", "name", "test-deleteapikey"),
				),
			},

			{
				Config: testAccDeleteapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteapikey.test", "name", "test-deleteapikey-updated"),
				),
			},
		},
	})
}

func TestDeleteapikeyResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteapikey.test", "name", "test-deleteapikey"),
				),
			},

			{
				Config: testAccDeleteapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteapikey.test", "name", "test-deleteapikey-updated"),
				),
			},
		},
	})
}

func TestDeleteapikeyResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeleteapikeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deleteapikey.test", "name", "test-deleteapikey"),
				),
			},

			{
				Config: testAccDeleteapikeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deleteapikey.test", "name", "test-deleteapikey-updated"),
				),
			},
		},
	})
}

func testAccDeleteapikeyResourceConfig_basic() string {
	return `

resource "umbrella_deleteapikey" "test" {
  name = "test-deleteapikey"
  
  name = "test-resource"
  
}

`
}

func testAccDeleteapikeyResourceConfig_update() string {
	return `
resource "umbrella_deleteapikey" "test" {
  name = "test-deleteapikey-updated"
  
  name = "test-resource-updated"
  
}
`
}
