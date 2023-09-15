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

// checks if the PermissionsResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsResp{}

// PermissionsResp struct for PermissionsResp
type PermissionsResp struct {
	Max int32 `json:"max"`
	Ttl int64 `json:"ttl"`
}

// NewPermissionsResp instantiates a new PermissionsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsResp(max int32, ttl int64) *PermissionsResp {
	this := PermissionsResp{}
	this.Max = max
	this.Ttl = ttl
	return &this
}

// NewPermissionsRespWithDefaults instantiates a new PermissionsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsRespWithDefaults() *PermissionsResp {
	this := PermissionsResp{}
	return &this
}

// GetMax returns the Max field value
func (o *PermissionsResp) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *PermissionsResp) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *PermissionsResp) SetMax(v int32) {
	o.Max = v
}

// GetTtl returns the Ttl field value
func (o *PermissionsResp) GetTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *PermissionsResp) GetTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *PermissionsResp) SetTtl(v int64) {
	o.Ttl = v
}

func (o PermissionsResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["max"] = o.Max
	toSerialize["ttl"] = o.Ttl
	return toSerialize, nil
}

type NullablePermissionsResp struct {
	value *PermissionsResp
	isSet bool
}

func (v NullablePermissionsResp) Get() *PermissionsResp {
	return v.value
}

func (v *NullablePermissionsResp) Set(val *PermissionsResp) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsResp) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsResp(val *PermissionsResp) *NullablePermissionsResp {
	return &NullablePermissionsResp{value: val, isSet: true}
}

func (v NullablePermissionsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}