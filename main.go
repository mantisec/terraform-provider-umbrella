package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/mantisec/terraform-provider-umbrella/internal/provider"
)

// These will be set by the goreleaser configuration
// to appropriate values for the compiled binary.
var version string = "dev"
var commit string = ""

// -----------------------------------------------------------------------------
// Provider entry point
// -----------------------------------------------------------------------------
func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/mantisec/umbrella",
		Debug:   debugMode,
	}

	if debugMode {
		fmt.Printf("Starting terraform-provider-umbrella %s (commit: %s) in debug mode\n", version, commit)
	}

	err := providerserver.Serve(context.Background(), provider.NewProvider, opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
