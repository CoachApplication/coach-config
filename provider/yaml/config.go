package yaml

import (
	"bytes"
	yaml "gopkg.in/yaml.v2"
	"io"

	api "github.com/james-nesbitt/coach-api"
	base "github.com/james-nesbitt/coach-base"
	base_config "github.com/james-nesbitt/coach-config"
	base_config_provider "github.com/james-nesbitt/coach-config/provider"
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
