package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/mantisec/terraform-provider-umbrella/internal/provider"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"umbrella": providerserver.NewProtocol6WithError(provider.NewProvider()),
}

func testAccPreCheck(t *testing.T) {
	// You can add environment variable checks here if needed
	// For example:
	// if v := os.Getenv("UMBRELLA_API_KEY"); v == "" {
	//     t.Fatal("UMBRELLA_API_KEY must be set for acceptance tests")
	// }
	// if v := os.Getenv("UMBRELLA_API_SECRET"); v == "" {
	//     t.Fatal("UMBRELLA_API_SECRET must be set for acceptance tests")
	// }

	// For now, just check if we're in test mode
	if os.Getenv("TF_ACC") == "" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}
}
