package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletelogoResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletelogoResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletelogo.test", "name", "test-deletelogo"),
				),
			},

			{
				Config: testAccDeletelogoResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletelogo.test", "name", "test-deletelogo-updated"),
				),
			},
		},
	})
}

func TestDeletelogoResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletelogoResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletelogo.test", "name", "test-deletelogo"),
				),
			},

			{
				Config: testAccDeletelogoResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletelogo.test", "name", "test-deletelogo-updated"),
				),
			},
		},
	})
}

func TestDeletelogoResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletelogoResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletelogo.test", "name", "test-deletelogo"),
				),
			},

			{
				Config: testAccDeletelogoResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletelogo.test", "name", "test-deletelogo-updated"),
				),
			},
		},
	})
}

func testAccDeletelogoResourceConfig_basic() string {
	return `

resource "umbrella_deletelogo" "test" {
  name = "test-deletelogo"
  
  name = "test-resource"
  
}

`
}

func testAccDeletelogoResourceConfig_update() string {
	return `
resource "umbrella_deletelogo" "test" {
  name = "test-deletelogo-updated"
  
  name = "test-resource-updated"
  
}
`
}
