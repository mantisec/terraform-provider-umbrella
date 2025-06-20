package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPuttrialconversionsResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "name", "test-puttrialconversions"),
				),
			},

			{
				Config: testAccPuttrialconversionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "name", "test-puttrialconversions-updated"),
				),
			},
		},
	})
}

func TestPuttrialconversionsResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "name", "test-puttrialconversions"),
				),
			},

			{
				Config: testAccPuttrialconversionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "name", "test-puttrialconversions-updated"),
				),
			},
		},
	})
}

func TestPuttrialconversionsResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPuttrialconversionsResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "name", "test-puttrialconversions"),
				),
			},

			{
				Config: testAccPuttrialconversionsResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_puttrialconversions.test", "name", "test-puttrialconversions-updated"),
				),
			},
		},
	})
}

func testAccPuttrialconversionsResourceConfig_basic() string {
	return `

resource "umbrella_puttrialconversions" "test" {
  name = "test-puttrialconversions"
  
  name = "test-resource"
  
  packageId = 123
  
}

`
}

func testAccPuttrialconversionsResourceConfig_update() string {
	return `
resource "umbrella_puttrialconversions" "test" {
  name = "test-puttrialconversions-updated"
  
  name = "test-resource-updated"
  
  packageId = 456
  
}
`
}
