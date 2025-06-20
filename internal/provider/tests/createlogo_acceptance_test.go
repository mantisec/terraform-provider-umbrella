package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatelogoResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatelogoResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createlogo.test", "name", "test-createlogo"),
				),
			},

			{
				Config: testAccCreatelogoResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createlogo.test", "name", "test-createlogo-updated"),
				),
			},
		},
	})
}

func TestCreatelogoResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatelogoResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createlogo.test", "name", "test-createlogo"),
				),
			},

			{
				Config: testAccCreatelogoResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createlogo.test", "name", "test-createlogo-updated"),
				),
			},
		},
	})
}

func TestCreatelogoResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatelogoResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createlogo.test", "name", "test-createlogo"),
				),
			},

			{
				Config: testAccCreatelogoResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createlogo.test", "name", "test-createlogo-updated"),
				),
			},
		},
	})
}

func testAccCreatelogoResourceConfig_basic() string {
	return `

resource "umbrella_createlogo" "test" {
  name = "test-createlogo"
  
  name = "test-resource"
  
}

`
}

func testAccCreatelogoResourceConfig_update() string {
	return `
resource "umbrella_createlogo" "test" {
  name = "test-createlogo-updated"
  
  name = "test-resource-updated"
  
}
`
}
