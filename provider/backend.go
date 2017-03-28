package provider

import (
	base_config "github.com/james-nesbitt/coach-config"
)

// Backend Provider consumed struct that handles retrieving Config for certain key-scope combinations
type Backend interface {
	// Handles answers whether or not this Backend handles that passed key-scope pair
	Handles(key, scope string) bool
	// Scopes provide a list of all scopes that this backend handles
	Scopes() []string
	// Keys provide a list of all keys that this backend handles
	Keys() []string
	// Get retrieve a Config for a key-scope pair
	Get(key, scope string) (base_config.Config, error)
}
