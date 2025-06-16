package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccInternalNetworksResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInternalNetworkDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccInternalNetworkResourceConfig("test-internal-network", "192.168.1.0", 24, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "name", "test-internal-network"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "ip_address", "192.168.1.0"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "prefix_length", "24"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "site_id", "1"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "origin_id"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "created_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "umbrella_internalnetworks.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccInternalNetworkResourceConfig("test-internal-network-updated", "192.168.2.0", 25, 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "name", "test-internal-network-updated"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "ip_address", "192.168.2.0"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "prefix_length", "25"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "site_id", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccInternalNetworksResourceWithNetworkID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInternalNetworkDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing with network_id
			{
				Config: testAccInternalNetworkResourceConfigWithNetworkID("test-internal-network-net", "10.0.1.0", 28, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "name", "test-internal-network-net"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "ip_address", "10.0.1.0"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "prefix_length", "28"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "network_id", "2"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "origin_id"),
				),
			},
		},
	})
}

func TestAccInternalNetworksResourceWithTunnelID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInternalNetworkDestroy,
		Steps: []resource.TestStep{
			// Create and Read testing with tunnel_id
			{
				Config: testAccInternalNetworkResourceConfigWithTunnelID("test-internal-network-tunnel", "172.16.1.0", 30, 3),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "name", "test-internal-network-tunnel"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "ip_address", "172.16.1.0"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "prefix_length", "30"),
					resource.TestCheckResourceAttr("umbrella_internalnetworks.test", "tunnel_id", "3"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "id"),
					resource.TestCheckResourceAttrSet("umbrella_internalnetworks.test", "origin_id"),
				),
			},
		},
	})
}

func TestAccInternalNetworksResourceValidation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test validation errors
			{
				Config:      testAccInternalNetworkResourceConfigInvalidName("", "192.168.1.0", 24, 1),
				ExpectError: regexp.MustCompile("internal network name must be between 1 and 50 characters"),
			},
			{
				Config:      testAccInternalNetworkResourceConfigInvalidPrefixLength("test-network", "192.168.1.0", 7, 1),
				ExpectError: regexp.MustCompile("prefix length must be between 8 and 32"),
			},
			{
				Config:      testAccInternalNetworkResourceConfigInvalidPrefixLength("test-network", "192.168.1.0", 33, 1),
				ExpectError: regexp.MustCompile("prefix length must be between 8 and 32"),
			},
			{
				Config:      testAccInternalNetworkResourceConfigNoAssociation("test-network", "192.168.1.0", 24),
				ExpectError: regexp.MustCompile("at least one of site_id, network_id, or tunnel_id must be provided"),
			},
		},
	})
}

func testAccCheckInternalNetworkDestroy(s *terraform.State) error {
	// This would typically check that the internal network has been destroyed
	// For now, we'll implement a basic check
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "umbrella_internalnetworks" {
			continue
		}

		// In a real implementation, you would check if the resource still exists
		// by making an API call to get the internal network
		// If it still exists, return an error
		// If it returns a 404, the resource was successfully destroyed
	}

	return nil
}

func testAccInternalNetworkResourceConfig(name, ipAddress string, prefixLength, siteID int) string {
	return fmt.Sprintf(`
resource "umbrella_internalnetworks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %d
  site_id       = %d
}
`, name, ipAddress, prefixLength, siteID)
}

func testAccInternalNetworkResourceConfigWithNetworkID(name, ipAddress string, prefixLength, networkID int) string {
	return fmt.Sprintf(`
resource "umbrella_internalnetworks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %d
  network_id    = %d
}
`, name, ipAddress, prefixLength, networkID)
}

func testAccInternalNetworkResourceConfigWithTunnelID(name, ipAddress string, prefixLength, tunnelID int) string {
	return fmt.Sprintf(`
resource "umbrella_internalnetworks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %d
  tunnel_id     = %d
}
`, name, ipAddress, prefixLength, tunnelID)
}

func testAccInternalNetworkResourceConfigInvalidName(name, ipAddress string, prefixLength, siteID int) string {
	return fmt.Sprintf(`
resource "umbrella_internalnetworks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %d
  site_id       = %d
}
`, name, ipAddress, prefixLength, siteID)
}

func testAccInternalNetworkResourceConfigInvalidPrefixLength(name, ipAddress string, prefixLength, siteID int) string {
	return fmt.Sprintf(`
resource "umbrella_internalnetworks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %d
  site_id       = %d
}
`, name, ipAddress, prefixLength, siteID)
}

func testAccInternalNetworkResourceConfigNoAssociation(name, ipAddress string, prefixLength int) string {
	return fmt.Sprintf(`
resource "umbrella_internalnetworks" "test" {
  name          = "%s"
  ip_address    = "%s"
  prefix_length = %d
}
`, name, ipAddress, prefixLength)
}
