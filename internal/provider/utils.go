package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// -----------------------------------------------------------------------------
// Shared utility functions
// -----------------------------------------------------------------------------

func setToStringSlice(ctx context.Context, v types.Set, diags *diag.Diagnostics) []string {
	if v.IsNull() || v.IsUnknown() {
		return []string{}
	}
	var out []string
	diags.Append(v.ElementsAs(ctx, &out, false)...)
	return out
}

func diffSlices(old, new []string) (toAdd, toDel []string) {
	want := map[string]struct{}{}
	have := map[string]struct{}{}
	for _, n := range new {
		want[n] = struct{}{}
	}
	for _, o := range old {
		have[o] = struct{}{}
	}
	for d := range want {
		if _, ok := have[d]; !ok {
			toAdd = append(toAdd, d)
		}
	}
	for d := range have {
		if _, ok := want[d]; !ok {
			toDel = append(toDel, d)
		}
	}
	return
}

// Helper function to compare string slices
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := make(map[string]bool)
	for _, v := range a {
		aMap[v] = true
	}
	for _, v := range b {
		if !aMap[v] {
			return false
		}
	}
	return true
}

// Helper function to convert string slice to attr.Value slice
func stringSliceToAttrValues(slice []string) []attr.Value {
	elems := []attr.Value{}
	for _, s := range slice {
		elems = append(elems, types.StringValue(s))
	}
	return elems
}
