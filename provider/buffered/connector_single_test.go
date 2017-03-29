package buffered_test

import (
	"bytes"
	"github.com/CoachApplication/config/provider/buffered"
	"github.com/CoachApplication/utils"
	"io"
	"testing"
)

func TestSingle_Keys(t *testing.T) {
	bc := buffered.NewSingle("key", "scope", []byte{})

	k := bc.Keys()
	if len(k) != 1 {
		t.Error("Single return the wrong number of ket", k)
	} else if k[0] != "key" {
		t.Error("Single return incorrect key values")
	}
}

func TestSingle_Scopes(t *testing.T) {
	bc := buffered.NewSingle("key", "scope", []byte{})

	s := bc.Scopes()
	if len(s) != 1 {
		t.Error("Single return the wrong number of scopes", s)
	} else if s[0] != "scope" {
		t.Error("Single return incorrect scope values")
	}
}

func TestSingle_Get(t *testing.T) {
	bc := buffered.NewSingle("key", "scope", []byte("test"))

	if _, err := bc.Get("no", "no"); err == nil {
		t.Error("BufferredSingle did not return an error when incorrect key-scope was requested")
	} else if rc, err := bc.Get("key", "scope"); err != nil {
		t.Error("BufferredSingle returned an error when asked for a valid key-scope pair")
	} else {
		b := bytes.NewBufferString("")
		b.ReadFrom(rc)
		rc.Close()
		val := b.String()
		if val != "test" {
			t.Error("Single reader has the wrong string: ", val)
		}
	}
}

func TestSingle_Set(t *testing.T) {
	bc := buffered.NewSingle("key", "scope", []byte("one"))

	rc := io.ReadCloser(utils.CloseDecorateReader(bytes.NewBufferString("two"), nil))
	if err := bc.Set("key", "scope", rc); err != nil {
		t.Error("Single returned an error when setting a new value: ", err.Error())
	} else {
		if rc, err := bc.Get("key", "scope"); err != nil {
			t.Error("BufferredSingle returned an error when asked for a valid key-scope pair")
		} else {
			b := bytes.NewBufferString("")
			b.ReadFrom(rc)
			rc.Close()
			val := b.String()
			if val != "two" {
				t.Error("Single reader has the wrong string: ", val)
			} else if val == "one" {
				t.Error("Single reader has the wrong string, it still has it's original value: ", val)
			}
		}
	}
}
