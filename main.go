package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// -----------------------------------------------------------------------------
// Provider entry point
// -----------------------------------------------------------------------------
func main() {
	providerserver.Serve(context.Background(), NewProvider, providerserver.ServeOpts{})
}
