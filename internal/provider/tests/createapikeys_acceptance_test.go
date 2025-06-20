package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreateapikeysResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateapikeysResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createapikeys.test", "name", "test-createapikeys"),
				),
			},

			{
				Config: testAccCreateapikeysResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createapikeys.test", "name", "test-createapikeys-updated"),
				),
			},
		},
	})
}

func TestCreateapikeysResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateapikeysResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createapikeys.test", "name", "test-createapikeys"),
				),
			},

			{
				Config: testAccCreateapikeysResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createapikeys.test", "name", "test-createapikeys-updated"),
				),
			},
		},
	})
}

func TestCreateapikeysResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreateapikeysResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createapikeys.test", "name", "test-createapikeys"),
				),
			},

			{
				Config: testAccCreateapikeysResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createapikeys.test", "name", "test-createapikeys-updated"),
				),
			},
		},
	})
}

func testAccCreateapikeysResourceConfig_basic() string {
	return `

resource "umbrella_createapikeys" "test" {
  name = "test-createapikeys"
  
  name = "test-resource"
  
}

`
}

func testAccCreateapikeysResourceConfig_update() string {
	return `
resource "umbrella_createapikeys" "test" {
  name = "test-createapikeys-updated"
  
  name = "test-resource-updated"
  
}
`
}
