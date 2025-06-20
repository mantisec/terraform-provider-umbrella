package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatedestinationsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinations.test", "name", "test-createdestinations"),
				),
			},

			{
				Config: testAccCreatedestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createdestinations.test", "name", "test-createdestinations-updated"),
				),
			},
		},
	})
}

func TestCreatedestinationsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinations.test", "name", "test-createdestinations"),
				),
			},

			{
				Config: testAccCreatedestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createdestinations.test", "name", "test-createdestinations-updated"),
				),
			},
		},
	})
}

func TestCreatedestinationsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatedestinationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createdestinations.test", "name", "test-createdestinations"),
				),
			},

			{
				Config: testAccCreatedestinationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createdestinations.test", "name", "test-createdestinations-updated"),
				),
			},
		},
	})
}

func testAccCreatedestinationsResourceConfig_basic() string {
	return `

resource "umbrella_createdestinations" "test" {
  name = "test-createdestinations"
  
  name = "test-resource"
  
}

`
}

func testAccCreatedestinationsResourceConfig_update() string {
	return `
resource "umbrella_createdestinations" "test" {
  name = "test-createdestinations-updated"
  
  name = "test-resource-updated"
  
}
`
}
