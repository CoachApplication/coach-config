package configuration

import (
	coach_api "github.com/james-nesbitt/coach-api"
	coach_base "github.com/james-nesbitt/coach-base"
)

// GetOperation Operation that retrieves a single ScopedConfig for a config Key
type GetOperationBase struct{}

// Id Provide a unique machine name string
func (gob *GetOperationBase) Id() string {
	return "config.get"
}

// Usage Define UI metadata for the Operation
func (gob *GetOperationBase) Ui() coach_api.Ui {
	return coach_base.NewUi(
		gob.Id(),                                                                                     // Id
		"Get Configuration",                                                                          // Label
		"Retrieve scoped Configuration from a configuration backend",                                 // Description
		"Use this Operation to retrieve stored configuration from the system configuration backend.", // Help
	).Ui()
}

// Properties that all implementations will likely need
func (gob *GetOperationBase) Properties() coach_api.Properties {
	props := coach_base.NewProperties()

	// Only the key is needed for the get operations
	props.Add(coach_api.Property(&KeyProperty{}))

	// The Exec method should add the scopevalues

	return props.Properties()
}

// Usage Define how the operations is intended to be used
func (gob *GetOperationBase) Usage() coach_api.Usage {
	return (&coach_base.InternalOperationUsage{}).Usage()
}

// ListOperation Operation that produces a list of Config keys
type ListOperationBase struct{}

// Id Provide a unique machine name string
func (lob *ListOperationBase) Id() string {
	return "config.get"
}

// Usage Define UI metadata for the Operation
func (lob *ListOperationBase) Ui() coach_api.Ui {
	return coach_base.NewUi(
		lob.Id(),                                                                                 // Id
		"List Configuration",                                                                     // Label
		"List Configurations avaialble from a configuration backend",                             // Description
		"Use this Operation to list stored configuration from the system configuration backend.", // Help
	).Ui()
}

// Properties that all implementations will likely need
func (lob *ListOperationBase) Properties() coach_api.Properties {
	props := coach_base.NewProperties()

	// Optionall allow a scope limitation for the config keys
	props.Add(coach_api.Property(&ScopeProperty{}))

	// The Exec method should add the keys property

	return props.Properties()
}

// Usage Define how the operations is intended to be used
func (lob *ListOperationBase) Usage() coach_api.Usage {
	return (&coach_base.InternalOperationUsage{}).Usage()
}
