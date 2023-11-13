/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"fmt"
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

type NullableAppRoleScope struct {
	value *AppRoleScope
	isSet bool
}

func (v NullableAppRoleScope) Get() *AppRoleScope {
	return v.value
}

func (v *NullableAppRoleScope) Set(val *AppRoleScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRoleScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRoleScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRoleScope(val *AppRoleScope) *NullableAppRoleScope {
	return &NullableAppRoleScope{value: val, isSet: true}
}

func (v NullableAppRoleScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRoleScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
