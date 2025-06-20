package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatepolicyidentitiesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatepolicyidentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatepolicyidentities.test", "name", "test-updatepolicyidentities"),
				),
			},

			{
				Config: testAccUpdatepolicyidentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatepolicyidentities.test", "name", "test-updatepolicyidentities-updated"),
				),
			},
		},
	})
}

func TestUpdatepolicyidentitiesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatepolicyidentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatepolicyidentities.test", "name", "test-updatepolicyidentities"),
				),
			},

			{
				Config: testAccUpdatepolicyidentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatepolicyidentities.test", "name", "test-updatepolicyidentities-updated"),
				),
			},
		},
	})
}

func TestUpdatepolicyidentitiesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatepolicyidentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatepolicyidentities.test", "name", "test-updatepolicyidentities"),
				),
			},

			{
				Config: testAccUpdatepolicyidentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatepolicyidentities.test", "name", "test-updatepolicyidentities-updated"),
				),
			},
		},
	})
}

func testAccUpdatepolicyidentitiesResourceConfig_basic() string {
	return `

resource "umbrella_updatepolicyidentities" "test" {
  name = "test-updatepolicyidentities"
  
  name = "test-resource"
  
}

`
}

func testAccUpdatepolicyidentitiesResourceConfig_update() string {
	return `
resource "umbrella_updatepolicyidentities" "test" {
  name = "test-updatepolicyidentities-updated"
  
  name = "test-resource-updated"
  
}
`
}
