/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
)

// checks if the SigningKeyGroupCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyGroupCreateRequest{}

// SigningKeyGroupCreateRequest struct for SigningKeyGroupCreateRequest
type SigningKeyGroupCreateRequest struct {
	Name  string                `json:"name"`
	Scope *UserPermissionLimits `json:"scope,omitempty"`
	Seed  *string               `json:"seed,omitempty"`
}

// NewSigningKeyGroupCreateRequest instantiates a new SigningKeyGroupCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKeyGroupCreateRequest(name string) *SigningKeyGroupCreateRequest {
	this := SigningKeyGroupCreateRequest{}
	this.Name = name
	return &this
}

// NewSigningKeyGroupCreateRequestWithDefaults instantiates a new SigningKeyGroupCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeyGroupCreateRequestWithDefaults() *SigningKeyGroupCreateRequest {
	this := SigningKeyGroupCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SigningKeyGroupCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SigningKeyGroupCreateRequest) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SigningKeyGroupCreateRequest) GetScope() UserPermissionLimits {
	if o == nil || IsNil(o.Scope) {
		var ret UserPermissionLimits
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupCreateRequest) GetScopeOk() (*UserPermissionLimits, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SigningKeyGroupCreateRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given UserPermissionLimits and assigns it to the Scope field.
func (o *SigningKeyGroupCreateRequest) SetScope(v UserPermissionLimits) {
	o.Scope = &v
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (o *SigningKeyGroupCreateRequest) GetSeed() string {
	if o == nil || IsNil(o.Seed) {
		var ret string
		return ret
	}
	return *o.Seed
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupCreateRequest) GetSeedOk() (*string, bool) {
	if o == nil || IsNil(o.Seed) {
		return nil, false
	}
	return o.Seed, true
}

// HasSeed returns a boolean if a field has been set.
func (o *SigningKeyGroupCreateRequest) HasSeed() bool {
	if o != nil && !IsNil(o.Seed) {
		return true
	}

	return false
}

// SetSeed gets a reference to the given string and assigns it to the Seed field.
func (o *SigningKeyGroupCreateRequest) SetSeed(v string) {
	o.Seed = &v
}

func (o SigningKeyGroupCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningKeyGroupCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Seed) {
		toSerialize["seed"] = o.Seed
	}
	return toSerialize, nil
}

type NullableSigningKeyGroupCreateRequest struct {
	value *SigningKeyGroupCreateRequest
	isSet bool
}

func (v NullableSigningKeyGroupCreateRequest) Get() *SigningKeyGroupCreateRequest {
	return v.value
}

func (v *NullableSigningKeyGroupCreateRequest) Set(val *SigningKeyGroupCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningKeyGroupCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningKeyGroupCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningKeyGroupCreateRequest(val *SigningKeyGroupCreateRequest) *NullableSigningKeyGroupCreateRequest {
	return &NullableSigningKeyGroupCreateRequest{value: val, isSet: true}
}

func (v NullableSigningKeyGroupCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningKeyGroupCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
