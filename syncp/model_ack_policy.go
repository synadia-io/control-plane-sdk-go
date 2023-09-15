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

// AckPolicy enums have been changed to match UnmarshalJSON in https://github.com/nats-io/jsm.go/blob/main/api/consumers.go
type AckPolicy string

// List of AckPolicy
const (
	ACKPOLICY_NONE     AckPolicy = "none"
	ACKPOLICY_ALL      AckPolicy = "all"
	ACKPOLICY_EXPLICIT AckPolicy = "explicit"
)

// All allowed values of AckPolicy enum
var AllowedAckPolicyEnumValues = []AckPolicy{
	"none",
	"all",
	"explicit",
}

func (v *AckPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AckPolicy(value)
	for _, existing := range AllowedAckPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AckPolicy", value)
}

// NewAckPolicyFromValue returns a pointer to a valid AckPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAckPolicyFromValue(v string) (*AckPolicy, error) {
	ev := AckPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AckPolicy: valid values are %v", v, AllowedAckPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AckPolicy) IsValid() bool {
	for _, existing := range AllowedAckPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AckPolicy value
func (v AckPolicy) Ptr() *AckPolicy {
	return &v
}

type NullableAckPolicy struct {
	value *AckPolicy
	isSet bool
}

func (v NullableAckPolicy) Get() *AckPolicy {
	return v.value
}

func (v *NullableAckPolicy) Set(val *AckPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAckPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAckPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAckPolicy(val *AckPolicy) *NullableAckPolicy {
	return &NullableAckPolicy{value: val, isSet: true}
}

func (v NullableAckPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAckPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
