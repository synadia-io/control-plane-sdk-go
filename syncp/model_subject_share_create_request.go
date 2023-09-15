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

// checks if the SubjectShareCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectShareCreateRequest{}

// SubjectShareCreateRequest struct for SubjectShareCreateRequest
type SubjectShareCreateRequest struct {
	TargetAccountNkeyPublic string `json:"target_account_nkey_public"`
}

// NewSubjectShareCreateRequest instantiates a new SubjectShareCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectShareCreateRequest(targetAccountNkeyPublic string) *SubjectShareCreateRequest {
	this := SubjectShareCreateRequest{}
	this.TargetAccountNkeyPublic = targetAccountNkeyPublic
	return &this
}

// NewSubjectShareCreateRequestWithDefaults instantiates a new SubjectShareCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectShareCreateRequestWithDefaults() *SubjectShareCreateRequest {
	this := SubjectShareCreateRequest{}
	return &this
}

// GetTargetAccountNkeyPublic returns the TargetAccountNkeyPublic field value
func (o *SubjectShareCreateRequest) GetTargetAccountNkeyPublic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAccountNkeyPublic
}

// GetTargetAccountNkeyPublicOk returns a tuple with the TargetAccountNkeyPublic field value
// and a boolean to check if the value has been set.
func (o *SubjectShareCreateRequest) GetTargetAccountNkeyPublicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAccountNkeyPublic, true
}

// SetTargetAccountNkeyPublic sets field value
func (o *SubjectShareCreateRequest) SetTargetAccountNkeyPublic(v string) {
	o.TargetAccountNkeyPublic = v
}

func (o SubjectShareCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectShareCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_account_nkey_public"] = o.TargetAccountNkeyPublic
	return toSerialize, nil
}

type NullableSubjectShareCreateRequest struct {
	value *SubjectShareCreateRequest
	isSet bool
}

func (v NullableSubjectShareCreateRequest) Get() *SubjectShareCreateRequest {
	return v.value
}

func (v *NullableSubjectShareCreateRequest) Set(val *SubjectShareCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectShareCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectShareCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectShareCreateRequest(val *SubjectShareCreateRequest) *NullableSubjectShareCreateRequest {
	return &NullableSubjectShareCreateRequest{value: val, isSet: true}
}

func (v NullableSubjectShareCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectShareCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
