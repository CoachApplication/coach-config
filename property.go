package configuration

import (
	api "github.com/james-nesbitt/coach-api"
	base "github.com/james-nesbitt/coach-base"
	base_errors "github.com/james-nesbitt/coach-base/errors"
	base_property "github.com/james-nesbitt/coach-base/property"
	"errors"
)

const (
	PROPERTY_KEY_KEY          = "config.key"
	PROPERTY_KEY_KEYS         = "config.keys"
	PROPERTY_KEY_SCOPE        = "config.scope"
	PROPERTY_KEY_SCOPES       = "config.scopes"
	PROPERTY_KEY_SCOPEDCONFIG = "config.scopedconfig"
)

// KeyProperty holds a string config key
type KeyProperty struct {
	base_property.StringProperty
}

// Id Identify the property
func (kp *KeyProperty) Id() string {
	return PROPERTY_KEY_KEY
}

// Ui Provide UI metadata for the Property
func (kp *KeyProperty) Ui() api.Ui {
	return base.NewUi(
		kp.Id(),
		"Configuration key",
		"Key used to uniquely identify the Configuration.",
		"",
	)
}

// Usage Provide Usage information about the element
func (kp *KeyProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}

// KeysProperty holds a set of string config keys
type KeysProperty struct {
	base_property.StringSliceProperty
}

// Id Identify the property
func (kp *KeysProperty) Id() string {
	return PROPERTY_KEY_KEYS
}

// Ui Provide UI metadata for the Property
func (kp *KeysProperty) Ui() api.Ui {
	return base.NewUi(
		kp.Id(),
		"Configuration keys",
		"Keys used to uniquely identify Configurations.",
		"",
	)
}

// Usage Provide Usage information about the element
func (kp *KeysProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}

// ScopeProperty holds a string config scope key
type ScopeProperty struct {
	base_property.StringProperty
}

// Id Identify the property
func (sp *ScopeProperty) Id() string {
	return PROPERTY_KEY_SCOPE
}

// Ui Provide UI metadata for the Property
func (sp *ScopeProperty) Ui() api.Ui {
	return base.NewUi(
		sp.Id(),
		"Configuration scope key",
		"Key used to uniquely identify the Configuration scope.",
		"",
	)
}

// Usage Provide Usage information about the element
func (sp *ScopeProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// ScopedConfigProperty is an api Property that holds a ScopedConfig struct
type ScopedConfigProperty struct {
	value ScopedConfig
}

// Id Identify the property
func (scp *ScopedConfigProperty) Id() string {
	return PROPERTY_KEY_SCOPEDCONFIG
}

// Id Identify the property
func (scp *ScopedConfigProperty) Type() string {
	return "coach.configuration.scopedconfig"
}

// Ui Provide UI metadata for the Property
func (scp *ScopedConfigProperty) Ui() api.Ui {
	return base.NewUi(
		scp.Id(),
		"Configuration ScopedConfig",
		"Configuration data provided using a ScopedConfing interface.",
		"",
	)
}

// Usage Provide Usage information about the element
func (scp *ScopedConfigProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}.Usage()
}

// Validate That the property contains a valid value
func (scp *ScopedConfigProperty) Validate() api.Result {
	res := base.NewResult()

	if scp.value == nil {
		res.AddError(errors.New("No ScopedConfig Value has been set"))
		res.MarkFailed()
	}

	res.MarkFinished()
	return res.Result()
}

//
func (scp *ScopedConfigProperty) Get() interface{} {
	return scp.value
}

//
func (scp *ScopedConfigProperty) Set(value interface{}) error {
	if typedValue, success := value.(ScopedConfig); success {
		scp.value = typedValue
		return nil
	} else {
		return base_errors.PropertyWrongValueTypeError{Id: scp.Id(), ExpectedType: scp.Type()}
	}
}
