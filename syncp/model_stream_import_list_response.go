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

// checks if the StreamImportListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamImportListResponse{}

// StreamImportListResponse struct for StreamImportListResponse
type StreamImportListResponse struct {
	Items []StreamImportViewResponse `json:"items"`
}

// NewStreamImportListResponse instantiates a new StreamImportListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamImportListResponse(items []StreamImportViewResponse) *StreamImportListResponse {
	this := StreamImportListResponse{}
	this.Items = items
	return &this
}

// NewStreamImportListResponseWithDefaults instantiates a new StreamImportListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamImportListResponseWithDefaults() *StreamImportListResponse {
	this := StreamImportListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *StreamImportListResponse) GetItems() []StreamImportViewResponse {
	if o == nil {
		var ret []StreamImportViewResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *StreamImportListResponse) GetItemsOk() ([]StreamImportViewResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *StreamImportListResponse) SetItems(v []StreamImportViewResponse) {
	o.Items = v
}

func (o StreamImportListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamImportListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableStreamImportListResponse struct {
	value *StreamImportListResponse
	isSet bool
}

func (v NullableStreamImportListResponse) Get() *StreamImportListResponse {
	return v.value
}

func (v *NullableStreamImportListResponse) Set(val *StreamImportListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamImportListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamImportListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamImportListResponse(val *StreamImportListResponse) *NullableStreamImportListResponse {
	return &NullableStreamImportListResponse{value: val, isSet: true}
}

func (v NullableStreamImportListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamImportListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
