package yaml

import (
	"testing"

	config_provider_buffered "github.com/CoachApplication/config/provider/buffered"
)

// Message JSON bytes
var MessageBytes = []byte(`Name: Alice
Body: Hello
Time: 1294706395881547000`)
var MessageStruct = Message{
	Name: "Alice",
	Body: "Hello",
	Time: 1294706395881547000,
}

// Message struct
type Message struct {
	Name string `yaml:"Name"`
	Body string `yaml:"Body"`
	Time int64  `yaml:"Time"`
}

func TestConfig_Get(t *testing.T) {
	c := NewConfig("key", "scope", config_provider_buffered.NewSingle("key", "scope", MessageBytes))

	var m Message
	res := c.Get(&m)
	<-res.Finished()

	if !res.Success() {
		t.Error("Config Config reported failure in Get() : ", res.Errors())
	} else {

		if m.Name != MessageStruct.Name {
			t.Error("Config provided incorrect data ==> Name : ", m.Name)
		}
		if m.Body != MessageStruct.Body {
			t.Error("Config provided incorrect data ==> Body : ", m.Body)
		}
		if m.Time != MessageStruct.Time {
			t.Error("Config provided incorrect data ==> Time : ", m.Time)
		}

	}

}

func TestConfig_Set(t *testing.T) {
	c := NewConfig("key", "scope", config_provider_buffered.NewSingle("key", "scope", []byte{}))

	res := c.Set(MessageStruct)
	<-res.Finished()

	if !res.Success() {
		t.Error("Config Config reported failure in Set()", res.Errors())
	} else {
		var m Message
		res = c.Get(&m)
		<-res.Finished()

		if !res.Success() {
			t.Error("Config Config reported failure in Get() : ", res.Errors(), m)
		} else {

			if m.Name != MessageStruct.Name {
				t.Error("Config provided incorrect data ==> Name : ", m.Name)
			}
			if m.Body != MessageStruct.Body {
				t.Error("Config provided incorrect data ==> Body : ", m.Body)
			}
			if m.Time != MessageStruct.Time {
				t.Error("Config provided incorrect data ==> Time : ", m.Time)
			}

		}
	}
}
