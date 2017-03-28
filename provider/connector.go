package provider

import (
	"io"
)

type Connector interface {
	Scopes() []string
	Keys() []string
	HasValue(key, scope string) bool
	Get(key, scope string) (io.ReadCloser, error)
	Set(key, scope string, source io.ReadCloser) error
}
