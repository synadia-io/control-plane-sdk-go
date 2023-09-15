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

// checks if the JetStreamVarzStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JetStreamVarzStats{}

// JetStreamVarzStats struct for JetStreamVarzStats
type JetStreamVarzStats struct {
	Accounts        int32             `json:"accounts"`
	Api             JetStreamAPIStats `json:"api"`
	HaAssets        int32             `json:"ha_assets"`
	Memory          int32             `json:"memory"`
	ReservedMemory  int32             `json:"reserved_memory"`
	ReservedStorage int32             `json:"reserved_storage"`
	Storage         int32             `json:"storage"`
}

// NewJetStreamVarzStats instantiates a new JetStreamVarzStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJetStreamVarzStats(accounts int32, api JetStreamAPIStats, haAssets int32, memory int32, reservedMemory int32, reservedStorage int32, storage int32) *JetStreamVarzStats {
	this := JetStreamVarzStats{}
	this.Accounts = accounts
	this.Api = api
	this.HaAssets = haAssets
	this.Memory = memory
	this.ReservedMemory = reservedMemory
	this.ReservedStorage = reservedStorage
	this.Storage = storage
	return &this
}

// NewJetStreamVarzStatsWithDefaults instantiates a new JetStreamVarzStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJetStreamVarzStatsWithDefaults() *JetStreamVarzStats {
	this := JetStreamVarzStats{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *JetStreamVarzStats) GetAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *JetStreamVarzStats) SetAccounts(v int32) {
	o.Accounts = v
}

// GetApi returns the Api field value
func (o *JetStreamVarzStats) GetApi() JetStreamAPIStats {
	if o == nil {
		var ret JetStreamAPIStats
		return ret
	}

	return o.Api
}

// GetApiOk returns a tuple with the Api field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetApiOk() (*JetStreamAPIStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Api, true
}

// SetApi sets field value
func (o *JetStreamVarzStats) SetApi(v JetStreamAPIStats) {
	o.Api = v
}

// GetHaAssets returns the HaAssets field value
func (o *JetStreamVarzStats) GetHaAssets() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HaAssets
}

// GetHaAssetsOk returns a tuple with the HaAssets field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetHaAssetsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HaAssets, true
}

// SetHaAssets sets field value
func (o *JetStreamVarzStats) SetHaAssets(v int32) {
	o.HaAssets = v
}

// GetMemory returns the Memory field value
func (o *JetStreamVarzStats) GetMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *JetStreamVarzStats) SetMemory(v int32) {
	o.Memory = v
}

// GetReservedMemory returns the ReservedMemory field value
func (o *JetStreamVarzStats) GetReservedMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReservedMemory
}

// GetReservedMemoryOk returns a tuple with the ReservedMemory field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetReservedMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservedMemory, true
}

// SetReservedMemory sets field value
func (o *JetStreamVarzStats) SetReservedMemory(v int32) {
	o.ReservedMemory = v
}

// GetReservedStorage returns the ReservedStorage field value
func (o *JetStreamVarzStats) GetReservedStorage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReservedStorage
}

// GetReservedStorageOk returns a tuple with the ReservedStorage field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetReservedStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservedStorage, true
}

// SetReservedStorage sets field value
func (o *JetStreamVarzStats) SetReservedStorage(v int32) {
	o.ReservedStorage = v
}

// GetStorage returns the Storage field value
func (o *JetStreamVarzStats) GetStorage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *JetStreamVarzStats) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *JetStreamVarzStats) SetStorage(v int32) {
	o.Storage = v
}

func (o JetStreamVarzStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JetStreamVarzStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	toSerialize["api"] = o.Api
	toSerialize["ha_assets"] = o.HaAssets
	toSerialize["memory"] = o.Memory
	toSerialize["reserved_memory"] = o.ReservedMemory
	toSerialize["reserved_storage"] = o.ReservedStorage
	toSerialize["storage"] = o.Storage
	return toSerialize, nil
}

type NullableJetStreamVarzStats struct {
	value *JetStreamVarzStats
	isSet bool
}

func (v NullableJetStreamVarzStats) Get() *JetStreamVarzStats {
	return v.value
}

func (v *NullableJetStreamVarzStats) Set(val *JetStreamVarzStats) {
	v.value = val
	v.isSet = true
}

func (v NullableJetStreamVarzStats) IsSet() bool {
	return v.isSet
}

func (v *NullableJetStreamVarzStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJetStreamVarzStats(val *JetStreamVarzStats) *NullableJetStreamVarzStats {
	return &NullableJetStreamVarzStats{value: val, isSet: true}
}

func (v NullableJetStreamVarzStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJetStreamVarzStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
