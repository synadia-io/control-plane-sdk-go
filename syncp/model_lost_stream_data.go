/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
)

// checks if the LostStreamData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LostStreamData{}

// LostStreamData LostStreamData indicates msgs that have been lost during file checks and recover due to corruption
type LostStreamData struct {
	// How many bytes were lost
	Bytes int32 `json:"bytes"`
	// Message IDs of lost messages
	Msgs []int32 `json:"msgs"`
}

// NewLostStreamData instantiates a new LostStreamData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLostStreamData(bytes int32, msgs []int32) *LostStreamData {
	this := LostStreamData{}
	this.Bytes = bytes
	this.Msgs = msgs
	return &this
}

// NewLostStreamDataWithDefaults instantiates a new LostStreamData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLostStreamDataWithDefaults() *LostStreamData {
	this := LostStreamData{}
	return &this
}

// GetBytes returns the Bytes field value
func (o *LostStreamData) GetBytes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *LostStreamData) GetBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *LostStreamData) SetBytes(v int32) {
	o.Bytes = v
}

// GetMsgs returns the Msgs field value
func (o *LostStreamData) GetMsgs() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Msgs
}

// GetMsgsOk returns a tuple with the Msgs field value
// and a boolean to check if the value has been set.
func (o *LostStreamData) GetMsgsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Msgs, true
}

// SetMsgs sets field value
func (o *LostStreamData) SetMsgs(v []int32) {
	o.Msgs = v
}

func (o LostStreamData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LostStreamData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bytes"] = o.Bytes
	toSerialize["msgs"] = o.Msgs
	return toSerialize, nil
}

type NullableLostStreamData struct {
	value *LostStreamData
	isSet bool
}

func (v NullableLostStreamData) Get() *LostStreamData {
	return v.value
}

func (v *NullableLostStreamData) Set(val *LostStreamData) {
	v.value = val
	v.isSet = true
}

func (v NullableLostStreamData) IsSet() bool {
	return v.isSet
}

func (v *NullableLostStreamData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLostStreamData(val *LostStreamData) *NullableLostStreamData {
	return &NullableLostStreamData{value: val, isSet: true}
}

func (v NullableLostStreamData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLostStreamData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
