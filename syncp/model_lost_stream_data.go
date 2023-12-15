/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the LostStreamData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LostStreamData{}

// LostStreamData LostStreamData indicates msgs that have been lost during file checks and recover due to corruption
type LostStreamData struct {
	// How many bytes were lost
	Bytes int32 `json:"bytes"`
	// Message IDs of lost messages
	Msgs []int32 `json:"msgs"`
}

func (o LostStreamData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bytes"] = o.Bytes
	toSerialize["msgs"] = o.Msgs
	return toSerialize, nil
}
