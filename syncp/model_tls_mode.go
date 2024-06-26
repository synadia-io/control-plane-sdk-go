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

// TLSMode the model 'TLSMode'
type TLSMode string

// List of TLSMode
const (
	TLSMODE_AUTO     TLSMode = "Auto"
	TLSMODE_REQUIRE  TLSMode = "Require"
	TLSMODE_IMPLICIT TLSMode = "Implicit"
)

// All allowed values of TLSMode enum
var AllowedTLSModeEnumValues = []TLSMode{
	"Auto",
	"Require",
	"Implicit",
}

func (v *TLSMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TLSMode(value)
	for _, existing := range AllowedTLSModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TLSMode", value)
}

// NewTLSModeFromValue returns a pointer to a valid TLSMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTLSModeFromValue(v string) (*TLSMode, error) {
	ev := TLSMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TLSMode: valid values are %v", v, AllowedTLSModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TLSMode) IsValid() bool {
	for _, existing := range AllowedTLSModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TLSMode value
func (v TLSMode) Ptr() *TLSMode {
	return &v
}
