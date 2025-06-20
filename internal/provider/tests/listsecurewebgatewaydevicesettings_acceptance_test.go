package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListsecurewebgatewaydevicesettingsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListsecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "name", "test-listsecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccListsecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "name", "test-listsecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func TestListsecurewebgatewaydevicesettingsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListsecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "name", "test-listsecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccListsecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "name", "test-listsecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func TestListsecurewebgatewaydevicesettingsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccListsecurewebgatewaydevicesettingsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "name", "test-listsecurewebgatewaydevicesettings"),
				),
			},

			{
				Config: testAccListsecurewebgatewaydevicesettingsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_listsecurewebgatewaydevicesettings.test", "name", "test-listsecurewebgatewaydevicesettings-updated"),
				),
			},
		},
	})
}

func testAccListsecurewebgatewaydevicesettingsResourceConfig_basic() string {
	return `

resource "umbrella_listsecurewebgatewaydevicesettings" "test" {
  name = "test-listsecurewebgatewaydevicesettings"
  
  name = "test-resource"
  
  originIds = ["item1", "item2"]
  
}

`
}

func testAccListsecurewebgatewaydevicesettingsResourceConfig_update() string {
	return `
resource "umbrella_listsecurewebgatewaydevicesettings" "test" {
  name = "test-listsecurewebgatewaydevicesettings-updated"
  
  name = "test-resource-updated"
  
  originIds = ["updated1", "updated2"]
  
}
`
}
