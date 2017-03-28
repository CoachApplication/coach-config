package configuration

import (
	"fmt"
)

// StandardScopedConfig A standard implementation of a ScopedConfig that just maintains an ordered map list
type StandardScopedConfig struct {
	cMap   map[string]Config
	cOrder []string
}

// NewStandardScopedConfig Constructor for StandardScopedConfig
func NewStandardScopedConfig() *StandardScopedConfig {
	return &StandardScopedConfig{}
}

// ScopedConfig Explicitly convert this struct to a config ScopedConfig interface
func (sc *StandardScopedConfig) ScopedConfig() ScopedConfig {
	return ScopedConfig(sc)
}

// Get a Config for a scope
func (sc *StandardScopedConfig) Get(scope string) (Config, error) {
	sc.safe()
	config, exists := sc.cMap[scope]
	if exists {
		return config, nil
	} else {
		return config, error(ConfigScopeNotFoundError{Scope: scope})
	}
}

// Set uses a passed Config to set a value to a scope
func (sc *StandardScopedConfig) Set(scope string, config Config) error {
	sc.safe()
	if _, exists := sc.cMap[scope]; !exists {
		sc.cOrder = append(sc.cOrder, scope)
	}
	sc.cMap[scope] = config
	return nil
}

// List available scopes
func (sc *StandardScopedConfig) Order() []string {
	sc.safe()
	return sc.cOrder
}

func (sc *StandardScopedConfig) safe() {
	if sc.cMap == nil {
		sc.cMap = map[string]Config{}
		sc.cOrder = []string{}
	}
}

type ConfigScopeNotFoundError struct {
	Scope string
}

func (csnfe ConfigScopeNotFoundError) Error() string {
	return fmt.Sprintf("Config was not found at the reqyested scope %s", csnfe.Scope)
}
