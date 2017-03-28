package json

import (
	"encoding/json"
	"io"

	"bytes"
	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
	base_config "github.com/CoachApplication/coach-config"
	base_config_provider "github.com/CoachApplication/coach-config/provider"
)

type Config struct {
	connector base_config_provider.Connector
	key       string
	scope     string
}

func NewConfig(key, scope string, con base_config_provider.Connector) *Config {
	return &Config{
		key:       key,
		scope:     scope,
		connector: con,
	}
}

func (jc *Config) Config() base_config.Config {
	return base_config.Config(jc)
}

// Marshall gets a configuration and apply it to a target struct
func (jc *Config) Get(target interface{}) api.Result {
	res := base.NewResult()

	go func(key, scope string) {
		if rc, err := jc.connector.Get(key, scope); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			d := json.NewDecoder(rc)
			if err := d.Decode(target); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else {
				res.MarkSucceeded()
			}
			rc.Close()
		}
		res.MarkFinished()
	}(jc.key, jc.scope)

	return res.Result()
}

// UnMarshall sets a Config value by converting a passed struct into a configuration
// The expects that the values assigned are permanently saved
func (jc *Config) Set(source interface{}) api.Result {
	res := base.NewResult()

	go func(key, scope string) {
		defer res.MarkFinished()
		// @TODO should we do this without holding all the bytes in plain memory?
		if b, err := json.Marshal(source); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			rc := io.ReadCloser(&readCloserWrapper{Reader: bytes.NewBuffer(b)})
			if err := jc.connector.Set(key, scope, rc); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else {
				res.MarkSucceeded()
			}
		}
	}(jc.key, jc.scope)

	return res.Result()
}

// A simple struct that wraps an io.Reader and adds a Close() to make it a io.ReadCloser
type readCloserWrapper struct{ io.Reader }
// Close the io.Closer interface method
func (rcw *readCloserWrapper) Close() error { return nil }
