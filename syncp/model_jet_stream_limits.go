/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JetStreamLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JetStreamLimits{}

// JetStreamLimits struct for JetStreamLimits
type JetStreamLimits struct {
	Consumer           *int64 `json:"consumer,omitempty"`
	DiskMaxStreamBytes *int64 `json:"disk_max_stream_bytes,omitempty"`
	DiskStorage        *int64 `json:"disk_storage,omitempty"`
	MaxAckPending      *int64 `json:"max_ack_pending,omitempty"`
	MaxBytesRequired   *bool  `json:"max_bytes_required,omitempty"`
	MemMaxStreamBytes  *int64 `json:"mem_max_stream_bytes,omitempty"`
	MemStorage         *int64 `json:"mem_storage,omitempty"`
	Streams            *int64 `json:"streams,omitempty"`
}

func (o JetStreamLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Consumer != nil {
		toSerialize["consumer"] = o.Consumer
	}
	if o.DiskMaxStreamBytes != nil {
		toSerialize["disk_max_stream_bytes"] = o.DiskMaxStreamBytes
	}
	if o.DiskStorage != nil {
		toSerialize["disk_storage"] = o.DiskStorage
	}
	if o.MaxAckPending != nil {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if o.MaxBytesRequired != nil {
		toSerialize["max_bytes_required"] = o.MaxBytesRequired
	}
	if o.MemMaxStreamBytes != nil {
		toSerialize["mem_max_stream_bytes"] = o.MemMaxStreamBytes
	}
	if o.MemStorage != nil {
		toSerialize["mem_storage"] = o.MemStorage
	}
	if o.Streams != nil {
		toSerialize["streams"] = o.Streams
	}
	return toSerialize, nil
}
