/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JetStreamStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JetStreamStats{}

// JetStreamStats Statistics about JetStream for this server.
type JetStreamStats struct {
	Accounts        int64             `json:"accounts"`
	Api             JetStreamAPIStats `json:"api"`
	HaAssets        int64             `json:"ha_assets"`
	Memory          uint64            `json:"memory"`
	ReservedMemory  uint64            `json:"reserved_memory"`
	ReservedStorage uint64            `json:"reserved_storage"`
	Storage         uint64            `json:"storage"`
}

func (o JetStreamStats) ToMap() (map[string]interface{}, error) {
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
