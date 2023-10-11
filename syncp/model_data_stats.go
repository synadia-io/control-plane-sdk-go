/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
)

// checks if the DataStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStats{}

// DataStats DataStats reports how may msg and bytes. Applicable for both sent and received.
type DataStats struct {
	Bytes int64 `json:"bytes"`
	Msgs  int64 `json:"msgs"`
}

// NewDataStats instantiates a new DataStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStats(bytes int64, msgs int64) *DataStats {
	this := DataStats{}
	this.Bytes = bytes
	this.Msgs = msgs
	return &this
}

// NewDataStatsWithDefaults instantiates a new DataStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStatsWithDefaults() *DataStats {
	this := DataStats{}
	return &this
}

// GetBytes returns the Bytes field value
func (o *DataStats) GetBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *DataStats) GetBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *DataStats) SetBytes(v int64) {
	o.Bytes = v
}

// GetMsgs returns the Msgs field value
func (o *DataStats) GetMsgs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Msgs
}

// GetMsgsOk returns a tuple with the Msgs field value
// and a boolean to check if the value has been set.
func (o *DataStats) GetMsgsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msgs, true
}

// SetMsgs sets field value
func (o *DataStats) SetMsgs(v int64) {
	o.Msgs = v
}

func (o DataStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bytes"] = o.Bytes
	toSerialize["msgs"] = o.Msgs
	return toSerialize, nil
}

type NullableDataStats struct {
	value *DataStats
	isSet bool
}

func (v NullableDataStats) Get() *DataStats {
	return v.value
}

func (v *NullableDataStats) Set(val *DataStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStats(val *DataStats) *NullableDataStats {
	return &NullableDataStats{value: val, isSet: true}
}

func (v NullableDataStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
