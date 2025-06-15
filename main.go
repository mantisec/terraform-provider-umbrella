package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/mantisec/terraform-provider-umbrella/internal/provider"
)

// -----------------------------------------------------------------------------
// Provider entry point
// -----------------------------------------------------------------------------
func main() {
	err := providerserver.Serve(context.Background(), provider.NewProvider, providerserver.ServeOpts{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
