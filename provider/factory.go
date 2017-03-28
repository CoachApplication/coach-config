package provider

import (
	base_config "github.com/CoachApplication/config"
)

// Factory generates Config values from Key and Scope pairs
type Factory interface {
	// Get a single Config for a matching key-scope pair
	Get(key, scope string) (base_config.Config, error)
}
