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

// checks if the AppUserUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserUpdateRequest{}

// AppUserUpdateRequest struct for AppUserUpdateRequest
type AppUserUpdateRequest struct {
	Name          *string `json:"name,omitempty"`
	ResetPassword *bool   `json:"reset_password,omitempty"`
	Role          *string `json:"role,omitempty"`
}

// NewAppUserUpdateRequest instantiates a new AppUserUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserUpdateRequest() *AppUserUpdateRequest {
	this := AppUserUpdateRequest{}
	return &this
}

// NewAppUserUpdateRequestWithDefaults instantiates a new AppUserUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserUpdateRequestWithDefaults() *AppUserUpdateRequest {
	this := AppUserUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppUserUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppUserUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppUserUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetResetPassword returns the ResetPassword field value if set, zero value otherwise.
func (o *AppUserUpdateRequest) GetResetPassword() bool {
	if o == nil || IsNil(o.ResetPassword) {
		var ret bool
		return ret
	}
	return *o.ResetPassword
}

// GetResetPasswordOk returns a tuple with the ResetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserUpdateRequest) GetResetPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.ResetPassword) {
		return nil, false
	}
	return o.ResetPassword, true
}

// HasResetPassword returns a boolean if a field has been set.
func (o *AppUserUpdateRequest) HasResetPassword() bool {
	if o != nil && !IsNil(o.ResetPassword) {
		return true
	}

	return false
}

// SetResetPassword gets a reference to the given bool and assigns it to the ResetPassword field.
func (o *AppUserUpdateRequest) SetResetPassword(v bool) {
	o.ResetPassword = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AppUserUpdateRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserUpdateRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AppUserUpdateRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AppUserUpdateRequest) SetRole(v string) {
	o.Role = &v
}

func (o AppUserUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ResetPassword) {
		toSerialize["reset_password"] = o.ResetPassword
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableAppUserUpdateRequest struct {
	value *AppUserUpdateRequest
	isSet bool
}

func (v NullableAppUserUpdateRequest) Get() *AppUserUpdateRequest {
	return v.value
}

func (v *NullableAppUserUpdateRequest) Set(val *AppUserUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserUpdateRequest(val *AppUserUpdateRequest) *NullableAppUserUpdateRequest {
	return &NullableAppUserUpdateRequest{value: val, isSet: true}
}

func (v NullableAppUserUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
