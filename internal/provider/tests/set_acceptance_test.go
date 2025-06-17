package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSetResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSetResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_set.test", "name", "test-set"),
				),
			},

			{
				Config: testAccSetResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_set.test", "name", "test-set-updated"),
				),
			},
		},
	})
}

func TestSetResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSetResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_set.test", "name", "test-set"),
				),
			},

			{
				Config: testAccSetResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_set.test", "name", "test-set-updated"),
				),
			},
		},
	})
}

func TestSetResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSetResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_set.test", "name", "test-set"),
				),
			},

			{
				Config: testAccSetResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_set.test", "name", "test-set-updated"),
				),
			},
		},
	})
}

func testAccSetResourceConfig_basic() string {
	return `

resource "umbrella_set" "test" {
  name = "test-set"
  
  name = "test-resource"
  
}

`
}

func testAccSetResourceConfig_update() string {
	return `
resource "umbrella_set" "test" {
  name = "test-set-updated"
  
  name = "test-resource-updated"
  
}
`
}
