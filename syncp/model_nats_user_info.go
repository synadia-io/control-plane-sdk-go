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

// checks if the NatsUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserInfo{}

// NatsUserInfo struct for NatsUserInfo
type NatsUserInfo struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	UserPublicKey string `json:"user_public_key"`
}

// NewNatsUserInfo instantiates a new NatsUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatsUserInfo(id string, name string, userPublicKey string) *NatsUserInfo {
	this := NatsUserInfo{}
	this.Id = id
	this.Name = name
	this.UserPublicKey = userPublicKey
	return &this
}

// NewNatsUserInfoWithDefaults instantiates a new NatsUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatsUserInfoWithDefaults() *NatsUserInfo {
	this := NatsUserInfo{}
	return &this
}

// GetId returns the Id field value
func (o *NatsUserInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NatsUserInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NatsUserInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *NatsUserInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NatsUserInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NatsUserInfo) SetName(v string) {
	o.Name = v
}

// GetUserPublicKey returns the UserPublicKey field value
func (o *NatsUserInfo) GetUserPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserPublicKey
}

// GetUserPublicKeyOk returns a tuple with the UserPublicKey field value
// and a boolean to check if the value has been set.
func (o *NatsUserInfo) GetUserPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserPublicKey, true
}

// SetUserPublicKey sets field value
func (o *NatsUserInfo) SetUserPublicKey(v string) {
	o.UserPublicKey = v
}

func (o NatsUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NatsUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["user_public_key"] = o.UserPublicKey
	return toSerialize, nil
}

type NullableNatsUserInfo struct {
	value *NatsUserInfo
	isSet bool
}

func (v NullableNatsUserInfo) Get() *NatsUserInfo {
	return v.value
}

func (v *NullableNatsUserInfo) Set(val *NatsUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNatsUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNatsUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatsUserInfo(val *NatsUserInfo) *NullableNatsUserInfo {
	return &NullableNatsUserInfo{value: val, isSet: true}
}

func (v NullableNatsUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatsUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
