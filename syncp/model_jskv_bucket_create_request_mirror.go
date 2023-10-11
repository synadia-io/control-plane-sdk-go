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

// checks if the JSKVBucketCreateRequestMirror type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSKVBucketCreateRequestMirror{}

// JSKVBucketCreateRequestMirror struct for JSKVBucketCreateRequestMirror
type JSKVBucketCreateRequestMirror struct {
	External      NullableStreamSourceExternal `json:"external,omitempty"`
	FilterSubject *string                      `json:"filter_subject,omitempty"`
	Name          string                       `json:"name"`
	OptStartSeq   *int32                       `json:"opt_start_seq,omitempty"`
	OptStartTime  NullableString               `json:"opt_start_time,omitempty"`
}

// NewJSKVBucketCreateRequestMirror instantiates a new JSKVBucketCreateRequestMirror object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSKVBucketCreateRequestMirror(name string) *JSKVBucketCreateRequestMirror {
	this := JSKVBucketCreateRequestMirror{}
	this.Name = name
	return &this
}

// NewJSKVBucketCreateRequestMirrorWithDefaults instantiates a new JSKVBucketCreateRequestMirror object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSKVBucketCreateRequestMirrorWithDefaults() *JSKVBucketCreateRequestMirror {
	this := JSKVBucketCreateRequestMirror{}
	return &this
}

// GetExternal returns the External field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSKVBucketCreateRequestMirror) GetExternal() StreamSourceExternal {
	if o == nil || IsNil(o.External.Get()) {
		var ret StreamSourceExternal
		return ret
	}
	return *o.External.Get()
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSKVBucketCreateRequestMirror) GetExternalOk() (*StreamSourceExternal, bool) {
	if o == nil {
		return nil, false
	}
	return o.External.Get(), o.External.IsSet()
}

// HasExternal returns a boolean if a field has been set.
func (o *JSKVBucketCreateRequestMirror) HasExternal() bool {
	if o != nil && o.External.IsSet() {
		return true
	}

	return false
}

// SetExternal gets a reference to the given NullableStreamSourceExternal and assigns it to the External field.
func (o *JSKVBucketCreateRequestMirror) SetExternal(v StreamSourceExternal) {
	o.External.Set(&v)
}

// SetExternalNil sets the value for External to be an explicit nil
func (o *JSKVBucketCreateRequestMirror) SetExternalNil() {
	o.External.Set(nil)
}

// UnsetExternal ensures that no value is present for External, not even an explicit nil
func (o *JSKVBucketCreateRequestMirror) UnsetExternal() {
	o.External.Unset()
}

// GetFilterSubject returns the FilterSubject field value if set, zero value otherwise.
func (o *JSKVBucketCreateRequestMirror) GetFilterSubject() string {
	if o == nil || IsNil(o.FilterSubject) {
		var ret string
		return ret
	}
	return *o.FilterSubject
}

// GetFilterSubjectOk returns a tuple with the FilterSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSKVBucketCreateRequestMirror) GetFilterSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.FilterSubject) {
		return nil, false
	}
	return o.FilterSubject, true
}

// HasFilterSubject returns a boolean if a field has been set.
func (o *JSKVBucketCreateRequestMirror) HasFilterSubject() bool {
	if o != nil && !IsNil(o.FilterSubject) {
		return true
	}

	return false
}

// SetFilterSubject gets a reference to the given string and assigns it to the FilterSubject field.
func (o *JSKVBucketCreateRequestMirror) SetFilterSubject(v string) {
	o.FilterSubject = &v
}

// GetName returns the Name field value
func (o *JSKVBucketCreateRequestMirror) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JSKVBucketCreateRequestMirror) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JSKVBucketCreateRequestMirror) SetName(v string) {
	o.Name = v
}

// GetOptStartSeq returns the OptStartSeq field value if set, zero value otherwise.
func (o *JSKVBucketCreateRequestMirror) GetOptStartSeq() int32 {
	if o == nil || IsNil(o.OptStartSeq) {
		var ret int32
		return ret
	}
	return *o.OptStartSeq
}

// GetOptStartSeqOk returns a tuple with the OptStartSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSKVBucketCreateRequestMirror) GetOptStartSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.OptStartSeq) {
		return nil, false
	}
	return o.OptStartSeq, true
}

// HasOptStartSeq returns a boolean if a field has been set.
func (o *JSKVBucketCreateRequestMirror) HasOptStartSeq() bool {
	if o != nil && !IsNil(o.OptStartSeq) {
		return true
	}

	return false
}

// SetOptStartSeq gets a reference to the given int32 and assigns it to the OptStartSeq field.
func (o *JSKVBucketCreateRequestMirror) SetOptStartSeq(v int32) {
	o.OptStartSeq = &v
}

// GetOptStartTime returns the OptStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSKVBucketCreateRequestMirror) GetOptStartTime() string {
	if o == nil || IsNil(o.OptStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.OptStartTime.Get()
}

// GetOptStartTimeOk returns a tuple with the OptStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSKVBucketCreateRequestMirror) GetOptStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptStartTime.Get(), o.OptStartTime.IsSet()
}

// HasOptStartTime returns a boolean if a field has been set.
func (o *JSKVBucketCreateRequestMirror) HasOptStartTime() bool {
	if o != nil && o.OptStartTime.IsSet() {
		return true
	}

	return false
}

// SetOptStartTime gets a reference to the given NullableString and assigns it to the OptStartTime field.
func (o *JSKVBucketCreateRequestMirror) SetOptStartTime(v string) {
	o.OptStartTime.Set(&v)
}

// SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil
func (o *JSKVBucketCreateRequestMirror) SetOptStartTimeNil() {
	o.OptStartTime.Set(nil)
}

// UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil
func (o *JSKVBucketCreateRequestMirror) UnsetOptStartTime() {
	o.OptStartTime.Unset()
}

func (o JSKVBucketCreateRequestMirror) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSKVBucketCreateRequestMirror) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.External.IsSet() {
		toSerialize["external"] = o.External.Get()
	}
	if !IsNil(o.FilterSubject) {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OptStartSeq) {
		toSerialize["opt_start_seq"] = o.OptStartSeq
	}
	if o.OptStartTime.IsSet() {
		toSerialize["opt_start_time"] = o.OptStartTime.Get()
	}
	return toSerialize, nil
}

type NullableJSKVBucketCreateRequestMirror struct {
	value *JSKVBucketCreateRequestMirror
	isSet bool
}

func (v NullableJSKVBucketCreateRequestMirror) Get() *JSKVBucketCreateRequestMirror {
	return v.value
}

func (v *NullableJSKVBucketCreateRequestMirror) Set(val *JSKVBucketCreateRequestMirror) {
	v.value = val
	v.isSet = true
}

func (v NullableJSKVBucketCreateRequestMirror) IsSet() bool {
	return v.isSet
}

func (v *NullableJSKVBucketCreateRequestMirror) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSKVBucketCreateRequestMirror(val *JSKVBucketCreateRequestMirror) *NullableJSKVBucketCreateRequestMirror {
	return &NullableJSKVBucketCreateRequestMirror{value: val, isSet: true}
}

func (v NullableJSKVBucketCreateRequestMirror) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSKVBucketCreateRequestMirror) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
