package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPasswordresetsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPasswordresetsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_passwordresets.test", "name", "test-passwordresets"),
				),
			},

			{
				Config: testAccPasswordresetsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_passwordresets.test", "name", "test-passwordresets-updated"),
				),
			},
		},
	})
}

func TestPasswordresetsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPasswordresetsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_passwordresets.test", "name", "test-passwordresets"),
				),
			},

			{
				Config: testAccPasswordresetsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_passwordresets.test", "name", "test-passwordresets-updated"),
				),
			},
		},
	})
}

func TestPasswordresetsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPasswordresetsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_passwordresets.test", "name", "test-passwordresets"),
				),
			},

			{
				Config: testAccPasswordresetsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_passwordresets.test", "name", "test-passwordresets-updated"),
				),
			},
		},
	})
}

func testAccPasswordresetsResourceConfig_basic() string {
	return `

resource "umbrella_passwordresets" "test" {
  name = "test-passwordresets"
  
  name = "test-resource"
  
}

`
}

func testAccPasswordresetsResourceConfig_update() string {
	return `
resource "umbrella_passwordresets" "test" {
  name = "test-passwordresets-updated"
  
  name = "test-resource-updated"
  
}
`
}
