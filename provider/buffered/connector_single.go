package buffered

import (
	"bytes"
	"errors"
	"github.com/CoachApplication/utils"
	"io"
)

// A testing connector that runs from a buffered slice of bytes
type Single struct {
	key   string
	scope string
	val   []byte
}

func NewSingle(key, scope string, val []byte) *Single {
	return &Single{
		key:   key,
		scope: scope,
		val:   val,
	}
}

func (tc *Single) Scopes() []string {
	return []string{tc.scope}
}
func (tc *Single) Keys() []string {
	return []string{tc.key}
}
func (tc *Single) HasValue(key, scope string) bool {
	return key == tc.key && scope == tc.scope
}
func (tc *Single) Get(key, scope string) (io.ReadCloser, error) {
	if key == tc.key && scope == tc.scope {
		return utils.CloseDecorateReader(io.Reader(bytes.NewBuffer(tc.val)), nil), nil
	} else {
		return nil, errors.New("Wrong key scope") // @TODO make a custom error for this
	}
}
func (tc *Single) Set(key, scope string, source io.ReadCloser) error {
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

func (tc *Single) Dump() []byte {
	return tc.val
}
