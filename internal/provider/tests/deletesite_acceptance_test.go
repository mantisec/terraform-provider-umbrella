package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletesiteResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesite.test", "name", "test-deletesite"),
				),
			},

			{
				Config: testAccDeletesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletesite.test", "name", "test-deletesite-updated"),
				),
			},
		},
	})
}

func TestDeletesiteResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesite.test", "name", "test-deletesite"),
				),
			},

			{
				Config: testAccDeletesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletesite.test", "name", "test-deletesite-updated"),
				),
			},
		},
	})
}

func TestDeletesiteResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesiteResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesite.test", "name", "test-deletesite"),
				),
			},

			{
				Config: testAccDeletesiteResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletesite.test", "name", "test-deletesite-updated"),
				),
			},
		},
	})
}

func testAccDeletesiteResourceConfig_basic() string {
	return `

resource "umbrella_deletesite" "test" {
  name = "test-deletesite"
  
  name = "test-resource"
  
}

`
}

func testAccDeletesiteResourceConfig_update() string {
	return `
resource "umbrella_deletesite" "test" {
  name = "test-deletesite-updated"
  
  name = "test-resource-updated"
  
}
`
}
