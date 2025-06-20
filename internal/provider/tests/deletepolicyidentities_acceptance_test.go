package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletepolicyidentitiesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletepolicyidentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletepolicyidentities.test", "name", "test-deletepolicyidentities"),
				),
			},

			{
				Config: testAccDeletepolicyidentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletepolicyidentities.test", "name", "test-deletepolicyidentities-updated"),
				),
			},
		},
	})
}

func TestDeletepolicyidentitiesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletepolicyidentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletepolicyidentities.test", "name", "test-deletepolicyidentities"),
				),
			},

			{
				Config: testAccDeletepolicyidentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletepolicyidentities.test", "name", "test-deletepolicyidentities-updated"),
				),
			},
		},
	})
}

func TestDeletepolicyidentitiesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletepolicyidentitiesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletepolicyidentities.test", "name", "test-deletepolicyidentities"),
				),
			},

			{
				Config: testAccDeletepolicyidentitiesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletepolicyidentities.test", "name", "test-deletepolicyidentities-updated"),
				),
			},
		},
	})
}

func testAccDeletepolicyidentitiesResourceConfig_basic() string {
	return `

resource "umbrella_deletepolicyidentities" "test" {
  name = "test-deletepolicyidentities"
  
  name = "test-resource"
  
}

`
}

func testAccDeletepolicyidentitiesResourceConfig_update() string {
	return `
resource "umbrella_deletepolicyidentities" "test" {
  name = "test-deletepolicyidentities-updated"
  
  name = "test-resource-updated"
  
}
`
}
