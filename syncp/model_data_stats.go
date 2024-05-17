/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the DataStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStats{}

// DataStats DataStats reports how may msg and bytes. Applicable for both sent and received.
type DataStats struct {
	Bytes int64 `json:"bytes"`
	Msgs  int64 `json:"msgs"`
}

func (o DataStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bytes"] = o.Bytes
	toSerialize["msgs"] = o.Msgs
	return toSerialize, nil
}
