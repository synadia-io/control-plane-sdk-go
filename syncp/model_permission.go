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

// checks if the Permission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permission{}

// Permission Permission defines allow/deny subjects
type Permission struct {
	// StringList is a wrapper for an array of strings
	Allow []string `json:"allow,omitempty"`
	// StringList is a wrapper for an array of strings
	Deny []string `json:"deny,omitempty"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission() *Permission {
	this := Permission{}
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permission) GetAllow() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permission) GetAllowOk() ([]string, bool) {
	if o == nil || IsNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *Permission) HasAllow() bool {
	if o != nil && IsNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *Permission) SetAllow(v []string) {
	o.Allow = v
}

// GetDeny returns the Deny field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Permission) GetDeny() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Deny
}

// GetDenyOk returns a tuple with the Deny field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Permission) GetDenyOk() ([]string, bool) {
	if o == nil || IsNil(o.Deny) {
		return nil, false
	}
	return o.Deny, true
}

// HasDeny returns a boolean if a field has been set.
func (o *Permission) HasDeny() bool {
	if o != nil && IsNil(o.Deny) {
		return true
	}

	return false
}

// SetDeny gets a reference to the given []string and assigns it to the Deny field.
func (o *Permission) SetDeny(v []string) {
	o.Deny = v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Permission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}
	if o.Deny != nil {
		toSerialize["deny"] = o.Deny
	}
	return toSerialize, nil
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}