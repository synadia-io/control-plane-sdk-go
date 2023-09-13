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

// checks if the ExternalStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalStream{}

// ExternalStream ExternalStream allows you to qualify access to a stream source in another account.
type ExternalStream struct {
	Api     string `json:"api"`
	Deliver string `json:"deliver"`
}

// NewExternalStream instantiates a new ExternalStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalStream(api string, deliver string) *ExternalStream {
	this := ExternalStream{}
	this.Api = api
	this.Deliver = deliver
	return &this
}

// NewExternalStreamWithDefaults instantiates a new ExternalStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalStreamWithDefaults() *ExternalStream {
	this := ExternalStream{}
	return &this
}

// GetApi returns the Api field value
func (o *ExternalStream) GetApi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Api
}

// GetApiOk returns a tuple with the Api field value
// and a boolean to check if the value has been set.
func (o *ExternalStream) GetApiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Api, true
}

// SetApi sets field value
func (o *ExternalStream) SetApi(v string) {
	o.Api = v
}

// GetDeliver returns the Deliver field value
func (o *ExternalStream) GetDeliver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Deliver
}

// GetDeliverOk returns a tuple with the Deliver field value
// and a boolean to check if the value has been set.
func (o *ExternalStream) GetDeliverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deliver, true
}

// SetDeliver sets field value
func (o *ExternalStream) SetDeliver(v string) {
	o.Deliver = v
}

func (o ExternalStream) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api"] = o.Api
	toSerialize["deliver"] = o.Deliver
	return toSerialize, nil
}

type NullableExternalStream struct {
	value *ExternalStream
	isSet bool
}

func (v NullableExternalStream) Get() *ExternalStream {
	return v.value
}

func (v *NullableExternalStream) Set(val *ExternalStream) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalStream) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalStream(val *ExternalStream) *NullableExternalStream {
	return &NullableExternalStream{value: val, isSet: true}
}

func (v NullableExternalStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
