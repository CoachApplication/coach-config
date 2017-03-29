package config

import (
	"context"

	api "github.com/CoachApplication/api"
)

type StandardWrapper struct {
	ctx context.Context
	ops api.Operations
}

func NewStandardWrapper(ops api.Operations, ctx context.Context) *StandardWrapper {
	if ctx == nil {
		ctx = context.Background()
	}
	return &StandardWrapper{
		ops: ops,
		ctx: ctx,
	}
}

func (sw *StandardWrapper) Wrapper() Wrapper {
	return Wrapper(sw)
}

func (sw *StandardWrapper) Get(key string) (ScopedConfig, error) {
	return nil, nil
}

func (sw *StandardWrapper) List() ([]string, error) {
	return []string{}, nil
}
