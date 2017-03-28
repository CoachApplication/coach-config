package yaml

import (
	"testing"

	coach_config_provider_bufferred "github.com/CoachApplication/coach-config/provider/bufferred"
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
	conn := coach_config_provider_bufferred.NewBufferedConnector("key", "scope", MessageBytes)
	c := NewConfig("key", "scope", conn)

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
	conn := coach_config_provider_bufferred.NewBufferedConnector("key", "scope", []byte{})
	c := NewConfig("key", "scope", conn)

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
