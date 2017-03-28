package configuration

// ScopedConfig An object which can provide scoped config objects
type ScopedConfig interface {
	// Get a Config for a scope
	Get(scope string) (Config, error)
	// Set uses a passed Config to set a value to a scope
	Set(scope string, config Config) error
	// List available scopes
	Order() []string
}
