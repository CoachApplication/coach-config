package provider

import (
	"fmt"
	base_config "github.com/CoachApplication/config"
)

/**
 *
 */

type BackendConfigProvider struct {
	backends []Backend
}

func NewBackendConfigProvider() *BackendConfigProvider {
	return &BackendConfigProvider{
		backends: []Backend{},
	}
}

func (bcp *BackendConfigProvider) Provider() Provider {
	return Provider(bcp)
}

func (bcp *BackendConfigProvider) Add(backend Backend) {
	bcp.backends = append(bcp.backends, backend)
}

func (bcp *BackendConfigProvider) Scopes() []string {
	scopes := uniqueStringSlice{}
	for _, backend := range bcp.backends {
		scopes.merge(backend.Scopes())
	}
	return scopes.slice()
}

func (bcp *BackendConfigProvider) Keys() []string {
	keys := uniqueStringSlice{}
	for _, backend := range bcp.backends {
		keys.merge(backend.Keys())
	}
	return keys.slice()
}

func (bcp *BackendConfigProvider) Get(key, scope string) (base_config.Config, error) {
	for _, backend := range bcp.backends {
		if backend.Handles(key, scope) {
			return backend.Get(key, scope)
		}
	}
	return nil, error(ConfigNotHandlerdError{Key: key, Scope: scope})
}

type uniqueStringSlice struct {
	s []string
}

func (uss *uniqueStringSlice) add(val string) {
	for _, has := range uss.s {
		if has == val {
			return
		}
	}
	uss.s = append(uss.s, val)
}
func (uss *uniqueStringSlice) merge(vals []string) {
	for _, val := range vals {
		uss.add(val)
	}
}
func (uss *uniqueStringSlice) slice() []string {
	return uss.s
}

type ConfigNotHandlerdError struct {
	Key   string
	Scope string
}

func (cnhe ConfigNotHandlerdError) Error() string {
	return fmt.Sprintf("No backend handles the requests key/scope pair: %s / %s", cnhe.Key, cnhe.Scope)
}
