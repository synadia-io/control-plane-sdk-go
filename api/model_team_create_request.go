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

// checks if the TeamCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamCreateRequest{}

// TeamCreateRequest struct for TeamCreateRequest
type TeamCreateRequest struct {
	Name string `json:"name"`
}

// NewTeamCreateRequest instantiates a new TeamCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamCreateRequest(name string) *TeamCreateRequest {
	this := TeamCreateRequest{}
	this.Name = name
	return &this
}

// NewTeamCreateRequestWithDefaults instantiates a new TeamCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamCreateRequestWithDefaults() *TeamCreateRequest {
	this := TeamCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TeamCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamCreateRequest) SetName(v string) {
	o.Name = v
}

func (o TeamCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTeamCreateRequest struct {
	value *TeamCreateRequest
	isSet bool
}

func (v NullableTeamCreateRequest) Get() *TeamCreateRequest {
	return v.value
}

func (v *NullableTeamCreateRequest) Set(val *TeamCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamCreateRequest(val *TeamCreateRequest) *NullableTeamCreateRequest {
	return &NullableTeamCreateRequest{value: val, isSet: true}
}

func (v NullableTeamCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
