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

// checks if the AppUserCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserCreateRequest{}

// AppUserCreateRequest struct for AppUserCreateRequest
type AppUserCreateRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

// NewAppUserCreateRequest instantiates a new AppUserCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserCreateRequest(name string, role string, username string) *AppUserCreateRequest {
	this := AppUserCreateRequest{}
	this.Name = name
	this.Role = role
	this.Username = username
	return &this
}

// NewAppUserCreateRequestWithDefaults instantiates a new AppUserCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserCreateRequestWithDefaults() *AppUserCreateRequest {
	this := AppUserCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AppUserCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserCreateRequest) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *AppUserCreateRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *AppUserCreateRequest) SetRole(v string) {
	o.Role = v
}

// GetUsername returns the Username field value
func (o *AppUserCreateRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AppUserCreateRequest) SetUsername(v string) {
	o.Username = v
}

func (o AppUserCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableAppUserCreateRequest struct {
	value *AppUserCreateRequest
	isSet bool
}

func (v NullableAppUserCreateRequest) Get() *AppUserCreateRequest {
	return v.value
}

func (v *NullableAppUserCreateRequest) Set(val *AppUserCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserCreateRequest(val *AppUserCreateRequest) *NullableAppUserCreateRequest {
	return &NullableAppUserCreateRequest{value: val, isSet: true}
}

func (v NullableAppUserCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
