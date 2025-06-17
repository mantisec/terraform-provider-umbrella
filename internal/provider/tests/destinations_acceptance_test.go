package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDestinationsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_destinations.test", "name", "test-destinations"),
				),
			},

			{
				Config: testAccDestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destinations.test", "name", "test-destinations-updated"),
				),
			},
		},
	})
}

func TestDestinationsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_destinations.test", "name", "test-destinations"),
				),
			},

			{
				Config: testAccDestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destinations.test", "name", "test-destinations-updated"),
				),
			},
		},
	})
}

func TestDestinationsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_destinations.test", "name", "test-destinations"),
				),
			},

			{
				Config: testAccDestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destinations.test", "name", "test-destinations-updated"),
				),
			},
		},
	})
}

func testAccDestinationsResourceConfig_basic() string {
	return `

resource "umbrella_destinations" "test" {
  name = "test-destinations"
  
  name = "test-resource"
  
}

`
}

func testAccDestinationsResourceConfig_update() string {
	return `
resource "umbrella_destinations" "test" {
  name = "test-destinations-updated"
  
  name = "test-resource-updated"
  
}
`
}
