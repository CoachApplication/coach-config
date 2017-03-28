package json

import (
	"testing"

	coach_config_provider "github.com/CoachApplication/coach-config/provider"
	coach_config_provider_bufferred "github.com/CoachApplication/coach-config/provider/bufferred"
)

// @TODO add more complex JSON data to test

// Message JSON bytes
var MessageBytes = []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
var MessageStruct = Message{
	Name: "Alice",
	Body: "Hello",
	Time: 1294706395881547000,
}

// Message struct
type Message struct {
	Name string
	Body string
	Time int64
}

func TestJsonConfig_Get(t *testing.T) {
	jb := &Backend{
		connector: coach_config_provider_bufferred.NewBufferedConnector("key", "scope", MessageBytes),
		usage:     &coach_config_provider.AllBackendUsage{},
	}

	if !jb.Handles("key", "scope") {
		t.Error("Backend didn't Handle() properly")
	} else if c, err := jb.Get("key", "scope"); err != nil {
		t.Error("Backend gave an error when retreiving valid key-scope Config")
	} else {
		var m Message
		res := c.Get(&m)
		<-res.Finished()

		if !res.Success() {
			t.Error("Backend Config reported failure in Get() : ", res.Errors())
		} else {

			if m.Name != MessageStruct.Name {
				t.Error("Backend provided incorrect data ==> Name : ", m.Name)
			}
			if m.Body != MessageStruct.Body {
				t.Error("Backend provided incorrect data ==> Body : ", m.Body)
			}
			if m.Time != MessageStruct.Time {
				t.Error("Backend provided incorrect data ==> Time : ", m.Time)
			}

		}
	}
}

func TestJsonConfig_Set(t *testing.T) {
	con := coach_config_provider_bufferred.NewBufferedConnector("key", "scope", []byte{})
	jb := &Backend{
		connector: con,
		usage:     &coach_config_provider.AllBackendUsage{},
	}

	if !jb.Handles("key", "scope") {
		t.Error("Backend didn't Handle() properly")
	} else if c, err := jb.Get("key", "scope"); err != nil {
		t.Error("Backend gave an error when retreiving valid key-scope Config")
	} else {
		res := c.Set(MessageStruct)
		//<-res.Finished()

		if !res.Success() {
			t.Error("Backend Config reported a failure when Set(): ", res.Errors())
		} else {

		}
	}
}
