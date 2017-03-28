package json

import (
	"encoding/json"
	"io"

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

func (jc *Config) Config() base_config.Config {
	return base_config.Config(jc)
}

// Marshall gets a configuration and apply it to a target struct
func (jc *Config) Get(target interface{}) api.Result {
	res := base.NewResult()

	go func(key, scope string) {
		defer res.MarkFinished()
		if rc, err := jc.connector.Get(key, scope); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			defer rc.Close()
			d := json.NewDecoder(rc)
			if err := d.Decode(target); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else {
				res.MarkSucceeded()
			}
		}
	}(jc.key, jc.scope)

	return res.Result()
}

// UnMarshall sets a Config value by converting a passed struct into a configuration
// The expects that the values assigned are permanently saved
func (jc *Config) Set(source interface{}) api.Result {
	res := base.NewResult()

	go func(key, scope string) {
		defer res.MarkFinished()

		r, w := io.Pipe() // technically r is a ReaderCloser
		defer w.Close()   // this we do to be responsible
		defer r.Close()   // this we do in case the connector isn't responsible

		e := json.NewEncoder(w)
		if err := e.Encode(source); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			if err := jc.connector.Set(key, scope, r); err != nil {
				res.AddError(err)
				res.MarkFailed()
			} else {
				res.MarkSucceeded()
			}
		}
	}(jc.key, jc.scope)

	return res.Result()
}
