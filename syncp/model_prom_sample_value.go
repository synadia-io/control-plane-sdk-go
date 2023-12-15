/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the PromSampleValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromSampleValue{}

// PromSampleValue struct for PromSampleValue
type PromSampleValue struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

func (o PromSampleValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["value"] = o.Value
	return toSerialize, nil
}
