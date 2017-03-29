package buffered

import (
	"bytes"
	"testing"
)

func MakeTestMultiple(t *testing.T) *Multiple {
	return &Multiple{
		vals: map[string]map[string][]byte{
			"one": {
				"A": []byte("A-one"),
				"B": []byte("B-one"),
			},
			"two": {
				"B": []byte("B-two"),
				"C": []byte("C-two"),
			},
			"three": {
				"C": []byte("C-three"),
				"D": []byte("D-three"),
			},
		},
	}
}

func TestMultiple_Keys(t *testing.T) {
	m := MakeTestMultiple(t)

	ks := m.Keys()
	if len(ks) == 0 {
		t.Error("Multiple.Keys() returned no keys")
	} else if len(ks) != 3 {
		t.Error("Multiple.Keys() returned wrong number of keys")
	} else {
		for _, k := range []string{"one", "two", "three"} {
			if !valIsInSlice(k, ks) {
				t.Error("Multiple.Keys() is missing a key: ", k)
			}
		}
	}
}

func TestMultiple_Scopes(t *testing.T) {
	m := MakeTestMultiple(t)

	ss := m.Scopes()
	if len(ss) == 0 {
		t.Error("Multiple.Scopes() returned no scopes")
	} else if len(ss) != 4 {
		t.Error("Multiple.Scopes() returned wrong number of scopes: ", ss)
	} else {
		for _, s := range []string{"A", "B", "C", "D"} {
			if !valIsInSlice(s, ss) {
				t.Error("Multiple.Scopes() is missing a scope: ", s)
			}
		}
	}
}

func TestMultiple_Get(t *testing.T) {
	m := MakeTestMultiple(t)

	if _, err := m.Get("no", "no"); err == nil {
		t.Error("Multiple.Get() returned no error on invalid key-scope pair")
	}
	if rc, err := m.Get("two", "C"); err != nil {
		t.Error("Multiple.Get() returned an error on valid key-scope pair: ", err.Error())
	} else {
		b := bytes.NewBufferString("")
		b.ReadFrom(rc)
		rc.Close()
		val := b.String()
		if val != "C-two" {
			t.Error("Multiple.Get() reader has the wrong string: ", val)
		}
	}
}

func TestMultiple_Add(t *testing.T) {
	m := MakeTestMultiple(t)

	m.Add("one", "A", "one-A-changes")

	if rc, err := m.Get("one", "A"); err != nil {
		t.Error("Multiple.Get() returned an error on valid key-scope pair after change: ", err.Error())
	} else {
		b := bytes.NewBufferString("")
		b.ReadFrom(rc)
		rc.Close()
		val := b.String()
		if val != "one-A-changes" {
			t.Error("Multiple.Get() reader has the wrong string after change: ", val)
		}
	}
	m.Add("three", "E", "three-E")

	if rc, err := m.Get("three", "E"); err != nil {
		t.Error("Multiple.Get() returned an error on valid key-scope pair: ", err.Error())
	} else {
		b := bytes.NewBufferString("")
		b.ReadFrom(rc)
		rc.Close()
		val := b.String()
		if val != "three-E" {
			t.Error("Multiple.Get() reader has the wrong string: ", val)
		}
	}
}

// is a value in a slice
func valIsInSlice(val string, slice []string) bool {
	for _, each := range slice {
		if each == val {
			return true
		}
	}
	return false
}
