package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/mantisec/terraform-provider-umbrella/internal/provider"
)

// -----------------------------------------------------------------------------
// Provider entry point
// -----------------------------------------------------------------------------
func main() {
	providerserver.Serve(context.Background(), provider.NewProvider, providerserver.ServeOpts{})
}
