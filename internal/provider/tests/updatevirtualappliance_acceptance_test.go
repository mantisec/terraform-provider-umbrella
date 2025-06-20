package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatevirtualapplianceResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatevirtualapplianceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatevirtualappliance.test", "name", "test-updatevirtualappliance"),
				),
			},

			{
				Config: testAccUpdatevirtualapplianceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatevirtualappliance.test", "name", "test-updatevirtualappliance-updated"),
				),
			},
		},
	})
}

func TestUpdatevirtualapplianceResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatevirtualapplianceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatevirtualappliance.test", "name", "test-updatevirtualappliance"),
				),
			},

			{
				Config: testAccUpdatevirtualapplianceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatevirtualappliance.test", "name", "test-updatevirtualappliance-updated"),
				),
			},
		},
	})
}

func TestUpdatevirtualapplianceResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatevirtualapplianceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatevirtualappliance.test", "name", "test-updatevirtualappliance"),
				),
			},

			{
				Config: testAccUpdatevirtualapplianceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatevirtualappliance.test", "name", "test-updatevirtualappliance-updated"),
				),
			},
		},
	})
}

func testAccUpdatevirtualapplianceResourceConfig_basic() string {
	return `

resource "umbrella_updatevirtualappliance" "test" {
  name = "test-updatevirtualappliance"
  
  name = "test-resource"
  
  siteId = 123
  
}

`
}

func testAccUpdatevirtualapplianceResourceConfig_update() string {
	return `
resource "umbrella_updatevirtualappliance" "test" {
  name = "test-updatevirtualappliance-updated"
  
  name = "test-resource-updated"
  
  siteId = 456
  
}
`
}
