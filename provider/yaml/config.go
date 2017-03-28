package yaml

import (
	"bytes"
	"io"

	yaml "gopkg.in/yaml.v2"

	api "github.com/CoachApplication/coach-api"
	base "github.com/CoachApplication/coach-base"
	base_config "github.com/CoachApplication/coach-config"
	base_config_provider "github.com/CoachApplication/coach-config/provider"
)

// YamlConfig Build Config by marshalling yml from a connector
type Config struct {
	key       string
	scope     string
	connector base_config_provider.Connector
}

func NewConfig(key, scope string, con base_config_provider.Connector) *Config {
	return &Config{
		key:       key,
		scope:     scope,
		connector: con,
	}
}

// Marshall gets a configuration and apply it to a target struct
func (ycc *Config) Config() base_config.Config {
	return base_config.Config(ycc)
}

// Marshall gets a configuration and apply it to a target struct
func (ycc *Config) Get(target interface{}) api.Result {

	res := base.NewResult()

	go func(key, scope string, con base_config_provider.Connector) {
		defer res.MarkFinished()

		if r, err := con.Get(key, scope); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			defer r.Close()
			buf := bytes.Buffer{}
			if _, err := buf.ReadFrom(r); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else if err := yaml.Unmarshal(buf.Bytes(), target); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else {
				res.MarkSucceeded()
			}
		}
	}(ycc.key, ycc.scope, ycc.connector)

	return res.Result()
}

// UnMarshall sets a Config value by converting a passed struct into a configuration
// The expects that the values assigned are permanently saved
func (ycc *Config) Set(source interface{}) api.Result {
	res := base.NewResult()

	go func(key, scope string, con base_config_provider.Connector) {
		defer res.MarkFinished()
		// @TODO should we do this without holding all the bytes in plain memory?
		if b, err := yaml.Marshal(source); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			rc := io.ReadCloser(&readCloserWrapper{Reader: bytes.NewBuffer(b)})
			if err := con.Set(key, scope, rc); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else {
				res.MarkSucceeded()
			}
		}
	}(ycc.key, ycc.scope, ycc.connector)

	return res.Result()
}

// A simple struct that wraps an io.Reader and adds a Close() to make it a io.ReadCloser
type readCloserWrapper struct{ io.Reader }

// Close the io.Closer interface method
func (rcw *readCloserWrapper) Close() error { return nil }
