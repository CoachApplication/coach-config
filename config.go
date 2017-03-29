package config

import (
	api "github.com/CoachApplication/api"
)

// Config encapsulation of configuration
type Config interface {
	// HasValue does this Config already have data, use the api.Result Success() to indicate that a value exists
	HasValue() api.Result
	// Marshall gets a configuration and apply it to a target struct
	Get(interface{}) api.Result
	// UnMarshall sets a Config value by converting a passed struct into a configuration
	// The expects that the values assigned are permanently saved
	Set(interface{}) api.Result
}
