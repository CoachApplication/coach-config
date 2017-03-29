package config

import (
	"context"

	"errors"
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
	if getOp, err := sw.ops.Get(OPERATION_ID_GET); err != nil {
		return nil, err
	} else if props := getOp.Properties(); len(props.Order()) == 0 {
		return nil, errors.New("Not enough properties in config.get Operation")
	} else if keyProp, err := props.Get(PROPERTY_ID_KEY); err != nil {
		return nil, errors.New("config.get operation had no key property")
	} else if err := keyProp.Set(key); err != nil {
		return nil, err
	} else {
		res := getOp.Exec(props)

		select {
		case <-res.Finished():
			if !res.Success() {
				if errs := res.Errors(); len(errs) == 0 {
					return nil, errors.New("Unknown error occured running config.get operations")
				} else {
					return nil, errs[len(errs)-1]
				}
			} else if scProp, err := res.Properties().Get(PROPERTY_ID_SCOPEDCONFIG); err != nil {
				return nil, err
			} else if sc := scProp.Get(); sc == nil {
				return nil, errors.New("config.get operation produced a nil scopedconfig property value")
			} else if scConv, good := sc.(ScopedConfig); !good {
				return nil, errors.New("config.get operation produced an invalid scopedconfig property value")
			} else {
				return scConv, nil
			}
		case <-sw.ctx.Done():
			return nil, errors.New("config.get operation time out: " + sw.ctx.Err().Error())
		}
	}
}

func (sw *StandardWrapper) List() ([]string, error) {
	if lOp, err := sw.ops.Get(OPERATION_ID_LIST); err != nil {
		return nil, err
	} else {
		res := lOp.Exec(lOp.Properties())

		select {
		case <-res.Finished():
			if !res.Success() {
				if errs := res.Errors(); len(errs) == 0 {
					return nil, errors.New("Unknown error occured running config.list operations")
				} else {
					return nil, errs[len(errs)-1]
				}
			} else if scProp, err := res.Properties().Get(PROPERTY_ID_KEYS); err != nil {
				return nil, err
			} else if ss := scProp.Get(); ss == nil {
				return nil, errors.New("config.list operation produced a nil keys property value")
			} else if ssConv, good := ss.([]string); !good {
				return nil, errors.New("config.list operation produced an invalid keys property value")
			} else {
				return ssConv, nil
			}
		case <-sw.ctx.Done():
			return nil, errors.New("config.list operation time out: " + sw.ctx.Err().Error())
		}
	}
}
