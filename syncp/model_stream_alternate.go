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

// checks if the StreamAlternate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamAlternate{}

// StreamAlternate struct for StreamAlternate
type StreamAlternate struct {
	Cluster string  `json:"cluster"`
	Domain  *string `json:"domain,omitempty"`
	Name    string  `json:"name"`
}

// NewStreamAlternate instantiates a new StreamAlternate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamAlternate(cluster string, name string) *StreamAlternate {
	this := StreamAlternate{}
	this.Cluster = cluster
	this.Name = name
	return &this
}

// NewStreamAlternateWithDefaults instantiates a new StreamAlternate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamAlternateWithDefaults() *StreamAlternate {
	this := StreamAlternate{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *StreamAlternate) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *StreamAlternate) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *StreamAlternate) SetCluster(v string) {
	o.Cluster = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *StreamAlternate) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamAlternate) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *StreamAlternate) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *StreamAlternate) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value
func (o *StreamAlternate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StreamAlternate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StreamAlternate) SetName(v string) {
	o.Name = v
}

func (o StreamAlternate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamAlternate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableStreamAlternate struct {
	value *StreamAlternate
	isSet bool
}

func (v NullableStreamAlternate) Get() *StreamAlternate {
	return v.value
}

func (v *NullableStreamAlternate) Set(val *StreamAlternate) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamAlternate) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamAlternate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamAlternate(val *StreamAlternate) *NullableStreamAlternate {
	return &NullableStreamAlternate{value: val, isSet: true}
}

func (v NullableStreamAlternate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamAlternate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
