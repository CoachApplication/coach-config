package buffered

import (
	"bytes"
	"errors"
	"io"
)

// A testing connector that runs from a buffered slice of bytes
type Connector struct {
	key   string
	scope string
	val   []byte
}

func NewConnector(key, scope string, val []byte) *Connector {
	return &Connector{
		key:   key,
		scope: scope,
		val:   val,
	}
}

func (tc *Connector) Scopes() []string {
	return []string{tc.scope}
}
func (tc *Connector) Keys() []string {
	return []string{tc.key}
}
func (tc *Connector) Get(key, scope string) (io.ReadCloser, error) {
	if key == tc.key && scope == tc.scope {
		return io.ReadCloser(&BufferCloser{Buffer: *bytes.NewBuffer(tc.val)}), nil
	} else {
		return nil, errors.New("Wrong key scope") // @TODO make a custom error for this
	}
}
func (tc *Connector) Set(key, scope string, source io.ReadCloser) error {
	buf := bytes.NewBuffer([]byte{})
	if _, err := buf.ReadFrom(source); err != nil {
		return err
	}
	tc.val = []byte(buf.String())
	if err := source.Close(); err != nil {
		return err
	}
	return nil
}

type BufferCloser struct {
	bytes.Buffer
}

func (bc *BufferCloser) Close() error {
	return nil
}
