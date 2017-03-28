package provider_test

import (
	"testing"

	config_provider "github.com/CoachApplication/coach-config/provider"
	utils "github.com/CoachApplication/coach-utils"
)

func testBackendConfigProvider(t *testing.T) *config_provider.BackendConfigProvider {
	tbcp := config_provider.NewBackendConfigProvider()
	tbcp.Add(NewTestBackend(t, "A", "B", "AB").Backend())
	tbcp.Add(NewTestBackend(t, "A", "C", "AC").Backend())
	tbcp.Add(NewTestBackend(t, "B", "D", "BD").Backend())
	tbcp.Add(NewTestBackend(t, "B", "D", "ED").Backend())
	return tbcp
}

func TestBackendConfigProvider_Scopes(t *testing.T) {
	tbcp := testBackendConfigProvider(t)

	ss := tbcp.Scopes()

	if len(ss) != 3 {
		t.Error("BackendConfigProvider gave incorrect number of Scope values ", ss, len(ss))
	}
	if !sliceHasValue(ss, "B") {
		t.Error("BackendConfigProvider Scopes() is missing a value")
	}
	if !sliceHasValue(ss, "C") {
		t.Error("BackendConfigProvider Scopes() is missing a value")
	}
	if !sliceHasValue(ss, "D") {
		t.Error("BackendConfigProvider Scopes() is missing a value")
	}
	if sliceHasValue(ss, "A") {
		t.Error("BackendConfigProvider Scopes() has a value that it shouldn't have")
	}
	if sliceHasValue(ss, "E") {
		t.Error("BackendConfigProvider Scopes() has a value that it shouldn't have")
	}
}

func TestBackendConfigProvider_Keys(t *testing.T) {
	tbcp := testBackendConfigProvider(t)

	ks := tbcp.Keys()

	if len(ks) != 2 {
		t.Error("BackendConfigProvider gave incorrect number of Keys values", ks, len(ks))
	}
	if !sliceHasValue(ks, "A") {
		t.Error("BackendConfigProvider Keys() is missing a value")
	}
	if !sliceHasValue(ks, "B") {
		t.Error("BackendConfigProvider Keys() is missing a value")
	}
	if sliceHasValue(ks, "D") {
		t.Error("BackendConfigProvider Keys() has a value that it shouldn't have")
	}
	if sliceHasValue(ks, "E") {
		t.Error("BackendConfigProvider Keys() has a value that it shouldn't have")
	}
}

func TestBackendConfigProvider_Get(t *testing.T) {
	tbcp := testBackendConfigProvider(t)

	if _, err := tbcp.Get("B", "A"); err == nil {
		t.Error("BackendConfigProvider gave a Config when it shouldn't have")
	}
	if _, err := tbcp.Get("E", "E"); err == nil {
		t.Error("BackendConfigProvider gave a Config when it shouldn't have")
	}

	if c, err := tbcp.Get("A", "C"); err != nil {
		t.Error("BackendConfigProvider didn't give a Config when it should have", err.Error())
	} else {
		s := "zero"
		c.Get(&s)
		if s != "AC" {
			t.Error("BackendConfigProvider provided the wrong Config, or the Config provides the wrong string.", s)
		}
	}
}

//
func Test_sliceHasValue(t *testing.T) {
	if sliceHasValue([]string{"A", "B", "C"}, "0") {
		t.Error("sliceHasValue improperly recognized a key")
	}
	if !sliceHasValue([]string{"A", "B", "C"}, "B") {
		t.Error("sliceHasValue improperly rejected a key")
	}
}

// check if a value is in a slice
func sliceHasValue(s []string, val string) bool {
	for _, sval := range s {
		if sval == val {
			return true
		}
	}
	return false
}

// We use the unique string slice in the backend provider, so we should test that it works
func Test_uniqueStringSlice(t *testing.T) {
	s := utils.UniqueStringSlice{}

	// test adding strings
	s.Add("A")
	sl1 := s.Slice()
	if len(sl1) != 1 {
		t.Error("uniqueStringSlice did not properly add a value")
	} else if sl1[0] != "A" {
		t.Error("uniqueStringSlice did not properly add a value")
	}

	// test adding further
	s.Add("B")
	s.Add("C")
	sl2 := s.Slice()
	if len(sl2) != 3 {
		t.Error("uniqueStringSlice did not properly add a value")
	} else if sl2[0] != "A" || sl2[1] != "B" || sl2[2] != "C" {
		t.Error("uniqueStringSlice did not properly add a value")
	}

	// test unique values
	s.Add("A")
	sl3 := s.Slice()
	if len(sl3) != 3 {
		t.Error("uniqueStringSlice improperly added a duplicate value")
	}
}
