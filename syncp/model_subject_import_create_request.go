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

// checks if the SubjectImportCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectImportCreateRequest{}

// SubjectImportCreateRequest struct for SubjectImportCreateRequest
type SubjectImportCreateRequest struct {
	JwtSettings Import `json:"jwt_settings"`
}

// NewSubjectImportCreateRequest instantiates a new SubjectImportCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectImportCreateRequest(jwtSettings Import) *SubjectImportCreateRequest {
	this := SubjectImportCreateRequest{}
	this.JwtSettings = jwtSettings
	return &this
}

// NewSubjectImportCreateRequestWithDefaults instantiates a new SubjectImportCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectImportCreateRequestWithDefaults() *SubjectImportCreateRequest {
	this := SubjectImportCreateRequest{}
	return &this
}

// GetJwtSettings returns the JwtSettings field value
func (o *SubjectImportCreateRequest) GetJwtSettings() Import {
	if o == nil {
		var ret Import
		return ret
	}

	return o.JwtSettings
}

// GetJwtSettingsOk returns a tuple with the JwtSettings field value
// and a boolean to check if the value has been set.
func (o *SubjectImportCreateRequest) GetJwtSettingsOk() (*Import, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwtSettings, true
}

// SetJwtSettings sets field value
func (o *SubjectImportCreateRequest) SetJwtSettings(v Import) {
	o.JwtSettings = v
}

func (o SubjectImportCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectImportCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jwt_settings"] = o.JwtSettings
	return toSerialize, nil
}

type NullableSubjectImportCreateRequest struct {
	value *SubjectImportCreateRequest
	isSet bool
}

func (v NullableSubjectImportCreateRequest) Get() *SubjectImportCreateRequest {
	return v.value
}

func (v *NullableSubjectImportCreateRequest) Set(val *SubjectImportCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectImportCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectImportCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectImportCreateRequest(val *SubjectImportCreateRequest) *NullableSubjectImportCreateRequest {
	return &NullableSubjectImportCreateRequest{value: val, isSet: true}
}

func (v NullableSubjectImportCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectImportCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}