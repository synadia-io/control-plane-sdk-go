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

// checks if the SigningKeyUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyUpdateRequest{}

// SigningKeyUpdateRequest struct for SigningKeyUpdateRequest
type SigningKeyUpdateRequest struct {
	Disabled bool    `json:"disabled"`
	Seed     *string `json:"seed,omitempty"`
}

// NewSigningKeyUpdateRequest instantiates a new SigningKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKeyUpdateRequest(disabled bool) *SigningKeyUpdateRequest {
	this := SigningKeyUpdateRequest{}
	this.Disabled = disabled
	return &this
}

// NewSigningKeyUpdateRequestWithDefaults instantiates a new SigningKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeyUpdateRequestWithDefaults() *SigningKeyUpdateRequest {
	this := SigningKeyUpdateRequest{}
	return &this
}

// GetDisabled returns the Disabled field value
func (o *SigningKeyUpdateRequest) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *SigningKeyUpdateRequest) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *SigningKeyUpdateRequest) SetDisabled(v bool) {
	o.Disabled = v
}

// GetSeed returns the Seed field value if set, zero value otherwise.
func (o *SigningKeyUpdateRequest) GetSeed() string {
	if o == nil || IsNil(o.Seed) {
		var ret string
		return ret
	}
	return *o.Seed
}

// GetSeedOk returns a tuple with the Seed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeyUpdateRequest) GetSeedOk() (*string, bool) {
	if o == nil || IsNil(o.Seed) {
		return nil, false
	}
	return o.Seed, true
}

// HasSeed returns a boolean if a field has been set.
func (o *SigningKeyUpdateRequest) HasSeed() bool {
	if o != nil && !IsNil(o.Seed) {
		return true
	}

	return false
}

// SetSeed gets a reference to the given string and assigns it to the Seed field.
func (o *SigningKeyUpdateRequest) SetSeed(v string) {
	o.Seed = &v
}

func (o SigningKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningKeyUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["disabled"] = o.Disabled
	if !IsNil(o.Seed) {
		toSerialize["seed"] = o.Seed
	}
	return toSerialize, nil
}

type NullableSigningKeyUpdateRequest struct {
	value *SigningKeyUpdateRequest
	isSet bool
}

func (v NullableSigningKeyUpdateRequest) Get() *SigningKeyUpdateRequest {
	return v.value
}

func (v *NullableSigningKeyUpdateRequest) Set(val *SigningKeyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningKeyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningKeyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningKeyUpdateRequest(val *SigningKeyUpdateRequest) *NullableSigningKeyUpdateRequest {
	return &NullableSigningKeyUpdateRequest{value: val, isSet: true}
}

func (v NullableSigningKeyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningKeyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
