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

// checks if the AccountSearchListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSearchListResponse{}

// AccountSearchListResponse struct for AccountSearchListResponse
type AccountSearchListResponse struct {
	Items []AccountSearch `json:"items"`
}

// NewAccountSearchListResponse instantiates a new AccountSearchListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSearchListResponse(items []AccountSearch) *AccountSearchListResponse {
	this := AccountSearchListResponse{}
	this.Items = items
	return &this
}

// NewAccountSearchListResponseWithDefaults instantiates a new AccountSearchListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSearchListResponseWithDefaults() *AccountSearchListResponse {
	this := AccountSearchListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *AccountSearchListResponse) GetItems() []AccountSearch {
	if o == nil {
		var ret []AccountSearch
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AccountSearchListResponse) GetItemsOk() ([]AccountSearch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AccountSearchListResponse) SetItems(v []AccountSearch) {
	o.Items = v
}

func (o AccountSearchListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSearchListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableAccountSearchListResponse struct {
	value *AccountSearchListResponse
	isSet bool
}

func (v NullableAccountSearchListResponse) Get() *AccountSearchListResponse {
	return v.value
}

func (v *NullableAccountSearchListResponse) Set(val *AccountSearchListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSearchListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSearchListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSearchListResponse(val *AccountSearchListResponse) *NullableAccountSearchListResponse {
	return &NullableAccountSearchListResponse{value: val, isSet: true}
}

func (v NullableAccountSearchListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSearchListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
