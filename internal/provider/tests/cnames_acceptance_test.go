package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCnamesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCnamesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_cnames.test", "name", "test-cnames"),
				),
			},

			{
				Config: testAccCnamesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_cnames.test", "name", "test-cnames-updated"),
				),
			},
		},
	})
}

func TestCnamesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCnamesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_cnames.test", "name", "test-cnames"),
				),
			},

			{
				Config: testAccCnamesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_cnames.test", "name", "test-cnames-updated"),
				),
			},
		},
	})
}

func TestCnamesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCnamesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_cnames.test", "name", "test-cnames"),
				),
			},

			{
				Config: testAccCnamesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_cnames.test", "name", "test-cnames-updated"),
				),
			},
		},
	})
}

func testAccCnamesResourceConfig_basic() string {
	return `

resource "umbrella_cnames" "test" {
  name = "test-cnames"
  
  name = "test-resource"
  
}

`
}

func testAccCnamesResourceConfig_update() string {
	return `
resource "umbrella_cnames" "test" {
  name = "test-cnames-updated"
  
  name = "test-resource-updated"
  
}
`
}
