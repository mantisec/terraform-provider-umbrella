package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletedestinationlistResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletedestinationlistResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletedestinationlist.test", "name", "test-deletedestinationlist"),
				),
			},

			{
				Config: testAccDeletedestinationlistResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletedestinationlist.test", "name", "test-deletedestinationlist-updated"),
				),
			},
		},
	})
}

func TestDeletedestinationlistResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletedestinationlistResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletedestinationlist.test", "name", "test-deletedestinationlist"),
				),
			},

			{
				Config: testAccDeletedestinationlistResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletedestinationlist.test", "name", "test-deletedestinationlist-updated"),
				),
			},
		},
	})
}

func TestDeletedestinationlistResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletedestinationlistResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletedestinationlist.test", "name", "test-deletedestinationlist"),
				),
			},

			{
				Config: testAccDeletedestinationlistResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletedestinationlist.test", "name", "test-deletedestinationlist-updated"),
				),
			},
		},
	})
}

func testAccDeletedestinationlistResourceConfig_basic() string {
	return `

resource "umbrella_deletedestinationlist" "test" {
  name = "test-deletedestinationlist"
  
  name = "test-resource"
  
}

`
}

func testAccDeletedestinationlistResourceConfig_update() string {
	return `
resource "umbrella_deletedestinationlist" "test" {
  name = "test-deletedestinationlist-updated"
  
  name = "test-resource-updated"
  
}
`
}
