package json

import (
	base_config "github.com/CoachApplication/config"
	base_config_provider "github.com/CoachApplication/config/provider"
)

type Factory struct {
	connector base_config_provider.Connector
}

// NewFactory Constructor for Factory
func NewFactory(con base_config_provider.Connector) *Factory {
	return &Factory{
		connector: con,
	}
}

func (f *Factory) Factory() base_config_provider.Factory {
	return base_config_provider.Factory(f)
}

func (f *Factory) Get(key, scope string) (base_config.Config, error) {
	return (&Config{
		connector: f.connector,
		key:       key,
		scope:     scope,
	}).Config(), nil
}
