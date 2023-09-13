/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the StreamExportCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamExportCreateRequest{}

// StreamExportCreateRequest struct for StreamExportCreateRequest
type StreamExportCreateRequest struct {
	IsPublic   bool   `json:"is_public"`
	StreamName string `json:"stream_name"`
}

// NewStreamExportCreateRequest instantiates a new StreamExportCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamExportCreateRequest(isPublic bool, streamName string) *StreamExportCreateRequest {
	this := StreamExportCreateRequest{}
	this.IsPublic = isPublic
	this.StreamName = streamName
	return &this
}

// NewStreamExportCreateRequestWithDefaults instantiates a new StreamExportCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamExportCreateRequestWithDefaults() *StreamExportCreateRequest {
	this := StreamExportCreateRequest{}
	return &this
}

// GetIsPublic returns the IsPublic field value
func (o *StreamExportCreateRequest) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *StreamExportCreateRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *StreamExportCreateRequest) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetStreamName returns the StreamName field value
func (o *StreamExportCreateRequest) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *StreamExportCreateRequest) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *StreamExportCreateRequest) SetStreamName(v string) {
	o.StreamName = v
}

func (o StreamExportCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamExportCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_public"] = o.IsPublic
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}

type NullableStreamExportCreateRequest struct {
	value *StreamExportCreateRequest
	isSet bool
}

func (v NullableStreamExportCreateRequest) Get() *StreamExportCreateRequest {
	return v.value
}

func (v *NullableStreamExportCreateRequest) Set(val *StreamExportCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamExportCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamExportCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamExportCreateRequest(val *StreamExportCreateRequest) *NullableStreamExportCreateRequest {
	return &NullableStreamExportCreateRequest{value: val, isSet: true}
}

func (v NullableStreamExportCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamExportCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
