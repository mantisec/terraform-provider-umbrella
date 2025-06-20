package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletedestinationsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletedestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletedestinations.test", "name", "test-deletedestinations"),
				),
			},

			{
				Config: testAccDeletedestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletedestinations.test", "name", "test-deletedestinations-updated"),
				),
			},
		},
	})
}

func TestDeletedestinationsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletedestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletedestinations.test", "name", "test-deletedestinations"),
				),
			},

			{
				Config: testAccDeletedestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletedestinations.test", "name", "test-deletedestinations-updated"),
				),
			},
		},
	})
}

func TestDeletedestinationsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletedestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletedestinations.test", "name", "test-deletedestinations"),
				),
			},

			{
				Config: testAccDeletedestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletedestinations.test", "name", "test-deletedestinations-updated"),
				),
			},
		},
	})
}

func testAccDeletedestinationsResourceConfig_basic() string {
	return `

resource "umbrella_deletedestinations" "test" {
  name = "test-deletedestinations"
  
  name = "test-resource"
  
}

`
}

func testAccDeletedestinationsResourceConfig_update() string {
	return `
resource "umbrella_deletedestinations" "test" {
  name = "test-deletedestinations-updated"
  
  name = "test-resource-updated"
  
}
`
}
