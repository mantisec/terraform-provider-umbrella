package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletesecurewebgatewaydevicesettingsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "name", "test-deletesecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccDeletesecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "name", "test-deletesecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func TestDeletesecurewebgatewaydevicesettingsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "name", "test-deletesecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccDeletesecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "name", "test-deletesecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func TestDeletesecurewebgatewaydevicesettingsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletesecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "name", "test-deletesecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccDeletesecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletesecurewebgatewaydevicesettings.test", "name", "test-deletesecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func testAccDeletesecurewebgatewaydevicesettingsResourceConfig_basic() string {
	return `

resource "umbrella_deletesecurewebgatewaydevicesettings" "test" {
  name = "test-deletesecurewebgatewaydevicesettings"
  
  name = "test-resource"
  
  originIds = ["item1", "item2"]
  
}

`
}

func testAccDeletesecurewebgatewaydevicesettingsResourceConfig_update() string {
	return `
resource "umbrella_deletesecurewebgatewaydevicesettings" "test" {
  name = "test-deletesecurewebgatewaydevicesettings-updated"
  
  name = "test-resource-updated"
  
  originIds = ["updated1", "updated2"]
  
}
`
}
