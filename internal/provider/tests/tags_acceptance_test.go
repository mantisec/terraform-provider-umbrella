package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTagsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTagsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_tags.test", "name", "test-tags"),
				),
			},

			{
				Config: testAccTagsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_tags.test", "name", "test-tags-updated"),
				),
			},
		},
	})
}

func TestTagsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTagsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_tags.test", "name", "test-tags"),
				),
			},

			{
				Config: testAccTagsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_tags.test", "name", "test-tags-updated"),
				),
			},
		},
	})
}

func TestTagsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTagsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_tags.test", "name", "test-tags"),
				),
			},

			{
				Config: testAccTagsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_tags.test", "name", "test-tags-updated"),
				),
			},
		},
	})
}

func testAccTagsResourceConfig_basic() string {
	return `

resource "umbrella_tags" "test" {
  name = "test-tags"
  
  name = "test-resource"
  
}

`
}

func testAccTagsResourceConfig_update() string {
	return `
resource "umbrella_tags" "test" {
  name = "test-tags-updated"
  
  name = "test-resource-updated"
  
}
`
}
