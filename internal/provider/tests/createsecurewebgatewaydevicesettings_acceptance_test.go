package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatesecurewebgatewaydevicesettingsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "name", "test-createsecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccCreatesecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "name", "test-createsecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func TestCreatesecurewebgatewaydevicesettingsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "name", "test-createsecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccCreatesecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "name", "test-createsecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func TestCreatesecurewebgatewaydevicesettingsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatesecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "name", "test-createsecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccCreatesecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createsecurewebgatewaydevicesettings.test", "name", "test-createsecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func testAccCreatesecurewebgatewaydevicesettingsResourceConfig_basic() string {
	return `

resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  name = "test-createsecurewebgatewaydevicesettings"
  
  name = "test-resource"
  
  originIds = ["item1", "item2"]
  
  value = "0"
  
}

`
}

func testAccCreatesecurewebgatewaydevicesettingsResourceConfig_update() string {
	return `
resource "umbrella_createsecurewebgatewaydevicesettings" "test" {
  name = "test-createsecurewebgatewaydevicesettings-updated"
  
  name = "test-resource-updated"
  
  originIds = ["updated1", "updated2"]
  
  value = "1"
  
}
`
}
