package provider

import (
	base_config "github.com/james-nesbitt/coach-config"
)

type Factory interface {
	Get(key, scope string) (base_config.Config, error)
}
