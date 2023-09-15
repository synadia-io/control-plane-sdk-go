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

// checks if the SigningKeyListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyListResponse{}

// SigningKeyListResponse struct for SigningKeyListResponse
type SigningKeyListResponse struct {
	Items []SigningKeyViewResponse `json:"items"`
}

// NewSigningKeyListResponse instantiates a new SigningKeyListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKeyListResponse(items []SigningKeyViewResponse) *SigningKeyListResponse {
	this := SigningKeyListResponse{}
	this.Items = items
	return &this
}

// NewSigningKeyListResponseWithDefaults instantiates a new SigningKeyListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeyListResponseWithDefaults() *SigningKeyListResponse {
	this := SigningKeyListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SigningKeyListResponse) GetItems() []SigningKeyViewResponse {
	if o == nil {
		var ret []SigningKeyViewResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SigningKeyListResponse) GetItemsOk() ([]SigningKeyViewResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SigningKeyListResponse) SetItems(v []SigningKeyViewResponse) {
	o.Items = v
}

func (o SigningKeyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningKeyListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSigningKeyListResponse struct {
	value *SigningKeyListResponse
	isSet bool
}

func (v NullableSigningKeyListResponse) Get() *SigningKeyListResponse {
	return v.value
}

func (v *NullableSigningKeyListResponse) Set(val *SigningKeyListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningKeyListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningKeyListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningKeyListResponse(val *SigningKeyListResponse) *NullableSigningKeyListResponse {
	return &NullableSigningKeyListResponse{value: val, isSet: true}
}

func (v NullableSigningKeyListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningKeyListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
