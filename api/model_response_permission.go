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

// checks if the ResponsePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePermission{}

// ResponsePermission ResponsePermission can be used to allow responses to any reply subject that is received on a valid subscription.
type ResponsePermission struct {
	Max int32 `json:"max"`
	Ttl int64 `json:"ttl"`
}

// NewResponsePermission instantiates a new ResponsePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePermission(max int32, ttl int64) *ResponsePermission {
	this := ResponsePermission{}
	this.Max = max
	this.Ttl = ttl
	return &this
}

// NewResponsePermissionWithDefaults instantiates a new ResponsePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePermissionWithDefaults() *ResponsePermission {
	this := ResponsePermission{}
	return &this
}

// GetMax returns the Max field value
func (o *ResponsePermission) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *ResponsePermission) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *ResponsePermission) SetMax(v int32) {
	o.Max = v
}

// GetTtl returns the Ttl field value
func (o *ResponsePermission) GetTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *ResponsePermission) GetTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *ResponsePermission) SetTtl(v int64) {
	o.Ttl = v
}

func (o ResponsePermission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max"] = o.Max
	toSerialize["ttl"] = o.Ttl
	return toSerialize, nil
}

type NullableResponsePermission struct {
	value *ResponsePermission
	isSet bool
}

func (v NullableResponsePermission) Get() *ResponsePermission {
	return v.value
}

func (v *NullableResponsePermission) Set(val *ResponsePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePermission(val *ResponsePermission) *NullableResponsePermission {
	return &NullableResponsePermission{value: val, isSet: true}
}

func (v NullableResponsePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
