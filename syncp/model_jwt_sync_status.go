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

// JwtSyncStatus the model 'JwtSyncStatus'
type JwtSyncStatus string

// List of JwtSyncStatus
const (
	JWTSYNCSTATUS_IN_PROGRESS JwtSyncStatus = "InProgress"
	JWTSYNCSTATUS_COMPLETE    JwtSyncStatus = "Complete"
	JWTSYNCSTATUS_FAILED      JwtSyncStatus = "Failed"
)

// All allowed values of JwtSyncStatus enum
var AllowedJwtSyncStatusEnumValues = []JwtSyncStatus{
	"InProgress",
	"Complete",
	"Failed",
}

func (v *JwtSyncStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JwtSyncStatus(value)
	for _, existing := range AllowedJwtSyncStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JwtSyncStatus", value)
}

// NewJwtSyncStatusFromValue returns a pointer to a valid JwtSyncStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJwtSyncStatusFromValue(v string) (*JwtSyncStatus, error) {
	ev := JwtSyncStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JwtSyncStatus: valid values are %v", v, AllowedJwtSyncStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JwtSyncStatus) IsValid() bool {
	for _, existing := range AllowedJwtSyncStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JwtSyncStatus value
func (v JwtSyncStatus) Ptr() *JwtSyncStatus {
	return &v
}
