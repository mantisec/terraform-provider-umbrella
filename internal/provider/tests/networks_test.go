package provider_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const providerConfig = `
provider "umbrella" {
  # Use environment variables for testing
  # Set UMBRELLA_API_KEY, UMBRELLA_API_SECRET, UMBRELLA_ORG_ID
}
`

func TestAccNetworksResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccNetworksResourceConfig("test-network", "192.168.1.0", 30, false, "OPEN"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-network"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "ip_address", "192.168.1.0"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "prefix_length", "30"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "is_dynamic", "false"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "status", "OPEN"),
					resource.TestCheckResourceAttrSet("umbrella_networks.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_networks.test", "origin_id"),
					resource.TestCheckResourceAttrSet("umbrella_networks.test", "created_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "umbrella_networks.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccNetworksResourceConfig("test-network-updated", "192.168.1.0", 30, false, "CLOSED"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-network-updated"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "status", "CLOSED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworksResourceDynamic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			// Create and Read testing for dynamic network
			{
				Config: testAccNetworksResourceConfigDynamic("test-dynamic-network", 32, true, "OPEN"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_networks.test", "name", "test-dynamic-network"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "prefix_length", "32"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "is_dynamic", "true"),
					resource.TestCheckResourceAttr("umbrella_networks.test", "status", "OPEN"),
					resource.TestCheckResourceAttrSet("umbrella_networks.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_networks.test", "origin_id"),
				),
			},
		},
	})
}

func testAccNetworksResourceConfig(name, ipAddress string, prefixLength int, isDynamic bool, status string) string {
	return providerConfig + fmt.Sprintf(`
resource "umbrella_networks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %s
  is_dynamic    = %s
  status        = "%s"
}
`, name, ipAddress, strconv.Itoa(prefixLength), strconv.FormatBool(isDynamic), status)
}

func testAccNetworksResourceConfigDynamic(name string, prefixLength int, isDynamic bool, status string) string {
	return providerConfig + fmt.Sprintf(`
resource "umbrella_networks" "test" {
  name          = "%s"
  prefix_length = %s
  is_dynamic    = %s
  status        = "%s"
}
`, name, strconv.Itoa(prefixLength), strconv.FormatBool(isDynamic), status)
}
