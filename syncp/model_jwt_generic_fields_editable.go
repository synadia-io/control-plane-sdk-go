/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
)

// checks if the JwtGenericFieldsEditable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JwtGenericFieldsEditable{}

// JwtGenericFieldsEditable filter jwt.#GenericFields to editable fields only
type JwtGenericFieldsEditable struct {
	// TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments
	Tags []string `json:"tags,omitempty"`
}

// NewJwtGenericFieldsEditable instantiates a new JwtGenericFieldsEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJwtGenericFieldsEditable() *JwtGenericFieldsEditable {
	this := JwtGenericFieldsEditable{}
	return &this
}

// NewJwtGenericFieldsEditableWithDefaults instantiates a new JwtGenericFieldsEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJwtGenericFieldsEditableWithDefaults() *JwtGenericFieldsEditable {
	this := JwtGenericFieldsEditable{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JwtGenericFieldsEditable) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JwtGenericFieldsEditable) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *JwtGenericFieldsEditable) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *JwtGenericFieldsEditable) SetTags(v []string) {
	o.Tags = v
}

func (o JwtGenericFieldsEditable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JwtGenericFieldsEditable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableJwtGenericFieldsEditable struct {
	value *JwtGenericFieldsEditable
	isSet bool
}

func (v NullableJwtGenericFieldsEditable) Get() *JwtGenericFieldsEditable {
	return v.value
}

func (v *NullableJwtGenericFieldsEditable) Set(val *JwtGenericFieldsEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableJwtGenericFieldsEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableJwtGenericFieldsEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwtGenericFieldsEditable(val *JwtGenericFieldsEditable) *NullableJwtGenericFieldsEditable {
	return &NullableJwtGenericFieldsEditable{value: val, isSet: true}
}

func (v NullableJwtGenericFieldsEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwtGenericFieldsEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
