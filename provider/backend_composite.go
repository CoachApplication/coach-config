package provider

import (
	base_config "github.com/james-nesbitt/coach-config"
)

/**
 * A backend that relies on a combination of a Backend, a Connector and a Usage struct to fulfill the Backend interface
 *
 * @NOTE there is no need to test this as all it does is hand off all methods to other structs.
 */

// Backend a config provider backend that can make Config objects from a
type CompositeBackend struct {
	connector Connector
	factory   Factory
	usage     BackendUsage
}

// NewCompositeBackend Constructor for
func NewCompositeBackend(con Connector, us BackendUsage, fac Factory) *CompositeBackend {
	return &CompositeBackend{
		connector: con,
		usage:     us,
		factory:   fac,
	}
}

func (fcb *CompositeBackend) Backend() Backend {
	return Backend(fcb)
}

func (fcb *CompositeBackend) Handles(key, scope string) bool {
	return fcb.usage.Handles(key, scope)
}
func (fcb *CompositeBackend) Scopes() []string {
	return fcb.connector.Scopes()
}
func (fcb *CompositeBackend) Keys() []string {
	return fcb.connector.Keys()
}
func (fcb *CompositeBackend) Get(key, scope string) (base_config.Config, error) {
	c, err := fcb.factory.Get(key, scope)
	return c, err
}
