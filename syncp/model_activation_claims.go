/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"fmt"
)

// checks if the ActivationClaims type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivationClaims{}

// ActivationClaims ActivationClaims holds the data specific to an activation JWT
type ActivationClaims struct {
	ClaimsData ClaimsData  `json:"ClaimsData"`
	Nats       *Activation `json:"nats,omitempty"`
}

type _ActivationClaims ActivationClaims

// NewActivationClaims instantiates a new ActivationClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationClaims(claimsData ClaimsData) *ActivationClaims {
	this := ActivationClaims{}
	this.ClaimsData = claimsData
	return &this
}

// NewActivationClaimsWithDefaults instantiates a new ActivationClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationClaimsWithDefaults() *ActivationClaims {
	this := ActivationClaims{}
	return &this
}

// GetClaimsData returns the ClaimsData field value
func (o *ActivationClaims) GetClaimsData() ClaimsData {
	if o == nil {
		var ret ClaimsData
		return ret
	}

	return o.ClaimsData
}

// GetClaimsDataOk returns a tuple with the ClaimsData field value
// and a boolean to check if the value has been set.
func (o *ActivationClaims) GetClaimsDataOk() (*ClaimsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimsData, true
}

// SetClaimsData sets field value
func (o *ActivationClaims) SetClaimsData(v ClaimsData) {
	o.ClaimsData = v
}

// GetNats returns the Nats field value if set, zero value otherwise.
func (o *ActivationClaims) GetNats() Activation {
	if o == nil || IsNil(o.Nats) {
		var ret Activation
		return ret
	}
	return *o.Nats
}

// GetNatsOk returns a tuple with the Nats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationClaims) GetNatsOk() (*Activation, bool) {
	if o == nil || IsNil(o.Nats) {
		return nil, false
	}
	return o.Nats, true
}

// HasNats returns a boolean if a field has been set.
func (o *ActivationClaims) HasNats() bool {
	if o != nil && !IsNil(o.Nats) {
		return true
	}

	return false
}

// SetNats gets a reference to the given Activation and assigns it to the Nats field.
func (o *ActivationClaims) SetNats(v Activation) {
	o.Nats = &v
}

func (o ActivationClaims) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivationClaims) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ClaimsData"] = o.ClaimsData
	if !IsNil(o.Nats) {
		toSerialize["nats"] = o.Nats
	}
	return toSerialize, nil
}

func (o *ActivationClaims) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClaimsData",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varActivationClaims := _ActivationClaims{}

	err = json.Unmarshal(bytes, &varActivationClaims)

	if err != nil {
		return err
	}

	*o = ActivationClaims(varActivationClaims)

	return err
}

type NullableActivationClaims struct {
	value *ActivationClaims
	isSet bool
}

func (v NullableActivationClaims) Get() *ActivationClaims {
	return v.value
}

func (v *NullableActivationClaims) Set(val *ActivationClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationClaims(val *ActivationClaims) *NullableActivationClaims {
	return &NullableActivationClaims{value: val, isSet: true}
}

func (v NullableActivationClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
