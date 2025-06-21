// *******************************************************
// terraform-provider-umbrella – initial fully‑working MVP
// -------------------------------------------------------
// This codebase contains a minimal yet complete, *compilable* Terraform
// provider that manages **Internal Domains** in Cisco Umbrella.  It
// demonstrates the core patterns (provider wiring, authenticated API
// client, resource CRUD, data/attribute mapping, and safe unit tests with
// HTTP mocks).  Further Umbrella domains (Tunnels, Networks, Destination
// Lists, SAML, …) can be added by following the same file‑per‑resource
// pattern and re‑using the shared apiClient.
// *******************************************************

// -------------------- go.mod ---------------------------

// -------------------------------------------------------

// -------------------- main.go --------------------------
package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/mantisec/terraform-provider-umbrella/internal/provider"
)

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/mantisec/umbrella",
	})
}
