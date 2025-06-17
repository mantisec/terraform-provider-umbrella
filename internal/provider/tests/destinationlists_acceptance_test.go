package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDestinationlistsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationlistsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_destinationlists.test", "name", "test-destinationlists"),
				),
			},

			{
				Config: testAccDestinationlistsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destinationlists.test", "name", "test-destinationlists-updated"),
				),
			},
		},
	})
}

func TestDestinationlistsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationlistsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_destinationlists.test", "name", "test-destinationlists"),
				),
			},

			{
				Config: testAccDestinationlistsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destinationlists.test", "name", "test-destinationlists-updated"),
				),
			},
		},
	})
}

func TestDestinationlistsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationlistsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_destinationlists.test", "name", "test-destinationlists"),
				),
			},

			{
				Config: testAccDestinationlistsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_destinationlists.test", "name", "test-destinationlists-updated"),
				),
			},
		},
	})
}

func testAccDestinationlistsResourceConfig_basic() string {
	return `

resource "umbrella_destinationlists" "test" {
  name = "test-destinationlists"
  
  name = "test-resource"
  
}

`
}

func testAccDestinationlistsResourceConfig_update() string {
	return `
resource "umbrella_destinationlists" "test" {
  name = "test-destinationlists-updated"
  
  name = "test-resource-updated"
  
}
`
}
