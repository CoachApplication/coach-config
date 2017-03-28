package provider

import (
	base_config "github.com/james-nesbitt/coach-config"
)

type Provider interface {
	Scopes() []string
	Keys() []string
	Get(key, scope string) (base_config.Config, error)
}
