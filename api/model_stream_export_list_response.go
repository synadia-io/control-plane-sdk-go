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

// checks if the StreamExportListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamExportListResponse{}

// StreamExportListResponse struct for StreamExportListResponse
type StreamExportListResponse struct {
	Items []StreamExportViewResponse `json:"items"`
}

// NewStreamExportListResponse instantiates a new StreamExportListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamExportListResponse(items []StreamExportViewResponse) *StreamExportListResponse {
	this := StreamExportListResponse{}
	this.Items = items
	return &this
}

// NewStreamExportListResponseWithDefaults instantiates a new StreamExportListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamExportListResponseWithDefaults() *StreamExportListResponse {
	this := StreamExportListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *StreamExportListResponse) GetItems() []StreamExportViewResponse {
	if o == nil {
		var ret []StreamExportViewResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *StreamExportListResponse) GetItemsOk() ([]StreamExportViewResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *StreamExportListResponse) SetItems(v []StreamExportViewResponse) {
	o.Items = v
}

func (o StreamExportListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamExportListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableStreamExportListResponse struct {
	value *StreamExportListResponse
	isSet bool
}

func (v NullableStreamExportListResponse) Get() *StreamExportListResponse {
	return v.value
}

func (v *NullableStreamExportListResponse) Set(val *StreamExportListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamExportListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamExportListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamExportListResponse(val *StreamExportListResponse) *NullableStreamExportListResponse {
	return &NullableStreamExportListResponse{value: val, isSet: true}
}

func (v NullableStreamExportListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamExportListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
