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

// NewJetStreamLimits instantiates a new JetStreamLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJetStreamLimits() *JetStreamLimits {
	this := JetStreamLimits{}
	return &this
}

// NewJetStreamLimitsWithDefaults instantiates a new JetStreamLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJetStreamLimitsWithDefaults() *JetStreamLimits {
	this := JetStreamLimits{}
	return &this
}

// GetConsumer returns the Consumer field value if set, zero value otherwise.
func (o *JetStreamLimits) GetConsumer() int64 {
	if o == nil || IsNil(o.Consumer) {
		var ret int64
		return ret
	}
	return *o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetConsumerOk() (*int64, bool) {
	if o == nil || IsNil(o.Consumer) {
		return nil, false
	}
	return o.Consumer, true
}

// HasConsumer returns a boolean if a field has been set.
func (o *JetStreamLimits) HasConsumer() bool {
	if o != nil && !IsNil(o.Consumer) {
		return true
	}

	return false
}

// SetConsumer gets a reference to the given int64 and assigns it to the Consumer field.
func (o *JetStreamLimits) SetConsumer(v int64) {
	o.Consumer = &v
}

// GetDiskMaxStreamBytes returns the DiskMaxStreamBytes field value if set, zero value otherwise.
func (o *JetStreamLimits) GetDiskMaxStreamBytes() int64 {
	if o == nil || IsNil(o.DiskMaxStreamBytes) {
		var ret int64
		return ret
	}
	return *o.DiskMaxStreamBytes
}

// GetDiskMaxStreamBytesOk returns a tuple with the DiskMaxStreamBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetDiskMaxStreamBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskMaxStreamBytes) {
		return nil, false
	}
	return o.DiskMaxStreamBytes, true
}

// HasDiskMaxStreamBytes returns a boolean if a field has been set.
func (o *JetStreamLimits) HasDiskMaxStreamBytes() bool {
	if o != nil && !IsNil(o.DiskMaxStreamBytes) {
		return true
	}

	return false
}

// SetDiskMaxStreamBytes gets a reference to the given int64 and assigns it to the DiskMaxStreamBytes field.
func (o *JetStreamLimits) SetDiskMaxStreamBytes(v int64) {
	o.DiskMaxStreamBytes = &v
}

// GetDiskStorage returns the DiskStorage field value if set, zero value otherwise.
func (o *JetStreamLimits) GetDiskStorage() int64 {
	if o == nil || IsNil(o.DiskStorage) {
		var ret int64
		return ret
	}
	return *o.DiskStorage
}

// GetDiskStorageOk returns a tuple with the DiskStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetDiskStorageOk() (*int64, bool) {
	if o == nil || IsNil(o.DiskStorage) {
		return nil, false
	}
	return o.DiskStorage, true
}

// HasDiskStorage returns a boolean if a field has been set.
func (o *JetStreamLimits) HasDiskStorage() bool {
	if o != nil && !IsNil(o.DiskStorage) {
		return true
	}

	return false
}

// SetDiskStorage gets a reference to the given int64 and assigns it to the DiskStorage field.
func (o *JetStreamLimits) SetDiskStorage(v int64) {
	o.DiskStorage = &v
}

// GetMaxAckPending returns the MaxAckPending field value if set, zero value otherwise.
func (o *JetStreamLimits) GetMaxAckPending() int64 {
	if o == nil || IsNil(o.MaxAckPending) {
		var ret int64
		return ret
	}
	return *o.MaxAckPending
}

// GetMaxAckPendingOk returns a tuple with the MaxAckPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetMaxAckPendingOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxAckPending) {
		return nil, false
	}
	return o.MaxAckPending, true
}

// HasMaxAckPending returns a boolean if a field has been set.
func (o *JetStreamLimits) HasMaxAckPending() bool {
	if o != nil && !IsNil(o.MaxAckPending) {
		return true
	}

	return false
}

// SetMaxAckPending gets a reference to the given int64 and assigns it to the MaxAckPending field.
func (o *JetStreamLimits) SetMaxAckPending(v int64) {
	o.MaxAckPending = &v
}

// GetMaxBytesRequired returns the MaxBytesRequired field value if set, zero value otherwise.
func (o *JetStreamLimits) GetMaxBytesRequired() bool {
	if o == nil || IsNil(o.MaxBytesRequired) {
		var ret bool
		return ret
	}
	return *o.MaxBytesRequired
}

// GetMaxBytesRequiredOk returns a tuple with the MaxBytesRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetMaxBytesRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.MaxBytesRequired) {
		return nil, false
	}
	return o.MaxBytesRequired, true
}

// HasMaxBytesRequired returns a boolean if a field has been set.
func (o *JetStreamLimits) HasMaxBytesRequired() bool {
	if o != nil && !IsNil(o.MaxBytesRequired) {
		return true
	}

	return false
}

// SetMaxBytesRequired gets a reference to the given bool and assigns it to the MaxBytesRequired field.
func (o *JetStreamLimits) SetMaxBytesRequired(v bool) {
	o.MaxBytesRequired = &v
}

// GetMemMaxStreamBytes returns the MemMaxStreamBytes field value if set, zero value otherwise.
func (o *JetStreamLimits) GetMemMaxStreamBytes() int64 {
	if o == nil || IsNil(o.MemMaxStreamBytes) {
		var ret int64
		return ret
	}
	return *o.MemMaxStreamBytes
}

// GetMemMaxStreamBytesOk returns a tuple with the MemMaxStreamBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetMemMaxStreamBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.MemMaxStreamBytes) {
		return nil, false
	}
	return o.MemMaxStreamBytes, true
}

// HasMemMaxStreamBytes returns a boolean if a field has been set.
func (o *JetStreamLimits) HasMemMaxStreamBytes() bool {
	if o != nil && !IsNil(o.MemMaxStreamBytes) {
		return true
	}

	return false
}

// SetMemMaxStreamBytes gets a reference to the given int64 and assigns it to the MemMaxStreamBytes field.
func (o *JetStreamLimits) SetMemMaxStreamBytes(v int64) {
	o.MemMaxStreamBytes = &v
}

// GetMemStorage returns the MemStorage field value if set, zero value otherwise.
func (o *JetStreamLimits) GetMemStorage() int64 {
	if o == nil || IsNil(o.MemStorage) {
		var ret int64
		return ret
	}
	return *o.MemStorage
}

// GetMemStorageOk returns a tuple with the MemStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetMemStorageOk() (*int64, bool) {
	if o == nil || IsNil(o.MemStorage) {
		return nil, false
	}
	return o.MemStorage, true
}

// HasMemStorage returns a boolean if a field has been set.
func (o *JetStreamLimits) HasMemStorage() bool {
	if o != nil && !IsNil(o.MemStorage) {
		return true
	}

	return false
}

// SetMemStorage gets a reference to the given int64 and assigns it to the MemStorage field.
func (o *JetStreamLimits) SetMemStorage(v int64) {
	o.MemStorage = &v
}

// GetStreams returns the Streams field value if set, zero value otherwise.
func (o *JetStreamLimits) GetStreams() int64 {
	if o == nil || IsNil(o.Streams) {
		var ret int64
		return ret
	}
	return *o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JetStreamLimits) GetStreamsOk() (*int64, bool) {
	if o == nil || IsNil(o.Streams) {
		return nil, false
	}
	return o.Streams, true
}

// HasStreams returns a boolean if a field has been set.
func (o *JetStreamLimits) HasStreams() bool {
	if o != nil && !IsNil(o.Streams) {
		return true
	}

	return false
}

// SetStreams gets a reference to the given int64 and assigns it to the Streams field.
func (o *JetStreamLimits) SetStreams(v int64) {
	o.Streams = &v
}

func (o JetStreamLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JetStreamLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Consumer) {
		toSerialize["consumer"] = o.Consumer
	}
	if !IsNil(o.DiskMaxStreamBytes) {
		toSerialize["disk_max_stream_bytes"] = o.DiskMaxStreamBytes
	}
	if !IsNil(o.DiskStorage) {
		toSerialize["disk_storage"] = o.DiskStorage
	}
	if !IsNil(o.MaxAckPending) {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if !IsNil(o.MaxBytesRequired) {
		toSerialize["max_bytes_required"] = o.MaxBytesRequired
	}
	if !IsNil(o.MemMaxStreamBytes) {
		toSerialize["mem_max_stream_bytes"] = o.MemMaxStreamBytes
	}
	if !IsNil(o.MemStorage) {
		toSerialize["mem_storage"] = o.MemStorage
	}
	if !IsNil(o.Streams) {
		toSerialize["streams"] = o.Streams
	}
	return toSerialize, nil
}

type NullableJetStreamLimits struct {
	value *JetStreamLimits
	isSet bool
}

func (v NullableJetStreamLimits) Get() *JetStreamLimits {
	return v.value
}

func (v *NullableJetStreamLimits) Set(val *JetStreamLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableJetStreamLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableJetStreamLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJetStreamLimits(val *JetStreamLimits) *NullableJetStreamLimits {
	return &NullableJetStreamLimits{value: val, isSet: true}
}

func (v NullableJetStreamLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJetStreamLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}