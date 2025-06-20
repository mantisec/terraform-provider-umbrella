package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPatchapplicationsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPatchapplicationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_patchapplications.test", "name", "test-patchapplications"),
				),
			},

			{
				Config: testAccPatchapplicationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_patchapplications.test", "name", "test-patchapplications-updated"),
				),
			},
		},
	})
}

func TestPatchapplicationsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPatchapplicationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_patchapplications.test", "name", "test-patchapplications"),
				),
			},

			{
				Config: testAccPatchapplicationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_patchapplications.test", "name", "test-patchapplications-updated"),
				),
			},
		},
	})
}

func TestPatchapplicationsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPatchapplicationsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_patchapplications.test", "name", "test-patchapplications"),
				),
			},

			{
				Config: testAccPatchapplicationsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_patchapplications.test", "name", "test-patchapplications-updated"),
				),
			},
		},
	})
}

func testAccPatchapplicationsResourceConfig_basic() string {
	return `

resource "umbrella_patchapplications" "test" {
  name = "test-patchapplications"
  
  name = "test-resource"
  
}

`
}

func testAccPatchapplicationsResourceConfig_update() string {
	return `
resource "umbrella_patchapplications" "test" {
  name = "test-patchapplications-updated"
  
  name = "test-resource-updated"
  
}
`
}
