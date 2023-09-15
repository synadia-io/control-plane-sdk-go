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

// AccountChartType the model 'AccountChartType'
type AccountChartType string

// List of AccountChartType
const (
	ACCOUNTCHARTTYPE_NUM_LEAF_NODES              AccountChartType = "NumLeafNodes"
	ACCOUNTCHARTTYPE_NUM_CONNECTIONS             AccountChartType = "NumConnections"
	ACCOUNTCHARTTYPE_NUM_SUBSCRIPTIONS           AccountChartType = "NumSubscriptions"
	ACCOUNTCHARTTYPE_BYTES_SENT                  AccountChartType = "BytesSent"
	ACCOUNTCHARTTYPE_BYTES_RECEIVED              AccountChartType = "BytesReceived"
	ACCOUNTCHARTTYPE_NUM_MESSAGES_SENT           AccountChartType = "NumMessagesSent"
	ACCOUNTCHARTTYPE_NUM_MESSAGES_RECEIVED       AccountChartType = "NumMessagesReceived"
	ACCOUNTCHARTTYPE_JET_STREAM_ENABLED          AccountChartType = "JetStreamEnabled"
	ACCOUNTCHARTTYPE_NUM_STREAMS                 AccountChartType = "NumStreams"
	ACCOUNTCHARTTYPE_NUM_STREAM_CONSUMERS        AccountChartType = "NumStreamConsumers"
	ACCOUNTCHARTTYPE_JET_STREAM_MEMORY_USED      AccountChartType = "JetStreamMemoryUsed"
	ACCOUNTCHARTTYPE_JET_STREAM_STORAGE_USED     AccountChartType = "JetStreamStorageUsed"
	ACCOUNTCHARTTYPE_JET_STREAM_MEMORY_RESERVED  AccountChartType = "JetStreamMemoryReserved"
	ACCOUNTCHARTTYPE_JET_STREAM_STORAGE_RESERVED AccountChartType = "JetStreamStorageReserved"
)

// All allowed values of AccountChartType enum
var AllowedAccountChartTypeEnumValues = []AccountChartType{
	"NumLeafNodes",
	"NumConnections",
	"NumSubscriptions",
	"BytesSent",
	"BytesReceived",
	"NumMessagesSent",
	"NumMessagesReceived",
	"JetStreamEnabled",
	"NumStreams",
	"NumStreamConsumers",
	"JetStreamMemoryUsed",
	"JetStreamStorageUsed",
	"JetStreamMemoryReserved",
	"JetStreamStorageReserved",
}

func (v *AccountChartType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountChartType(value)
	for _, existing := range AllowedAccountChartTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountChartType", value)
}

// NewAccountChartTypeFromValue returns a pointer to a valid AccountChartType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountChartTypeFromValue(v string) (*AccountChartType, error) {
	ev := AccountChartType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountChartType: valid values are %v", v, AllowedAccountChartTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountChartType) IsValid() bool {
	for _, existing := range AllowedAccountChartTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountChartType value
func (v AccountChartType) Ptr() *AccountChartType {
	return &v
}

type NullableAccountChartType struct {
	value *AccountChartType
	isSet bool
}

func (v NullableAccountChartType) Get() *AccountChartType {
	return v.value
}

func (v *NullableAccountChartType) Set(val *AccountChartType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountChartType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountChartType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountChartType(val *AccountChartType) *NullableAccountChartType {
	return &NullableAccountChartType{value: val, isSet: true}
}

func (v NullableAccountChartType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountChartType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}