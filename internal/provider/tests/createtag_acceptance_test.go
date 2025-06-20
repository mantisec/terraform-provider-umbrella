package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatetagResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatetagResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createtag.test", "name", "test-createtag"),
				),
			},

			{
				Config: testAccCreatetagResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createtag.test", "name", "test-createtag-updated"),
				),
			},
		},
	})
}

func TestCreatetagResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatetagResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createtag.test", "name", "test-createtag"),
				),
			},

			{
				Config: testAccCreatetagResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createtag.test", "name", "test-createtag-updated"),
				),
			},
		},
	})
}

func TestCreatetagResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatetagResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createtag.test", "name", "test-createtag"),
				),
			},

			{
				Config: testAccCreatetagResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createtag.test", "name", "test-createtag-updated"),
				),
			},
		},
	})
}

func testAccCreatetagResourceConfig_basic() string {
	return `

resource "umbrella_createtag" "test" {
  name = "test-createtag"
  
  name = "test-resource"
  
}

`
}

func testAccCreatetagResourceConfig_update() string {
	return `
resource "umbrella_createtag" "test" {
  name = "test-createtag-updated"
  
  name = "test-resource-updated"
  
}
`
}
