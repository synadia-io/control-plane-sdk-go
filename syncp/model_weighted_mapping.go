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

// checks if the WeightedMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeightedMapping{}

// WeightedMapping Mapping for publishes
type WeightedMapping struct {
	Cluster *string `json:"cluster,omitempty"`
	// Subject is a string that represents a NATS subject
	Subject string `json:"subject"`
	Weight  *int32 `json:"weight,omitempty"`
}

// NewWeightedMapping instantiates a new WeightedMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeightedMapping(subject string) *WeightedMapping {
	this := WeightedMapping{}
	this.Subject = subject
	return &this
}

// NewWeightedMappingWithDefaults instantiates a new WeightedMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeightedMappingWithDefaults() *WeightedMapping {
	this := WeightedMapping{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *WeightedMapping) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeightedMapping) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *WeightedMapping) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *WeightedMapping) SetCluster(v string) {
	o.Cluster = &v
}

// GetSubject returns the Subject field value
func (o *WeightedMapping) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *WeightedMapping) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *WeightedMapping) SetSubject(v string) {
	o.Subject = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *WeightedMapping) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WeightedMapping) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *WeightedMapping) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *WeightedMapping) SetWeight(v int32) {
	o.Weight = &v
}

func (o WeightedMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeightedMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableWeightedMapping struct {
	value *WeightedMapping
	isSet bool
}

func (v NullableWeightedMapping) Get() *WeightedMapping {
	return v.value
}

func (v *NullableWeightedMapping) Set(val *WeightedMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableWeightedMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableWeightedMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeightedMapping(val *WeightedMapping) *NullableWeightedMapping {
	return &NullableWeightedMapping{value: val, isSet: true}
}

func (v NullableWeightedMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeightedMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
