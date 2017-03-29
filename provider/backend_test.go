package provider_test

import (
	"errors"
	"reflect"
	"testing"

	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	config "github.com/CoachApplication/config"
	config_provider "github.com/CoachApplication/config/provider"
)

/**
 * This file contains a TestBackend struct that can be used for other tests, and it had tests for that backend
 * This file contains no real system tests, just code to assist other testing.
 */

func TestBackend_Handles(t *testing.T) {
	tb := NewTestBackend(t, "key", "scope", "one").Backend()

	if tb.Handles("no", "no") {
		t.Error("TestBackend says that it handles key-scopes that it shouldn't")
	}
	if tb.Handles("key", "no") {
		t.Error("TestBackend says that it handles key-scopes that it shouldn't")
	}
	if tb.Handles("no", "scope") {
		t.Error("TestBackend says that it handles key-scopes that it shouldn't")
	}
	if !tb.Handles("key", "scope") {
		t.Error("TestBackend says that it doesn't handle key-scopes that it should")
	}
}

func TestBackend_Scopes(t *testing.T) {
	tb := NewTestBackend(t, "key", "scope", "one").Backend()

	s := tb.Scopes()
	if len(s) != 1 {
		t.Error("TestBackend returned the wrong number of scopes")
	} else if s[0] != "scope" {
		t.Error("TestBackend provided the wrong scope in Scopes()")
	}
}

func TestBackend_Keys(t *testing.T) {
	tb := NewTestBackend(t, "key", "scope", "one").Backend()

	k := tb.Keys()
	if len(k) != 1 {
		t.Error("TestBackend returned the wrong number of keys")
	} else if k[0] != "key" {
		t.Error("TestBackend provided the wrong key in Keys()")
	}
}

func TestBackend_Get(t *testing.T) {
	tb := NewTestBackend(t, "key", "scope", "one").Backend()

	if _, err := tb.Get("no", "no"); err == nil {
		t.Error("TestBackend did not return an error when it was asked for a Config that it shouldn't provide.")
	}
	if c, err := tb.Get("key", "scope"); err != nil {
		t.Error("TestBackend returned an error when it was asked for a Config that it should provide.")
	} else {
		s := "zero"
		c.Get(&s)
		if s != "one" {
			t.Error("TestBackend provided the wrong Config, or the Config failed to assign the correct string value.")
		}
	}
}

// TestBackend a Backend that returns a TestStringConfig for a provided key-scope pair
type TestBackend struct {
	t     *testing.T
	key   string
	scope string
	val   string
}

// NewTestBackend TestBackend Constructor
func NewTestBackend(t *testing.T, key string, scope string, val string) *TestBackend {
	return &TestBackend{
		t:     t,
		key:   key,
		scope: scope,
		val:   val,
	}
}

// Backend Explicitly convert this to a Backend interface
func (tb *TestBackend) Backend() config_provider.Backend {
	return config_provider.Backend(tb)
}

// Handles answers whether or not this Backend handles that passed key-scope pair
func (tb *TestBackend) Handles(key, scope string) bool {
	return key == tb.key && scope == tb.scope
}

// Scopes provide a list of all scopes that this backend handles
func (tb *TestBackend) Scopes() []string {
	return []string{tb.scope}
}

// Keys provide a list of all keys that this backend handles
func (tb *TestBackend) Keys() []string {
	return []string{tb.key}
}

// Get retrieve a Config for a key-scope pair
func (tb *TestBackend) Get(key, scope string) (config.Config, error) {
	if key == tb.key && scope == tb.scope {
		return NewTestStringConfig(tb.t, tb.key+"-"+tb.scope, tb.val).Config(), nil
	} else {
		return nil, errors.New("Incorrect key/scope requested")
	}
}

/**
 * This is copied directly from the config - there are tests there
 */

// TestConfig A testing Config implementation
type TestStringConfig struct {
	t   *testing.T
	id  string
	val string
}

// NewTestStringConfig TestStringConfig constructor
func NewTestStringConfig(t *testing.T, id, val string) *TestStringConfig {
	return &TestStringConfig{
		t:   t,
		id:  id,
		val: val,
	}
}

// Config explicitly convert this to an Config interface
func (tsc *TestStringConfig) Config() config.Config {
	return config.Config(tsc)
}

// HasValue indicates if the Config already has a value
func (tsc *TestStringConfig) HasValue() api.Result {
	if tsc.val != "" {
		return base.MakeSuccessfulResult()
	} else {
		return base.MakeFailedResult()
	}
}

// Marshall gets a configuration and apply it to a target struct
func (tsc *TestStringConfig) Get(target interface{}) api.Result {
	res := base.NewResult()
	if _, success := target.(*string); success {
		res.MarkSucceeded()
		tv := reflect.ValueOf(target)
		tsc.t.Log("reflect:", tv.Type(), tv.String())
		if tv.Kind() != reflect.Ptr || tv.IsNil() {
			res.AddError(errors.New("Invalid target"))
			tsc.t.Error("Not a good pointer")
		} else {
			tv.Elem().SetString(tsc.val)
		}
	} else {
		res.MarkFailed()
		res.AddError(errors.New("Incorrect val type"))
	}

	res.MarkFinished()
	return res.Result()
}

// UnMarshall sets a Config value by converting a passed struct into a configuration
// The expects that the values assigned are permanently saved
func (tsc *TestStringConfig) Set(source interface{}) api.Result {
	res := base.NewResult()
	if _, success := source.(string); success {
		res.MarkSucceeded()
		sv := reflect.ValueOf(source)
		tsc.val = sv.String()
		tsc.t.Log("reflect:", tsc.val, sv.Type(), sv.String())
	} else {
		res.MarkFailed()
		res.AddError(errors.New("Incorrect val type"))
	}
	res.MarkFinished()
	return res.Result()
}
