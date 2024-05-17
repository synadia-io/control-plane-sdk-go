/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the OperatorLimitsPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorLimitsPatch{}

// OperatorLimitsPatch struct for OperatorLimitsPatch
type OperatorLimitsPatch struct {
	Conn               *int64                      `json:"conn,omitempty"`
	Consumer           *int64                      `json:"consumer,omitempty"`
	Data               *int64                      `json:"data,omitempty"`
	DisallowBearer     *bool                       `json:"disallow_bearer,omitempty"`
	DiskMaxStreamBytes *int64                      `json:"disk_max_stream_bytes,omitempty"`
	DiskStorage        *int64                      `json:"disk_storage,omitempty"`
	Exports            *int64                      `json:"exports,omitempty"`
	Imports            *int64                      `json:"imports,omitempty"`
	Leaf               *int64                      `json:"leaf,omitempty"`
	MaxAckPending      *int64                      `json:"max_ack_pending,omitempty"`
	MaxBytesRequired   *bool                       `json:"max_bytes_required,omitempty"`
	MemMaxStreamBytes  *int64                      `json:"mem_max_stream_bytes,omitempty"`
	MemStorage         *int64                      `json:"mem_storage,omitempty"`
	Payload            *int64                      `json:"payload,omitempty"`
	Streams            *int64                      `json:"streams,omitempty"`
	Subs               *int64                      `json:"subs,omitempty"`
	TieredLimits       map[string]*JetStreamLimits `json:"tiered_limits,omitempty"`
	Wildcards          *bool                       `json:"wildcards,omitempty"`
}

func (o OperatorLimitsPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Conn != nil {
		toSerialize["conn"] = o.Conn
	}
	if o.Consumer != nil {
		toSerialize["consumer"] = o.Consumer
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DisallowBearer != nil {
		toSerialize["disallow_bearer"] = o.DisallowBearer
	}
	if o.DiskMaxStreamBytes != nil {
		toSerialize["disk_max_stream_bytes"] = o.DiskMaxStreamBytes
	}
	if o.DiskStorage != nil {
		toSerialize["disk_storage"] = o.DiskStorage
	}
	if o.Exports != nil {
		toSerialize["exports"] = o.Exports
	}
	if o.Imports != nil {
		toSerialize["imports"] = o.Imports
	}
	if o.Leaf != nil {
		toSerialize["leaf"] = o.Leaf
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
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Streams != nil {
		toSerialize["streams"] = o.Streams
	}
	if o.Subs != nil {
		toSerialize["subs"] = o.Subs
	}
	if o.TieredLimits != nil {
		toSerialize["tiered_limits"] = o.TieredLimits
	}
	if o.Wildcards != nil {
		toSerialize["wildcards"] = o.Wildcards
	}
	return toSerialize, nil
}
