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

// checks if the ServerStatsJetstream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStatsJetstream{}

// ServerStatsJetstream struct for ServerStatsJetstream
type ServerStatsJetstream struct {
	Config NullableJetStreamVarzConfig `json:"config,omitempty"`
	Meta   NullableJetStreamVarzMeta   `json:"meta,omitempty"`
	Stats  NullableJetStreamVarzStats  `json:"stats,omitempty"`
}

// NewServerStatsJetstream instantiates a new ServerStatsJetstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStatsJetstream() *ServerStatsJetstream {
	this := ServerStatsJetstream{}
	return &this
}

// NewServerStatsJetstreamWithDefaults instantiates a new ServerStatsJetstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatsJetstreamWithDefaults() *ServerStatsJetstream {
	this := ServerStatsJetstream{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStatsJetstream) GetConfig() JetStreamVarzConfig {
	if o == nil || IsNil(o.Config.Get()) {
		var ret JetStreamVarzConfig
		return ret
	}
	return *o.Config.Get()
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStatsJetstream) GetConfigOk() (*JetStreamVarzConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Config.Get(), o.Config.IsSet()
}

// HasConfig returns a boolean if a field has been set.
func (o *ServerStatsJetstream) HasConfig() bool {
	if o != nil && o.Config.IsSet() {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NullableJetStreamVarzConfig and assigns it to the Config field.
func (o *ServerStatsJetstream) SetConfig(v JetStreamVarzConfig) {
	o.Config.Set(&v)
}

// SetConfigNil sets the value for Config to be an explicit nil
func (o *ServerStatsJetstream) SetConfigNil() {
	o.Config.Set(nil)
}

// UnsetConfig ensures that no value is present for Config, not even an explicit nil
func (o *ServerStatsJetstream) UnsetConfig() {
	o.Config.Unset()
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStatsJetstream) GetMeta() JetStreamVarzMeta {
	if o == nil || IsNil(o.Meta.Get()) {
		var ret JetStreamVarzMeta
		return ret
	}
	return *o.Meta.Get()
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStatsJetstream) GetMetaOk() (*JetStreamVarzMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta.Get(), o.Meta.IsSet()
}

// HasMeta returns a boolean if a field has been set.
func (o *ServerStatsJetstream) HasMeta() bool {
	if o != nil && o.Meta.IsSet() {
		return true
	}

	return false
}

// SetMeta gets a reference to the given NullableJetStreamVarzMeta and assigns it to the Meta field.
func (o *ServerStatsJetstream) SetMeta(v JetStreamVarzMeta) {
	o.Meta.Set(&v)
}

// SetMetaNil sets the value for Meta to be an explicit nil
func (o *ServerStatsJetstream) SetMetaNil() {
	o.Meta.Set(nil)
}

// UnsetMeta ensures that no value is present for Meta, not even an explicit nil
func (o *ServerStatsJetstream) UnsetMeta() {
	o.Meta.Unset()
}

// GetStats returns the Stats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStatsJetstream) GetStats() JetStreamVarzStats {
	if o == nil || IsNil(o.Stats.Get()) {
		var ret JetStreamVarzStats
		return ret
	}
	return *o.Stats.Get()
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStatsJetstream) GetStatsOk() (*JetStreamVarzStats, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stats.Get(), o.Stats.IsSet()
}

// HasStats returns a boolean if a field has been set.
func (o *ServerStatsJetstream) HasStats() bool {
	if o != nil && o.Stats.IsSet() {
		return true
	}

	return false
}

// SetStats gets a reference to the given NullableJetStreamVarzStats and assigns it to the Stats field.
func (o *ServerStatsJetstream) SetStats(v JetStreamVarzStats) {
	o.Stats.Set(&v)
}

// SetStatsNil sets the value for Stats to be an explicit nil
func (o *ServerStatsJetstream) SetStatsNil() {
	o.Stats.Set(nil)
}

// UnsetStats ensures that no value is present for Stats, not even an explicit nil
func (o *ServerStatsJetstream) UnsetStats() {
	o.Stats.Unset()
}

func (o ServerStatsJetstream) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStatsJetstream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Config.IsSet() {
		toSerialize["config"] = o.Config.Get()
	}
	if o.Meta.IsSet() {
		toSerialize["meta"] = o.Meta.Get()
	}
	if o.Stats.IsSet() {
		toSerialize["stats"] = o.Stats.Get()
	}
	return toSerialize, nil
}

type NullableServerStatsJetstream struct {
	value *ServerStatsJetstream
	isSet bool
}

func (v NullableServerStatsJetstream) Get() *ServerStatsJetstream {
	return v.value
}

func (v *NullableServerStatsJetstream) Set(val *ServerStatsJetstream) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStatsJetstream) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStatsJetstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStatsJetstream(val *ServerStatsJetstream) *NullableServerStatsJetstream {
	return &NullableServerStatsJetstream{value: val, isSet: true}
}

func (v NullableServerStatsJetstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStatsJetstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
