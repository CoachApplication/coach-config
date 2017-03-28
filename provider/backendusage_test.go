package provider_test

import (
	"testing"

	config_provider "github.com/CoachApplication/coach-config/provider"
)

func TestAllBackendUsage_Handles(t *testing.T) {
	us := config_provider.BackendUsage(&config_provider.AllBackendUsage{})

	if !us.Handles("one", "two") {
		t.Error("AllBackendUsage said that it doesn't handle a specific key-scope pair")
	}
	if !us.Handles(config_provider.BACKENDUSAGE_KEY_ALL, "two") {
		t.Error("AllBackendUsage said that it doesn't handle a specific scope")
	}
	if !us.Handles("one", config_provider.BACKENDUSAGE_SCOPE_ALL) {
		t.Error("AllBackendUsage said that it doesn't handle a specific key")
	}
}
