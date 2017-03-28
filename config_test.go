package configuration

import (
	"errors"
	"reflect"
	"testing"

	api "github.com/james-nesbitt/coach-api"
	base "github.com/james-nesbitt/coach-base"
)

/**
 * Here we test the accessors on our test config.  It is not a system test, but rather jsut a test to make sure that
 * out testing struct behaves properly.
 */

func Test_TestStringConfig_Get(t *testing.T) {
	c := NewTestStringConfig(t, "first", "one")

	var s string = "zero"
	res := c.Get(&s)
	<-res.Finished()

	if !res.Success() {
		t.Error("TestConfig Get returned failed status")
		for _, err := range res.Errors() {
			t.Error(err.Error())
		}
	}

	if s != "one" {
		t.Errorf("StringConfig did not retrieve the correct value to our test string : %s", s)
	}
}

func Test_TestStringConfig_Set(t *testing.T) {
	var res api.Result

	c := NewTestStringConfig(t, "first", "one")
	str := "two"
	res = c.Set(str)
	<-res.Finished()

	if !res.Success() {
		for _, err := range res.Errors() {
			t.Error(err.Error())
		}
		t.Error("TestConfig Set returned failed status")
	}

	var s string = "zero"
	res = c.Get(&s)
	<-res.Finished()

	if !res.Success() {
		for _, err := range res.Errors() {
			t.Error(err.Error())
		}
		t.Error("TestConfig Get returned failed status")
	}

	if s == "one" {
		t.Error("StringConfig did not apply the correct value to our test string, it still had it's constructor value", s)
	} else if s != "two" {
		t.Errorf("StringConfig did not apply the correct value to our test string : %s", s)
	}
}

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
func (tsc *TestStringConfig) Config() Config {
	return Config(tsc)
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
