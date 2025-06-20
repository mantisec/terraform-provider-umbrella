package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatecustomertrialextensionsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "name", "test-createcustomertrialextensions"),
				),
			},

			{
				Config: testAccCreatecustomertrialextensionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "name", "test-createcustomertrialextensions-updated"),
				),
			},
		},
	})
}

func TestCreatecustomertrialextensionsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "name", "test-createcustomertrialextensions"),
				),
			},

			{
				Config: testAccCreatecustomertrialextensionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "name", "test-createcustomertrialextensions-updated"),
				),
			},
		},
	})
}

func TestCreatecustomertrialextensionsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomertrialextensionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "name", "test-createcustomertrialextensions"),
				),
			},

			{
				Config: testAccCreatecustomertrialextensionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcustomertrialextensions.test", "name", "test-createcustomertrialextensions-updated"),
				),
			},
		},
	})
}

func testAccCreatecustomertrialextensionsResourceConfig_basic() string {
	return `

resource "umbrella_createcustomertrialextensions" "test" {
  name = "test-createcustomertrialextensions"
  
  name = "test-resource"
  
  trialExtensionDays = 123
  
}

`
}

func testAccCreatecustomertrialextensionsResourceConfig_update() string {
	return `
resource "umbrella_createcustomertrialextensions" "test" {
  name = "test-createcustomertrialextensions-updated"
  
  name = "test-resource-updated"
  
  trialExtensionDays = 456
  
}
`
}
