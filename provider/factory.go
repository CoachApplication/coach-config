package provider

import (
	base_config "github.com/CoachApplication/coach-config"
)

type Factory interface {
	Get(key, scope string) (base_config.Config, error)
}
