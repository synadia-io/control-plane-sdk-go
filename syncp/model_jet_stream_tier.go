/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JetStreamTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JetStreamTier{}

// JetStreamTier struct for JetStreamTier
type JetStreamTier struct {
	Consumers int64                  `json:"consumers"`
	Limits    JetStreamAccountLimits `json:"limits"`
	Memory    uint64                 `json:"memory"`
	Storage   uint64                 `json:"storage"`
	Streams   int64                  `json:"streams"`
}

func (o JetStreamTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consumers"] = o.Consumers
	toSerialize["limits"] = o.Limits
	toSerialize["memory"] = o.Memory
	toSerialize["storage"] = o.Storage
	toSerialize["streams"] = o.Streams
	return toSerialize, nil
}
