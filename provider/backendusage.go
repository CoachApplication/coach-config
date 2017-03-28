package provider

const (
	// Wildcard Key
	BACKENDUSAGE_KEY_ALL = "--all--"
	// Wildcard Scope
	BACKENDUSAGE_SCOPE_ALL = "--all--"
)

type BackendUsage interface {
	Handles(key, scope string) bool
}

type AllBackendUsage struct{}

func (abu *AllBackendUsage) BackendUsage() BackendUsage {
	return BackendUsage(abu)
}

func (abu *AllBackendUsage) Handles(key, scope string) bool {
	return true
}
