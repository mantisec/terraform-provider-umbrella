package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreateinternaldomainResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createinternaldomain.test", "name", "test-createinternaldomain"),
				),
			},

			{
				Config: testAccCreateinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createinternaldomain.test", "name", "test-createinternaldomain-updated"),
				),
			},
		},
	})
}

func TestCreateinternaldomainResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createinternaldomain.test", "name", "test-createinternaldomain"),
				),
			},

			{
				Config: testAccCreateinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createinternaldomain.test", "name", "test-createinternaldomain-updated"),
				),
			},
		},
	})
}

func TestCreateinternaldomainResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateinternaldomainResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createinternaldomain.test", "name", "test-createinternaldomain"),
				),
			},

			{
				Config: testAccCreateinternaldomainResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createinternaldomain.test", "name", "test-createinternaldomain-updated"),
				),
			},
		},
	})
}

func testAccCreateinternaldomainResourceConfig_basic() string {
	return `

resource "umbrella_createinternaldomain" "test" {
  name = "test-createinternaldomain"
  
  name = "test-resource"
  
}

`
}

func testAccCreateinternaldomainResourceConfig_update() string {
	return `
resource "umbrella_createinternaldomain" "test" {
  name = "test-createinternaldomain-updated"
  
  name = "test-resource-updated"
  
}
`
}
