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

// checks if the AppUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserInfo{}

// AppUserInfo struct for AppUserInfo
type AppUserInfo struct {
	Id         string         `json:"id"`
	Identifier NullableString `json:"identifier"`
	Name       string         `json:"name"`
}

// NewAppUserInfo instantiates a new AppUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserInfo(id string, identifier NullableString, name string) *AppUserInfo {
	this := AppUserInfo{}
	this.Id = id
	this.Identifier = identifier
	this.Name = name
	return &this
}

// NewAppUserInfoWithDefaults instantiates a new AppUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserInfoWithDefaults() *AppUserInfo {
	this := AppUserInfo{}
	return &this
}

// GetId returns the Id field value
func (o *AppUserInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUserInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUserInfo) SetId(v string) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserInfo) GetIdentifier() string {
	if o == nil || o.Identifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserInfo) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// SetIdentifier sets field value
func (o *AppUserInfo) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}

// GetName returns the Name field value
func (o *AppUserInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserInfo) SetName(v string) {
	o.Name = v
}

func (o AppUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["identifier"] = o.Identifier.Get()
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAppUserInfo struct {
	value *AppUserInfo
	isSet bool
}

func (v NullableAppUserInfo) Get() *AppUserInfo {
	return v.value
}

func (v *NullableAppUserInfo) Set(val *AppUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserInfo(val *AppUserInfo) *NullableAppUserInfo {
	return &NullableAppUserInfo{value: val, isSet: true}
}

func (v NullableAppUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}