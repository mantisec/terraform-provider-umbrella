package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTrialextensionsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrialextensionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_trialextensions.test", "name", "test-trialextensions"),
				),
			},

			{
				Config: testAccTrialextensionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_trialextensions.test", "name", "test-trialextensions-updated"),
				),
			},
		},
	})
}

func TestTrialextensionsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrialextensionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_trialextensions.test", "name", "test-trialextensions"),
				),
			},

			{
				Config: testAccTrialextensionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_trialextensions.test", "name", "test-trialextensions-updated"),
				),
			},
		},
	})
}

func TestTrialextensionsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrialextensionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_trialextensions.test", "name", "test-trialextensions"),
				),
			},

			{
				Config: testAccTrialextensionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_trialextensions.test", "name", "test-trialextensions-updated"),
				),
			},
		},
	})
}

func testAccTrialextensionsResourceConfig_basic() string {
	return `

resource "umbrella_trialextensions" "test" {
  name = "test-trialextensions"
  
  name = "test-resource"
  
}

`
}

func testAccTrialextensionsResourceConfig_update() string {
	return `
resource "umbrella_trialextensions" "test" {
  name = "test-trialextensions-updated"
  
  name = "test-resource-updated"
  
}
`
}
