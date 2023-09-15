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

// AlertRuleSeverity the model 'AlertRuleSeverity'
type AlertRuleSeverity string

// List of AlertRuleSeverity
const (
	ALERTRULESEVERITY_INFO     AlertRuleSeverity = "Info"
	ALERTRULESEVERITY_WARN     AlertRuleSeverity = "Warn"
	ALERTRULESEVERITY_ERROR    AlertRuleSeverity = "Error"
	ALERTRULESEVERITY_CRITICAL AlertRuleSeverity = "Critical"
)

// All allowed values of AlertRuleSeverity enum
var AllowedAlertRuleSeverityEnumValues = []AlertRuleSeverity{
	"Info",
	"Warn",
	"Error",
	"Critical",
}

func (v *AlertRuleSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertRuleSeverity(value)
	for _, existing := range AllowedAlertRuleSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertRuleSeverity", value)
}

// NewAlertRuleSeverityFromValue returns a pointer to a valid AlertRuleSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertRuleSeverityFromValue(v string) (*AlertRuleSeverity, error) {
	ev := AlertRuleSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertRuleSeverity: valid values are %v", v, AllowedAlertRuleSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertRuleSeverity) IsValid() bool {
	for _, existing := range AllowedAlertRuleSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertRuleSeverity value
func (v AlertRuleSeverity) Ptr() *AlertRuleSeverity {
	return &v
}

type NullableAlertRuleSeverity struct {
	value *AlertRuleSeverity
	isSet bool
}

func (v NullableAlertRuleSeverity) Get() *AlertRuleSeverity {
	return v.value
}

func (v *NullableAlertRuleSeverity) Set(val *AlertRuleSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRuleSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRuleSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRuleSeverity(val *AlertRuleSeverity) *NullableAlertRuleSeverity {
	return &NullableAlertRuleSeverity{value: val, isSet: true}
}

func (v NullableAlertRuleSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRuleSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}