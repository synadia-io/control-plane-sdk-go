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

// checks if the AccountConnectionsListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountConnectionsListResponse{}

// AccountConnectionsListResponse struct for AccountConnectionsListResponse
type AccountConnectionsListResponse struct {
	Items []Connz `json:"items"`
}

// NewAccountConnectionsListResponse instantiates a new AccountConnectionsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConnectionsListResponse(items []Connz) *AccountConnectionsListResponse {
	this := AccountConnectionsListResponse{}
	this.Items = items
	return &this
}

// NewAccountConnectionsListResponseWithDefaults instantiates a new AccountConnectionsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConnectionsListResponseWithDefaults() *AccountConnectionsListResponse {
	this := AccountConnectionsListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *AccountConnectionsListResponse) GetItems() []Connz {
	if o == nil {
		var ret []Connz
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AccountConnectionsListResponse) GetItemsOk() ([]Connz, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AccountConnectionsListResponse) SetItems(v []Connz) {
	o.Items = v
}

func (o AccountConnectionsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConnectionsListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableAccountConnectionsListResponse struct {
	value *AccountConnectionsListResponse
	isSet bool
}

func (v NullableAccountConnectionsListResponse) Get() *AccountConnectionsListResponse {
	return v.value
}

func (v *NullableAccountConnectionsListResponse) Set(val *AccountConnectionsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConnectionsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConnectionsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConnectionsListResponse(val *AccountConnectionsListResponse) *NullableAccountConnectionsListResponse {
	return &NullableAccountConnectionsListResponse{value: val, isSet: true}
}

func (v NullableAccountConnectionsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConnectionsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
