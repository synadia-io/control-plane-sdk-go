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

// checks if the StreamShareListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamShareListResponse{}

// StreamShareListResponse struct for StreamShareListResponse
type StreamShareListResponse struct {
	Items []StreamShareViewResponse `json:"items"`
}

// NewStreamShareListResponse instantiates a new StreamShareListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamShareListResponse(items []StreamShareViewResponse) *StreamShareListResponse {
	this := StreamShareListResponse{}
	this.Items = items
	return &this
}

// NewStreamShareListResponseWithDefaults instantiates a new StreamShareListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamShareListResponseWithDefaults() *StreamShareListResponse {
	this := StreamShareListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *StreamShareListResponse) GetItems() []StreamShareViewResponse {
	if o == nil {
		var ret []StreamShareViewResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *StreamShareListResponse) GetItemsOk() ([]StreamShareViewResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *StreamShareListResponse) SetItems(v []StreamShareViewResponse) {
	o.Items = v
}

func (o StreamShareListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamShareListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableStreamShareListResponse struct {
	value *StreamShareListResponse
	isSet bool
}

func (v NullableStreamShareListResponse) Get() *StreamShareListResponse {
	return v.value
}

func (v *NullableStreamShareListResponse) Set(val *StreamShareListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamShareListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamShareListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamShareListResponse(val *StreamShareListResponse) *NullableStreamShareListResponse {
	return &NullableStreamShareListResponse{value: val, isSet: true}
}

func (v NullableStreamShareListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamShareListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
