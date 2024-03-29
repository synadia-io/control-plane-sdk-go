/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JetStreamAPIStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JetStreamAPIStats{}

// JetStreamAPIStats struct for JetStreamAPIStats
type JetStreamAPIStats struct {
	Errors   int32  `json:"errors"`
	Inflight *int32 `json:"inflight,omitempty"`
	Total    int32  `json:"total"`
}

func (o JetStreamAPIStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	if o.Inflight != nil {
		toSerialize["inflight"] = o.Inflight
	}
	toSerialize["total"] = o.Total
	return toSerialize, nil
}
