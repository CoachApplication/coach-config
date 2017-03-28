package configuration

import (
	api "github.com/james-nesbitt/coach-api"
)

// Config encapsulation of configuration
type Config interface {
	// Marshall gets a configuration and apply it to a target struct
	Get(interface{}) api.Result
	// UnMarshall sets a Config value by converting a passed struct into a configuration
	// The expects that the values assigned are permanently saved
	Set(interface{}) api.Result
}
