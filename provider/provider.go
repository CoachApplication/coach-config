package provider

import (
	base_config "github.com/CoachApplication/config"
)

type Provider interface {
	Scopes() []string
	Keys() []string
	Get(key, scope string) (base_config.Config, error)
}
