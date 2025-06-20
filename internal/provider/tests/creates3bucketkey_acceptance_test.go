package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCreates3bucketkeyResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreates3bucketkeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_creates3bucketkey.test", "name", "test-creates3bucketkey"),
				),
			},

			{
				Config: testAccCreates3bucketkeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_creates3bucketkey.test", "name", "test-creates3bucketkey-updated"),
				),
			},
		},
	})
}

func TestCreates3bucketkeyResource_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreates3bucketkeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_creates3bucketkey.test", "name", "test-creates3bucketkey"),
				),
			},

			{
				Config: testAccCreates3bucketkeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_creates3bucketkey.test", "name", "test-creates3bucketkey-updated"),
				),
			},
		},
	})
}

func TestCreates3bucketkeyResource_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCreates3bucketkeyResourceConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("umbrella_creates3bucketkey.test", "name", "test-creates3bucketkey"),
				),
			},

			{
				Config: testAccCreates3bucketkeyResourceConfig_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_creates3bucketkey.test", "name", "test-creates3bucketkey-updated"),
				),
			},
		},
	})
}

func testAccCreates3bucketkeyResourceConfig_basic() string {
	return `

resource "umbrella_creates3bucketkey" "test" {
  name = "test-creates3bucketkey"
  
  name = "test-resource"
  
}

`
}

func testAccCreates3bucketkeyResourceConfig_update() string {
	return `
resource "umbrella_creates3bucketkey" "test" {
  name = "test-creates3bucketkey-updated"
  
  name = "test-resource-updated"
  
}
`
}
