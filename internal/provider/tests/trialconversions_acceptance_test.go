package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestTrialconversionsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrialconversionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_trialconversions.test", "name", "test-trialconversions"),
				),
			},

			{
				Config: testAccTrialconversionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_trialconversions.test", "name", "test-trialconversions-updated"),
				),
			},
		},
	})
}

func TestTrialconversionsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrialconversionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_trialconversions.test", "name", "test-trialconversions"),
				),
			},

			{
				Config: testAccTrialconversionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_trialconversions.test", "name", "test-trialconversions-updated"),
				),
			},
		},
	})
}

func TestTrialconversionsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTrialconversionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_trialconversions.test", "name", "test-trialconversions"),
				),
			},

			{
				Config: testAccTrialconversionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_trialconversions.test", "name", "test-trialconversions-updated"),
				),
			},
		},
	})
}

func testAccTrialconversionsResourceConfig_basic() string {
	return `

resource "umbrella_trialconversions" "test" {
  name = "test-trialconversions"
  
  name = "test-resource"
  
}

`
}

func testAccTrialconversionsResourceConfig_update() string {
	return `
resource "umbrella_trialconversions" "test" {
  name = "test-trialconversions-updated"
  
  name = "test-resource-updated"
  
}
`
}
