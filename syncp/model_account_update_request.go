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

// checks if the AccountUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUpdateRequest{}

// AccountUpdateRequest struct for AccountUpdateRequest
type AccountUpdateRequest struct {
	JwtSettings          *AccountJWTSettings `json:"jwt_settings,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	UserJwtExpiresInSecs *int64              `json:"user_jwt_expires_in_secs,omitempty"`
}

// NewAccountUpdateRequest instantiates a new AccountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdateRequest() *AccountUpdateRequest {
	this := AccountUpdateRequest{}
	return &this
}

// NewAccountUpdateRequestWithDefaults instantiates a new AccountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateRequestWithDefaults() *AccountUpdateRequest {
	this := AccountUpdateRequest{}
	return &this
}

// GetJwtSettings returns the JwtSettings field value if set, zero value otherwise.
func (o *AccountUpdateRequest) GetJwtSettings() AccountJWTSettings {
	if o == nil || IsNil(o.JwtSettings) {
		var ret AccountJWTSettings
		return ret
	}
	return *o.JwtSettings
}

// GetJwtSettingsOk returns a tuple with the JwtSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateRequest) GetJwtSettingsOk() (*AccountJWTSettings, bool) {
	if o == nil || IsNil(o.JwtSettings) {
		return nil, false
	}
	return o.JwtSettings, true
}

// HasJwtSettings returns a boolean if a field has been set.
func (o *AccountUpdateRequest) HasJwtSettings() bool {
	if o != nil && !IsNil(o.JwtSettings) {
		return true
	}

	return false
}

// SetJwtSettings gets a reference to the given AccountJWTSettings and assigns it to the JwtSettings field.
func (o *AccountUpdateRequest) SetJwtSettings(v AccountJWTSettings) {
	o.JwtSettings = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field value if set, zero value otherwise.
func (o *AccountUpdateRequest) GetUserJwtExpiresInSecs() int64 {
	if o == nil || IsNil(o.UserJwtExpiresInSecs) {
		var ret int64
		return ret
	}
	return *o.UserJwtExpiresInSecs
}

// GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdateRequest) GetUserJwtExpiresInSecsOk() (*int64, bool) {
	if o == nil || IsNil(o.UserJwtExpiresInSecs) {
		return nil, false
	}
	return o.UserJwtExpiresInSecs, true
}

// HasUserJwtExpiresInSecs returns a boolean if a field has been set.
func (o *AccountUpdateRequest) HasUserJwtExpiresInSecs() bool {
	if o != nil && !IsNil(o.UserJwtExpiresInSecs) {
		return true
	}

	return false
}

// SetUserJwtExpiresInSecs gets a reference to the given int64 and assigns it to the UserJwtExpiresInSecs field.
func (o *AccountUpdateRequest) SetUserJwtExpiresInSecs(v int64) {
	o.UserJwtExpiresInSecs = &v
}

func (o AccountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JwtSettings) {
		toSerialize["jwt_settings"] = o.JwtSettings
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserJwtExpiresInSecs) {
		toSerialize["user_jwt_expires_in_secs"] = o.UserJwtExpiresInSecs
	}
	return toSerialize, nil
}

type NullableAccountUpdateRequest struct {
	value *AccountUpdateRequest
	isSet bool
}

func (v NullableAccountUpdateRequest) Get() *AccountUpdateRequest {
	return v.value
}

func (v *NullableAccountUpdateRequest) Set(val *AccountUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdateRequest(val *AccountUpdateRequest) *NullableAccountUpdateRequest {
	return &NullableAccountUpdateRequest{value: val, isSet: true}
}

func (v NullableAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
