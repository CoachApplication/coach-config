package json

import (
	"encoding/json"
	"io"

	"bytes"
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	base_config "github.com/CoachApplication/config"
	base_config_provider "github.com/CoachApplication/config/provider"
	utils "github.com/CoachApplication/utils"
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

func (jc *Config) HasValue() bool {
	return jc.connector.HasValue(jc.key, jc.scope)
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
			rc := io.ReadCloser(utils.CloseDecorateReader(bytes.NewBuffer(b), nil))
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
