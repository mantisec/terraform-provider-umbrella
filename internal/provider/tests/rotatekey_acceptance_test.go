package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRotatekeyResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRotatekeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_rotatekey.test", "name", "test-rotatekey"),
				),
			},

			{
				Config: testAccRotatekeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_rotatekey.test", "name", "test-rotatekey-updated"),
				),
			},
		},
	})
}

func TestRotatekeyResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRotatekeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_rotatekey.test", "name", "test-rotatekey"),
				),
			},

			{
				Config: testAccRotatekeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_rotatekey.test", "name", "test-rotatekey-updated"),
				),
			},
		},
	})
}

func TestRotatekeyResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRotatekeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_rotatekey.test", "name", "test-rotatekey"),
				),
			},

			{
				Config: testAccRotatekeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_rotatekey.test", "name", "test-rotatekey-updated"),
				),
			},
		},
	})
}

func testAccRotatekeyResourceConfig_basic() string {
	return `

resource "umbrella_rotatekey" "test" {
  name = "test-rotatekey"
  
  name = "test-resource"
  
}

`
}

func testAccRotatekeyResourceConfig_update() string {
	return `
resource "umbrella_rotatekey" "test" {
  name = "test-rotatekey-updated"
  
  name = "test-resource-updated"
  
}
`
}
