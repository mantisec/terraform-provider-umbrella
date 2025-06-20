package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatecnameResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcname.test", "name", "test-createcname"),
				),
			},

			{
				Config: testAccCreatecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcname.test", "name", "test-createcname-updated"),
				),
			},
		},
	})
}

func TestCreatecnameResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcname.test", "name", "test-createcname"),
				),
			},

			{
				Config: testAccCreatecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcname.test", "name", "test-createcname-updated"),
				),
			},
		},
	})
}

func TestCreatecnameResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecnameResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcname.test", "name", "test-createcname"),
				),
			},

			{
				Config: testAccCreatecnameResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcname.test", "name", "test-createcname-updated"),
				),
			},
		},
	})
}

func testAccCreatecnameResourceConfig_basic() string {
	return `

resource "umbrella_createcname" "test" {
  name = "test-createcname"
  
  name = "test-resource"
  
}

`
}

func testAccCreatecnameResourceConfig_update() string {
	return `
resource "umbrella_createcname" "test" {
  name = "test-createcname-updated"
  
  name = "test-resource-updated"
  
}
`
}
