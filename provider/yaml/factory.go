package yaml

import (
	base_config "github.com/CoachApplication/coach-config"
	base_config_provider "github.com/CoachApplication/coach-config/provider"
)

/**
 * Here we have the config provider related architecture that will convert yaml files to config
 */

// Factory a config provider backend that can make Config objects from a connector
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

	/** @TODO should we check if the path exists first?  If we do then you can't write to new configs? */

	return NewConfig(key, scope, f.connector).Config(), nil
}
