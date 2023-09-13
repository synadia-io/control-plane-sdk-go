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

// checks if the SubjectImportListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectImportListResponse{}

// SubjectImportListResponse struct for SubjectImportListResponse
type SubjectImportListResponse struct {
	Items []SubjectImportViewResponse `json:"items"`
}

// NewSubjectImportListResponse instantiates a new SubjectImportListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectImportListResponse(items []SubjectImportViewResponse) *SubjectImportListResponse {
	this := SubjectImportListResponse{}
	this.Items = items
	return &this
}

// NewSubjectImportListResponseWithDefaults instantiates a new SubjectImportListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectImportListResponseWithDefaults() *SubjectImportListResponse {
	this := SubjectImportListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *SubjectImportListResponse) GetItems() []SubjectImportViewResponse {
	if o == nil {
		var ret []SubjectImportViewResponse
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SubjectImportListResponse) GetItemsOk() ([]SubjectImportViewResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SubjectImportListResponse) SetItems(v []SubjectImportViewResponse) {
	o.Items = v
}

func (o SubjectImportListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectImportListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableSubjectImportListResponse struct {
	value *SubjectImportListResponse
	isSet bool
}

func (v NullableSubjectImportListResponse) Get() *SubjectImportListResponse {
	return v.value
}

func (v *NullableSubjectImportListResponse) Set(val *SubjectImportListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectImportListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectImportListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectImportListResponse(val *SubjectImportListResponse) *NullableSubjectImportListResponse {
	return &NullableSubjectImportListResponse{value: val, isSet: true}
}

func (v NullableSubjectImportListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectImportListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
