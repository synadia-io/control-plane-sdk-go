/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PromSampleValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromSampleValue{}

// PromSampleValue struct for PromSampleValue
type PromSampleValue struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

// NewPromSampleValue instantiates a new PromSampleValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromSampleValue(timestamp int64, value float64) *PromSampleValue {
	this := PromSampleValue{}
	this.Timestamp = timestamp
	this.Value = value
	return &this
}

// NewPromSampleValueWithDefaults instantiates a new PromSampleValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromSampleValueWithDefaults() *PromSampleValue {
	this := PromSampleValue{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *PromSampleValue) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PromSampleValue) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PromSampleValue) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetValue returns the Value field value
func (o *PromSampleValue) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PromSampleValue) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PromSampleValue) SetValue(v float64) {
	o.Value = v
}

func (o PromSampleValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromSampleValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullablePromSampleValue struct {
	value *PromSampleValue
	isSet bool
}

func (v NullablePromSampleValue) Get() *PromSampleValue {
	return v.value
}

func (v *NullablePromSampleValue) Set(val *PromSampleValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePromSampleValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePromSampleValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromSampleValue(val *PromSampleValue) *NullablePromSampleValue {
	return &NullablePromSampleValue{value: val, isSet: true}
}

func (v NullablePromSampleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromSampleValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
