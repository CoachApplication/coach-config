package json_test

import (
	"testing"

	config_provider_buffered "github.com/CoachApplication/config/provider/buffered"
	config_provider_json "github.com/CoachApplication/config/provider/json"
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
	c := config_provider_json.NewConfig("key", "scope", config_provider_buffered.NewSingle("key", "scope", MessageBytes))

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

func TestJsonConfig_Set(t *testing.T) {
	c := config_provider_json.NewConfig("key", "scope", config_provider_buffered.NewSingle("key", "scope", MessageBytes))

	res := c.Set(MessageStruct)
	<-res.Finished()

	if !res.Success() {
		t.Error("Backend Config reported a failure when Set(): ", res.Errors())
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
