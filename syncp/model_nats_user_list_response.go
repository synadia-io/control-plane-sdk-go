/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"fmt"
)

// checks if the NatsUserListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserListResponse{}

// NatsUserListResponse struct for NatsUserListResponse
type NatsUserListResponse struct {
	Items []NatsUserViewResponse `json:"items"`
}

type _NatsUserListResponse NatsUserListResponse

// NewNatsUserListResponse instantiates a new NatsUserListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatsUserListResponse(items []NatsUserViewResponse) *NatsUserListResponse {
	this := NatsUserListResponse{}
	this.Items = items
	return &this
}

// NewNatsUserListResponseWithDefaults instantiates a new NatsUserListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatsUserListResponseWithDefaults() *NatsUserListResponse {
	this := NatsUserListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *NatsUserListResponse) GetItems() []NatsUserViewResponse {
	if o == nil {
		var ret []NatsUserViewResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NatsUserListResponse) GetItemsOk() ([]NatsUserViewResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NatsUserListResponse) SetItems(v []NatsUserViewResponse) {
	o.Items = v
}

func (o NatsUserListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NatsUserListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *NatsUserListResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNatsUserListResponse := _NatsUserListResponse{}

	err = json.Unmarshal(bytes, &varNatsUserListResponse)

	if err != nil {
		return err
	}

	*o = NatsUserListResponse(varNatsUserListResponse)

	return err
}

type NullableNatsUserListResponse struct {
	value *NatsUserListResponse
	isSet bool
}

func (v NullableNatsUserListResponse) Get() *NatsUserListResponse {
	return v.value
}

func (v *NullableNatsUserListResponse) Set(val *NatsUserListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNatsUserListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNatsUserListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatsUserListResponse(val *NatsUserListResponse) *NullableNatsUserListResponse {
	return &NullableNatsUserListResponse{value: val, isSet: true}
}

func (v NullableNatsUserListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatsUserListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
