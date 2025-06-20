package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpdatecustomeraccessrequestResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomeraccessrequestResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomeraccessrequest.test", "name", "test-updatecustomeraccessrequest"),
				),
			},

			{
				Config: testAccUpdatecustomeraccessrequestResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecustomeraccessrequest.test", "name", "test-updatecustomeraccessrequest-updated"),
				),
			},
		},
	})
}

func TestUpdatecustomeraccessrequestResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomeraccessrequestResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomeraccessrequest.test", "name", "test-updatecustomeraccessrequest"),
				),
			},

			{
				Config: testAccUpdatecustomeraccessrequestResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecustomeraccessrequest.test", "name", "test-updatecustomeraccessrequest-updated"),
				),
			},
		},
	})
}

func TestUpdatecustomeraccessrequestResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUpdatecustomeraccessrequestResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_updatecustomeraccessrequest.test", "name", "test-updatecustomeraccessrequest"),
				),
			},

			{
				Config: testAccUpdatecustomeraccessrequestResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_updatecustomeraccessrequest.test", "name", "test-updatecustomeraccessrequest-updated"),
				),
			},
		},
	})
}

func testAccUpdatecustomeraccessrequestResourceConfig_basic() string {
	return `

resource "umbrella_updatecustomeraccessrequest" "test" {
  name = "test-updatecustomeraccessrequest"
  
  name = "test-resource"
  
}

`
}

func testAccUpdatecustomeraccessrequestResourceConfig_update() string {
	return `
resource "umbrella_updatecustomeraccessrequest" "test" {
  name = "test-updatecustomeraccessrequest-updated"
  
  name = "test-resource-updated"
  
}
`
}
