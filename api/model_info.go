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

// checks if the Info type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Info{}

// Info struct for Info
type Info struct {
	Description *string `json:"description,omitempty"`
	InfoUrl     *string `json:"info_url,omitempty"`
}

// NewInfo instantiates a new Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfo() *Info {
	this := Info{}
	return &this
}

// NewInfoWithDefaults instantiates a new Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoWithDefaults() *Info {
	this := Info{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Info) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Info) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Info) SetDescription(v string) {
	o.Description = &v
}

// GetInfoUrl returns the InfoUrl field value if set, zero value otherwise.
func (o *Info) GetInfoUrl() string {
	if o == nil || IsNil(o.InfoUrl) {
		var ret string
		return ret
	}
	return *o.InfoUrl
}

// GetInfoUrlOk returns a tuple with the InfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetInfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InfoUrl) {
		return nil, false
	}
	return o.InfoUrl, true
}

// HasInfoUrl returns a boolean if a field has been set.
func (o *Info) HasInfoUrl() bool {
	if o != nil && !IsNil(o.InfoUrl) {
		return true
	}

	return false
}

// SetInfoUrl gets a reference to the given string and assigns it to the InfoUrl field.
func (o *Info) SetInfoUrl(v string) {
	o.InfoUrl = &v
}

func (o Info) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Info) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InfoUrl) {
		toSerialize["info_url"] = o.InfoUrl
	}
	return toSerialize, nil
}

type NullableInfo struct {
	value *Info
	isSet bool
}

func (v NullableInfo) Get() *Info {
	return v.value
}

func (v *NullableInfo) Set(val *Info) {
	v.value = val
	v.isSet = true
}

func (v NullableInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfo(val *Info) *NullableInfo {
	return &NullableInfo{value: val, isSet: true}
}

func (v NullableInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
