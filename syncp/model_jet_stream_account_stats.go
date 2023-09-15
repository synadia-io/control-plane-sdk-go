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

// checks if the JetStreamAccountStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JetStreamAccountStats{}

// JetStreamAccountStats JetStreamAccountStats returns current statistics about the account's JetStream usage.
type JetStreamAccountStats struct {
	JetStreamTier JetStreamTier             `json:"JetStreamTier"`
	Api           JetStreamAPIStats         `json:"api"`
	Domain        *string                   `json:"domain,omitempty"`
	Tiers         *map[string]JetStreamTier `json:"tiers,omitempty"`
}

// NewJetStreamAccountStats instantiates a new JetStreamAccountStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJetStreamAccountStats(jetStreamTier JetStreamTier, api JetStreamAPIStats) *JetStreamAccountStats {
	this := JetStreamAccountStats{}
	this.JetStreamTier = jetStreamTier
	this.Api = api
	return &this
}

// NewJetStreamAccountStatsWithDefaults instantiates a new JetStreamAccountStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJetStreamAccountStatsWithDefaults() *JetStreamAccountStats {
	this := JetStreamAccountStats{}
	return &this
}

// GetJetStreamTier returns the JetStreamTier field value
func (o *JetStreamAccountStats) GetJetStreamTier() JetStreamTier {
	if o == nil {
		var ret JetStreamTier
		return ret
	}

	return o.JetStreamTier
}

// GetJetStreamTierOk returns a tuple with the JetStreamTier field value
// and a boolean to check if the value has been set.
func (o *JetStreamAccountStats) GetJetStreamTierOk() (*JetStreamTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JetStreamTier, true
}

// SetJetStreamTier sets field value
func (o *JetStreamAccountStats) SetJetStreamTier(v JetStreamTier) {
	o.JetStreamTier = v
}

// GetApi returns the Api field value
func (o *JetStreamAccountStats) GetApi() JetStreamAPIStats {
	if o == nil {
		var ret JetStreamAPIStats
		return ret
	}

	return o.Api
}

// GetApiOk returns a tuple with the Api field value
// and a boolean to check if the value has been set.
func (o *JetStreamAccountStats) GetApiOk() (*JetStreamAPIStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Api, true
}

// SetApi sets field value
func (o *JetStreamAccountStats) SetApi(v JetStreamAPIStats) {
	o.Api = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *JetStreamAccountStats) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamAccountStats) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *JetStreamAccountStats) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *JetStreamAccountStats) SetDomain(v string) {
	o.Domain = &v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *JetStreamAccountStats) GetTiers() map[string]JetStreamTier {
	if o == nil || IsNil(o.Tiers) {
		var ret map[string]JetStreamTier
		return ret
	}
	return *o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamAccountStats) GetTiersOk() (*map[string]JetStreamTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *JetStreamAccountStats) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given map[string]JetStreamTier and assigns it to the Tiers field.
func (o *JetStreamAccountStats) SetTiers(v map[string]JetStreamTier) {
	o.Tiers = &v
}

func (o JetStreamAccountStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JetStreamAccountStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["JetStreamTier"] = o.JetStreamTier
	toSerialize["api"] = o.Api
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}
	return toSerialize, nil
}

type NullableJetStreamAccountStats struct {
	value *JetStreamAccountStats
	isSet bool
}

func (v NullableJetStreamAccountStats) Get() *JetStreamAccountStats {
	return v.value
}

func (v *NullableJetStreamAccountStats) Set(val *JetStreamAccountStats) {
	v.value = val
	v.isSet = true
}

func (v NullableJetStreamAccountStats) IsSet() bool {
	return v.isSet
}

func (v *NullableJetStreamAccountStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJetStreamAccountStats(val *JetStreamAccountStats) *NullableJetStreamAccountStats {
	return &NullableJetStreamAccountStats{value: val, isSet: true}
}

func (v NullableJetStreamAccountStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJetStreamAccountStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
