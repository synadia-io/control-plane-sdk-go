/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
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

type NullableAppRoleEffect struct {
	value *AppRoleEffect
	isSet bool
}

func (v NullableAppRoleEffect) Get() *AppRoleEffect {
	return v.value
}

func (v *NullableAppRoleEffect) Set(val *AppRoleEffect) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRoleEffect) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRoleEffect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRoleEffect(val *AppRoleEffect) *NullableAppRoleEffect {
	return &NullableAppRoleEffect{value: val, isSet: true}
}

func (v NullableAppRoleEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRoleEffect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
