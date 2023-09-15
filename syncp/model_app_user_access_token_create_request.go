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

// checks if the AppUserAccessTokenCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAccessTokenCreateRequest{}

// AppUserAccessTokenCreateRequest struct for AppUserAccessTokenCreateRequest
type AppUserAccessTokenCreateRequest struct {
	Expires NullableString `json:"expires"`
	Name    string         `json:"name"`
}

// NewAppUserAccessTokenCreateRequest instantiates a new AppUserAccessTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserAccessTokenCreateRequest(expires NullableString, name string) *AppUserAccessTokenCreateRequest {
	this := AppUserAccessTokenCreateRequest{}
	this.Expires = expires
	this.Name = name
	return &this
}

// NewAppUserAccessTokenCreateRequestWithDefaults instantiates a new AppUserAccessTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserAccessTokenCreateRequestWithDefaults() *AppUserAccessTokenCreateRequest {
	this := AppUserAccessTokenCreateRequest{}
	return &this
}

// GetExpires returns the Expires field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserAccessTokenCreateRequest) GetExpires() string {
	if o == nil || o.Expires.Get() == nil {
		var ret string
		return ret
	}

	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserAccessTokenCreateRequest) GetExpiresOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// SetExpires sets field value
func (o *AppUserAccessTokenCreateRequest) SetExpires(v string) {
	o.Expires.Set(&v)
}

// GetName returns the Name field value
func (o *AppUserAccessTokenCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserAccessTokenCreateRequest) SetName(v string) {
	o.Name = v
}

func (o AppUserAccessTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserAccessTokenCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expires"] = o.Expires.Get()
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAppUserAccessTokenCreateRequest struct {
	value *AppUserAccessTokenCreateRequest
	isSet bool
}

func (v NullableAppUserAccessTokenCreateRequest) Get() *AppUserAccessTokenCreateRequest {
	return v.value
}

func (v *NullableAppUserAccessTokenCreateRequest) Set(val *AppUserAccessTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserAccessTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserAccessTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserAccessTokenCreateRequest(val *AppUserAccessTokenCreateRequest) *NullableAppUserAccessTokenCreateRequest {
	return &NullableAppUserAccessTokenCreateRequest{value: val, isSet: true}
}

func (v NullableAppUserAccessTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserAccessTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}