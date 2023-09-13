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

// DeliverPolicy the model 'DeliverPolicy'
type DeliverPolicy string

// List of DeliverPolicy
const (
	DELIVERPOLICY_ALL               DeliverPolicy = "all"
	DELIVERPOLICY_LAST              DeliverPolicy = "last"
	DELIVERPOLICY_NEW               DeliverPolicy = "new"
	DELIVERPOLICY_BY_START_SEQUENCE DeliverPolicy = "by_start_sequence"
	DELIVERPOLICY_BY_START_TIME     DeliverPolicy = "by_start_time"
	DELIVERPOLICY_LAST_PER_SUBJECT  DeliverPolicy = "last_per_subject"
)

// All allowed values of DeliverPolicy enum
var AllowedDeliverPolicyEnumValues = []DeliverPolicy{
	"all",
	"last",
	"new",
	"by_start_sequence",
	"by_start_time",
	"last_per_subject",
}

func (v *DeliverPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliverPolicy(value)
	for _, existing := range AllowedDeliverPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliverPolicy", value)
}

// NewDeliverPolicyFromValue returns a pointer to a valid DeliverPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliverPolicyFromValue(v string) (*DeliverPolicy, error) {
	ev := DeliverPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliverPolicy: valid values are %v", v, AllowedDeliverPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliverPolicy) IsValid() bool {
	for _, existing := range AllowedDeliverPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeliverPolicy value
func (v DeliverPolicy) Ptr() *DeliverPolicy {
	return &v
}

type NullableDeliverPolicy struct {
	value *DeliverPolicy
	isSet bool
}

func (v NullableDeliverPolicy) Get() *DeliverPolicy {
	return v.value
}

func (v *NullableDeliverPolicy) Set(val *DeliverPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverPolicy(val *DeliverPolicy) *NullableDeliverPolicy {
	return &NullableDeliverPolicy{value: val, isSet: true}
}

func (v NullableDeliverPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
