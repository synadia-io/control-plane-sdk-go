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

// checks if the JSApiError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSApiError{}

// JSApiError struct for JSApiError
type JSApiError struct {
	Code        int32   `json:"code"`
	Description *string `json:"description,omitempty"`
	ErrCode     *int32  `json:"err_code,omitempty"`
}

// NewJSApiError instantiates a new JSApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSApiError(code int32) *JSApiError {
	this := JSApiError{}
	this.Code = code
	return &this
}

// NewJSApiErrorWithDefaults instantiates a new JSApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSApiErrorWithDefaults() *JSApiError {
	this := JSApiError{}
	return &this
}

// GetCode returns the Code field value
func (o *JSApiError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *JSApiError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *JSApiError) SetCode(v int32) {
	o.Code = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JSApiError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSApiError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JSApiError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JSApiError) SetDescription(v string) {
	o.Description = &v
}

// GetErrCode returns the ErrCode field value if set, zero value otherwise.
func (o *JSApiError) GetErrCode() int32 {
	if o == nil || IsNil(o.ErrCode) {
		var ret int32
		return ret
	}
	return *o.ErrCode
}

// GetErrCodeOk returns a tuple with the ErrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSApiError) GetErrCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrCode) {
		return nil, false
	}
	return o.ErrCode, true
}

// HasErrCode returns a boolean if a field has been set.
func (o *JSApiError) HasErrCode() bool {
	if o != nil && !IsNil(o.ErrCode) {
		return true
	}

	return false
}

// SetErrCode gets a reference to the given int32 and assigns it to the ErrCode field.
func (o *JSApiError) SetErrCode(v int32) {
	o.ErrCode = &v
}

func (o JSApiError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSApiError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ErrCode) {
		toSerialize["err_code"] = o.ErrCode
	}
	return toSerialize, nil
}

type NullableJSApiError struct {
	value *JSApiError
	isSet bool
}

func (v NullableJSApiError) Get() *JSApiError {
	return v.value
}

func (v *NullableJSApiError) Set(val *JSApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableJSApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableJSApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSApiError(val *JSApiError) *NullableJSApiError {
	return &NullableJSApiError{value: val, isSet: true}
}

func (v NullableJSApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
