package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestVirtualappliancesResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualappliancesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_virtualappliances.test", "name", "test-virtualappliances"),
				),
			},

			{
				Config: testAccVirtualappliancesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_virtualappliances.test", "name", "test-virtualappliances-updated"),
				),
			},
		},
	})
}

func TestVirtualappliancesResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualappliancesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_virtualappliances.test", "name", "test-virtualappliances"),
				),
			},

			{
				Config: testAccVirtualappliancesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_virtualappliances.test", "name", "test-virtualappliances-updated"),
				),
			},
		},
	})
}

func TestVirtualappliancesResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualappliancesResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_virtualappliances.test", "name", "test-virtualappliances"),
				),
			},

			{
				Config: testAccVirtualappliancesResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_virtualappliances.test", "name", "test-virtualappliances-updated"),
				),
			},
		},
	})
}

func testAccVirtualappliancesResourceConfig_basic() string {
	return `

resource "umbrella_virtualappliances" "test" {
  name = "test-virtualappliances"
  
  name = "test-resource"
  
}

`
}

func testAccVirtualappliancesResourceConfig_update() string {
	return `
resource "umbrella_virtualappliances" "test" {
  name = "test-virtualappliances-updated"
  
  name = "test-resource-updated"
  
}
`
}
