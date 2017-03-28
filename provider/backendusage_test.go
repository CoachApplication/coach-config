package provider

import (
	"testing"
)

func TestAllBackendUsage_Handles(t *testing.T) {
	us := BackendUsage(&AllBackendUsage{})

	if !us.Handles("one", "two") {
		t.Error("AllBackendUsage said that it doesn't handle a specific key-scope pair")
	}
	if !us.Handles(BACKENDUSAGE_KEY_ALL, "two") {
		t.Error("AllBackendUsage said that it doesn't handle a specific scope")
	}
	if !us.Handles("one", BACKENDUSAGE_SCOPE_ALL) {
		t.Error("AllBackendUsage said that it doesn't handle a specific key")
	}
}
