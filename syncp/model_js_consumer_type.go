/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"fmt"
)

// JSConsumerType the model 'JSConsumerType'
type JSConsumerType string

// List of JSConsumerType
const (
	JSCONSUMERTYPE_PUSH JSConsumerType = "push"
	JSCONSUMERTYPE_PULL JSConsumerType = "pull"
)

// All allowed values of JSConsumerType enum
var AllowedJSConsumerTypeEnumValues = []JSConsumerType{
	"push",
	"pull",
}

func (v *JSConsumerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JSConsumerType(value)
	for _, existing := range AllowedJSConsumerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JSConsumerType", value)
}

// NewJSConsumerTypeFromValue returns a pointer to a valid JSConsumerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJSConsumerTypeFromValue(v string) (*JSConsumerType, error) {
	ev := JSConsumerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JSConsumerType: valid values are %v", v, AllowedJSConsumerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JSConsumerType) IsValid() bool {
	for _, existing := range AllowedJSConsumerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JSConsumerType value
func (v JSConsumerType) Ptr() *JSConsumerType {
	return &v
}

type NullableJSConsumerType struct {
	value *JSConsumerType
	isSet bool
}

func (v NullableJSConsumerType) Get() *JSConsumerType {
	return v.value
}

func (v *NullableJSConsumerType) Set(val *JSConsumerType) {
	v.value = val
	v.isSet = true
}

func (v NullableJSConsumerType) IsSet() bool {
	return v.isSet
}

func (v *NullableJSConsumerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSConsumerType(val *JSConsumerType) *NullableJSConsumerType {
	return &NullableJSConsumerType{value: val, isSet: true}
}

func (v NullableJSConsumerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSConsumerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
