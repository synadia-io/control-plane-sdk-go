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

// checks if the Authorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Authorization{}

// Authorization struct for Authorization
type Authorization struct {
	// StringList is a wrapper for an array of strings
	AllowedAccounts []string `json:"allowed_accounts,omitempty"`
	// StringList is a wrapper for an array of strings
	AuthUsers []string `json:"auth_users"`
	Xkey      *string  `json:"xkey,omitempty"`
}

// NewAuthorization instantiates a new Authorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorization(authUsers []string) *Authorization {
	this := Authorization{}
	this.AuthUsers = authUsers
	return &this
}

// NewAuthorizationWithDefaults instantiates a new Authorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationWithDefaults() *Authorization {
	this := Authorization{}
	return &this
}

// GetAllowedAccounts returns the AllowedAccounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Authorization) GetAllowedAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedAccounts
}

// GetAllowedAccountsOk returns a tuple with the AllowedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Authorization) GetAllowedAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedAccounts) {
		return nil, false
	}
	return o.AllowedAccounts, true
}

// HasAllowedAccounts returns a boolean if a field has been set.
func (o *Authorization) HasAllowedAccounts() bool {
	if o != nil && IsNil(o.AllowedAccounts) {
		return true
	}

	return false
}

// SetAllowedAccounts gets a reference to the given []string and assigns it to the AllowedAccounts field.
func (o *Authorization) SetAllowedAccounts(v []string) {
	o.AllowedAccounts = v
}

// GetAuthUsers returns the AuthUsers field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Authorization) GetAuthUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthUsers
}

// GetAuthUsersOk returns a tuple with the AuthUsers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Authorization) GetAuthUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthUsers) {
		return nil, false
	}
	return o.AuthUsers, true
}

// SetAuthUsers sets field value
func (o *Authorization) SetAuthUsers(v []string) {
	o.AuthUsers = v
}

// GetXkey returns the Xkey field value if set, zero value otherwise.
func (o *Authorization) GetXkey() string {
	if o == nil || IsNil(o.Xkey) {
		var ret string
		return ret
	}
	return *o.Xkey
}

// GetXkeyOk returns a tuple with the Xkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetXkeyOk() (*string, bool) {
	if o == nil || IsNil(o.Xkey) {
		return nil, false
	}
	return o.Xkey, true
}

// HasXkey returns a boolean if a field has been set.
func (o *Authorization) HasXkey() bool {
	if o != nil && !IsNil(o.Xkey) {
		return true
	}

	return false
}

// SetXkey gets a reference to the given string and assigns it to the Xkey field.
func (o *Authorization) SetXkey(v string) {
	o.Xkey = &v
}

func (o Authorization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedAccounts != nil {
		toSerialize["allowed_accounts"] = o.AllowedAccounts
	}
	if o.AuthUsers != nil {
		toSerialize["auth_users"] = o.AuthUsers
	}
	if !IsNil(o.Xkey) {
		toSerialize["xkey"] = o.Xkey
	}
	return toSerialize, nil
}

type NullableAuthorization struct {
	value *Authorization
	isSet bool
}

func (v NullableAuthorization) Get() *Authorization {
	return v.value
}

func (v *NullableAuthorization) Set(val *Authorization) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorization(val *Authorization) *NullableAuthorization {
	return &NullableAuthorization{value: val, isSet: true}
}

func (v NullableAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
