package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatepasswordresetsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatepasswordresetsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "name", "test-createpasswordresets"),
				),
			},

			{
				Config: testAccCreatepasswordresetsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "name", "test-createpasswordresets-updated"),
				),
			},
		},
	})
}

func TestCreatepasswordresetsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatepasswordresetsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "name", "test-createpasswordresets"),
				),
			},

			{
				Config: testAccCreatepasswordresetsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "name", "test-createpasswordresets-updated"),
				),
			},
		},
	})
}

func TestCreatepasswordresetsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatepasswordresetsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "name", "test-createpasswordresets"),
				),
			},

			{
				Config: testAccCreatepasswordresetsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createpasswordresets.test", "name", "test-createpasswordresets-updated"),
				),
			},
		},
	})
}

func testAccCreatepasswordresetsResourceConfig_basic() string {
	return `

resource "umbrella_createpasswordresets" "test" {
  name = "test-createpasswordresets"
  
  name = "test-resource"
  
  adminEmails = ["item1", "item2"]
  
}

`
}

func testAccCreatepasswordresetsResourceConfig_update() string {
	return `
resource "umbrella_createpasswordresets" "test" {
  name = "test-createpasswordresets-updated"
  
  name = "test-resource-updated"
  
  adminEmails = ["updated1", "updated2"]
  
}
`
}
