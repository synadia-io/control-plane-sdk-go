/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"fmt"

	"encoding/json"
)

// AppRoleEffect the model 'AppRoleEffect'
type AppRoleEffect string

// List of AppRoleEffect
const (
	APPROLEEFFECT_ALLOW AppRoleEffect = "ALLOW"
	APPROLEEFFECT_DENY  AppRoleEffect = "DENY"
)

// All allowed values of AppRoleEffect enum
var AllowedAppRoleEffectEnumValues = []AppRoleEffect{
	"ALLOW",
	"DENY",
}

func (v *AppRoleEffect) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppRoleEffect(value)
	for _, existing := range AllowedAppRoleEffectEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppRoleEffect", value)
}

// NewAppRoleEffectFromValue returns a pointer to a valid AppRoleEffect
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppRoleEffectFromValue(v string) (*AppRoleEffect, error) {
	ev := AppRoleEffect(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppRoleEffect: valid values are %v", v, AllowedAppRoleEffectEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppRoleEffect) IsValid() bool {
	for _, existing := range AllowedAppRoleEffectEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppRoleEffect value
func (v AppRoleEffect) Ptr() *AppRoleEffect {
	return &v
}
