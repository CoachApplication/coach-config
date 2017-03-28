package provider

import (
	api "github.com/CoachApplication/api"
	base "github.com/CoachApplication/base"
	base_errors "github.com/CoachApplication/base/errors"
	base_config "github.com/CoachApplication/config"
)

type GetOperation struct {
	base_config.GetOperationBase
	provider Provider
}

func NewGetOperation(provider Provider) *GetOperation {
	return &GetOperation{
		provider: provider,
	}
}

func (gon *GetOperation) Operation() api.Operation {
	return api.Operation(gon)
}

func (gon *GetOperation) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

func (gon *GetOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	key := ""
	if keyProp, err := props.Get(base_config.PROPERTY_KEY_KEY); err == nil {
		key = keyProp.Get().(string)
	} else {
		res.AddError(err)
		res.AddError(base_errors.RequiredPropertyWasEmptyError{Key: base_config.PROPERTY_KEY_KEY})
	}

	scopedConfig := base_config.NewStandardScopedConfig()
	for _, scope := range gon.provider.Scopes() {
		if config, err := gon.provider.Get(key, scope); err == nil {
			scopedConfig.Set(scope, config)
		} else {
			res.AddError(err)
		}
	}

	scopedConfigProp := base_config.ScopedConfigProperty{}
	scopedConfigProp.Set(scopedConfig)
	res.AddProperty(api.Property(&scopedConfigProp))

	return res.Result()
}

type ListOperation struct {
	base_config.ListOperationBase
	provider Provider
}

func NewListOperation(provider Provider) *ListOperation {
	return &ListOperation{
		provider: provider,
	}
}

func (lo *ListOperation) Operation() api.Operation {
	return api.Operation(lo)
}

func (lo *ListOperation) Validate() api.Result {
	return base.MakeSuccessfulResult()
}

func (lo *ListOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	key := ""
	if keyProp, err := props.Get(base_config.PROPERTY_KEY_KEY); err == nil {
		key = keyProp.Get().(string)
	} else {
		res.AddError(err)
		res.AddError(base_errors.RequiredPropertyWasEmptyError{Key: base_config.PROPERTY_KEY_KEY})
	}

	scopedConfig := base_config.NewStandardScopedConfig()
	for _, scope := range lo.provider.Scopes() {
		if config, err := lo.provider.Get(key, scope); err == nil {
			scopedConfig.Set(scope, config)
		} else {
			res.AddError(err)
		}
	}

	scopedConfigProp := base_config.ScopedConfigProperty{}
	scopedConfigProp.Set(scopedConfig)
	res.AddProperty(api.Property(&scopedConfigProp))

	return res.Result()
}
