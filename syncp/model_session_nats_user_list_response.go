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

// checks if the SessionNatsUserListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionNatsUserListResponse{}

// SessionNatsUserListResponse struct for SessionNatsUserListResponse
type SessionNatsUserListResponse struct {
	Items []SessionNatsUserListResponseItemsInner `json:"items"`
}

// NewSessionNatsUserListResponse instantiates a new SessionNatsUserListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionNatsUserListResponse(items []SessionNatsUserListResponseItemsInner) *SessionNatsUserListResponse {
	this := SessionNatsUserListResponse{}
	this.Items = items
	return &this
}

// NewSessionNatsUserListResponseWithDefaults instantiates a new SessionNatsUserListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionNatsUserListResponseWithDefaults() *SessionNatsUserListResponse {
	this := SessionNatsUserListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SessionNatsUserListResponse) GetItems() []SessionNatsUserListResponseItemsInner {
	if o == nil {
		var ret []SessionNatsUserListResponseItemsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SessionNatsUserListResponse) GetItemsOk() ([]SessionNatsUserListResponseItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SessionNatsUserListResponse) SetItems(v []SessionNatsUserListResponseItemsInner) {
	o.Items = v
}

func (o SessionNatsUserListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionNatsUserListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSessionNatsUserListResponse struct {
	value *SessionNatsUserListResponse
	isSet bool
}

func (v NullableSessionNatsUserListResponse) Get() *SessionNatsUserListResponse {
	return v.value
}

func (v *NullableSessionNatsUserListResponse) Set(val *SessionNatsUserListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionNatsUserListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionNatsUserListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionNatsUserListResponse(val *SessionNatsUserListResponse) *NullableSessionNatsUserListResponse {
	return &NullableSessionNatsUserListResponse{value: val, isSet: true}
}

func (v NullableSessionNatsUserListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionNatsUserListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}