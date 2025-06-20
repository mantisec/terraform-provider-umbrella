package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreatecustomeraccessrequestResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomeraccessrequestResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomeraccessrequest.test", "name", "test-createcustomeraccessrequest"),
				),
			},

			{
				Config: testAccCreatecustomeraccessrequestResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcustomeraccessrequest.test", "name", "test-createcustomeraccessrequest-updated"),
				),
			},
		},
	})
}

func TestCreatecustomeraccessrequestResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomeraccessrequestResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomeraccessrequest.test", "name", "test-createcustomeraccessrequest"),
				),
			},

			{
				Config: testAccCreatecustomeraccessrequestResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcustomeraccessrequest.test", "name", "test-createcustomeraccessrequest-updated"),
				),
			},
		},
	})
}

func TestCreatecustomeraccessrequestResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreatecustomeraccessrequestResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_createcustomeraccessrequest.test", "name", "test-createcustomeraccessrequest"),
				),
			},

			{
				Config: testAccCreatecustomeraccessrequestResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_createcustomeraccessrequest.test", "name", "test-createcustomeraccessrequest-updated"),
				),
			},
		},
	})
}

func testAccCreatecustomeraccessrequestResourceConfig_basic() string {
	return `

resource "umbrella_createcustomeraccessrequest" "test" {
  name = "test-createcustomeraccessrequest"
  
  name = "test-resource"
  
}

`
}

func testAccCreatecustomeraccessrequestResourceConfig_update() string {
	return `
resource "umbrella_createcustomeraccessrequest" "test" {
  name = "test-createcustomeraccessrequest-updated"
  
  name = "test-resource-updated"
  
}
`
}
