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

// RetentionPolicy the model 'RetentionPolicy'
type RetentionPolicy string

// List of RetentionPolicy
const (
	RETENTIONPOLICY_LIMITS    RetentionPolicy = "limits"
	RETENTIONPOLICY_INTEREST  RetentionPolicy = "interest"
	RETENTIONPOLICY_WORKQUEUE RetentionPolicy = "workqueue"
)

// All allowed values of RetentionPolicy enum
var AllowedRetentionPolicyEnumValues = []RetentionPolicy{
	"limits",
	"interest",
	"workqueue",
}

func (v *RetentionPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RetentionPolicy(value)
	for _, existing := range AllowedRetentionPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RetentionPolicy", value)
}

// NewRetentionPolicyFromValue returns a pointer to a valid RetentionPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRetentionPolicyFromValue(v string) (*RetentionPolicy, error) {
	ev := RetentionPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RetentionPolicy: valid values are %v", v, AllowedRetentionPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RetentionPolicy) IsValid() bool {
	for _, existing := range AllowedRetentionPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionPolicy value
func (v RetentionPolicy) Ptr() *RetentionPolicy {
	return &v
}

type NullableRetentionPolicy struct {
	value *RetentionPolicy
	isSet bool
}

func (v NullableRetentionPolicy) Get() *RetentionPolicy {
	return v.value
}

func (v *NullableRetentionPolicy) Set(val *RetentionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionPolicy(val *RetentionPolicy) *NullableRetentionPolicy {
	return &NullableRetentionPolicy{value: val, isSet: true}
}

func (v NullableRetentionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
