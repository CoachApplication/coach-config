package buffered

import (
	"bytes"
	"errors"
	"github.com/CoachApplication/config/provider"
	"github.com/CoachApplication/utils"
	"io"
)

// Multiple config/provider/Connector that holds multiple buffered values in a nested string map
type Multiple struct {
	vals map[string]map[string][]byte // format is [key][scope] = val
}

func NewMultiple(vals map[string]map[string][]byte) *Multiple {
	return &Multiple{
		vals: vals,
	}
}

func (m *Multiple) Connector() provider.Connector {
	return provider.Connector(m)
}

func (m *Multiple) Add(key, scope, value string) error {
	if _, good := m.vals[key]; !good {
		m.vals[key] = map[string][]byte{}
	}
	m.vals[key][scope] = []byte(value)
	return nil
}

func (m *Multiple) Scopes() []string {
	ss := utils.UniqueStringSlice{}

	for _, each := range m.vals {
		for s, _ := range each {
			ss.Add(s)
		}
	}

	return ss.Slice()
}

func (m *Multiple) Keys() []string {
	ks := utils.UniqueStringSlice{}

	for k, _ := range m.vals {
		ks.Add(k)
	}

	return ks.Slice()
}

func (m *Multiple) HasValue(key, scope string) bool {
	if s, good := m.vals[key]; !good {
		return false
	} else if b, good := s[scope]; !good {
		return false
	} else {
		return len(b) > 0
	}
}

func (m *Multiple) Get(key, scope string) (io.ReadCloser, error) {
	if s, good := m.vals[key]; !good {
		return nil, errors.New("No such key exists")
	} else if b, good := s[scope]; !good {
		return nil, errors.New("No such scope exists for that key")
	} else {
		return utils.CloseDecorateReader(io.Reader(bytes.NewBuffer(b)), nil), nil
	}
}

func (m *Multiple) Set(key, scope string, source io.ReadCloser) error {
	if s, good := m.vals[key]; !good {
		return errors.New("No such key exists")
	} else if _, good := s[scope]; !good {
		return errors.New("No such scope exists for that key")
	} else {
		buf := bytes.NewBuffer([]byte{})
		if _, err := buf.ReadFrom(source); err != nil {
			return err
		}
		m.vals[key][scope] = []byte(buf.String())
		return nil
	}
}
