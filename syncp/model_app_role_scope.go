/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"fmt"

	"encoding/json"
)

// AppRoleScope the model 'AppRoleScope'
type AppRoleScope string

// List of AppRoleScope
const (
	APPROLESCOPE_APP       AppRoleScope = "App"
	APPROLESCOPE_TEAM      AppRoleScope = "Team"
	APPROLESCOPE_SYSTEM    AppRoleScope = "System"
	APPROLESCOPE_ACCOUNT   AppRoleScope = "Account"
	APPROLESCOPE_NATS_USER AppRoleScope = "NatsUser"
)

// All allowed values of AppRoleScope enum
var AllowedAppRoleScopeEnumValues = []AppRoleScope{
	"App",
	"Team",
	"System",
	"Account",
	"NatsUser",
}

func (v *AppRoleScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppRoleScope(value)
	for _, existing := range AllowedAppRoleScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppRoleScope", value)
}

// NewAppRoleScopeFromValue returns a pointer to a valid AppRoleScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppRoleScopeFromValue(v string) (*AppRoleScope, error) {
	ev := AppRoleScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppRoleScope: valid values are %v", v, AllowedAppRoleScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppRoleScope) IsValid() bool {
	for _, existing := range AllowedAppRoleScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppRoleScope value
func (v AppRoleScope) Ptr() *AppRoleScope {
	return &v
}
