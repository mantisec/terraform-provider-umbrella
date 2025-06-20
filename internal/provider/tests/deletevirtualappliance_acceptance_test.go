package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeletevirtualapplianceResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletevirtualapplianceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletevirtualappliance.test", "name", "test-deletevirtualappliance"),
				),
			},

			{
				Config: testAccDeletevirtualapplianceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletevirtualappliance.test", "name", "test-deletevirtualappliance-updated"),
				),
			},
		},
	})
}

func TestDeletevirtualapplianceResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletevirtualapplianceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletevirtualappliance.test", "name", "test-deletevirtualappliance"),
				),
			},

			{
				Config: testAccDeletevirtualapplianceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletevirtualappliance.test", "name", "test-deletevirtualappliance-updated"),
				),
			},
		},
	})
}

func TestDeletevirtualapplianceResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletevirtualapplianceResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_deletevirtualappliance.test", "name", "test-deletevirtualappliance"),
				),
			},

			{
				Config: testAccDeletevirtualapplianceResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_deletevirtualappliance.test", "name", "test-deletevirtualappliance-updated"),
				),
			},
		},
	})
}

func testAccDeletevirtualapplianceResourceConfig_basic() string {
	return `

resource "umbrella_deletevirtualappliance" "test" {
  name = "test-deletevirtualappliance"
  
  name = "test-resource"
  
}

`
}

func testAccDeletevirtualapplianceResourceConfig_update() string {
	return `
resource "umbrella_deletevirtualappliance" "test" {
  name = "test-deletevirtualappliance-updated"
  
  name = "test-resource-updated"
  
}
`
}
