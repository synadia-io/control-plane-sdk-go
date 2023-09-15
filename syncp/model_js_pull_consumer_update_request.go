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

// checks if the JSPullConsumerUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSPullConsumerUpdateRequest{}

// JSPullConsumerUpdateRequest struct for JSPullConsumerUpdateRequest
type JSPullConsumerUpdateRequest struct {
	AckWait       *int64  `json:"ack_wait,omitempty"`
	Backoff       []int64 `json:"backoff,omitempty"`
	Description   *string `json:"description,omitempty"`
	FilterSubject *string `json:"filter_subject,omitempty"`
	MaxAckPending *int32  `json:"max_ack_pending,omitempty"`
	MaxDeliver    *int32  `json:"max_deliver,omitempty"`
	SampleFreq    *string `json:"sample_freq,omitempty"`
}

// NewJSPullConsumerUpdateRequest instantiates a new JSPullConsumerUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSPullConsumerUpdateRequest() *JSPullConsumerUpdateRequest {
	this := JSPullConsumerUpdateRequest{}
	return &this
}

// NewJSPullConsumerUpdateRequestWithDefaults instantiates a new JSPullConsumerUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSPullConsumerUpdateRequestWithDefaults() *JSPullConsumerUpdateRequest {
	this := JSPullConsumerUpdateRequest{}
	return &this
}

// GetAckWait returns the AckWait field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetAckWait() int64 {
	if o == nil || IsNil(o.AckWait) {
		var ret int64
		return ret
	}
	return *o.AckWait
}

// GetAckWaitOk returns a tuple with the AckWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetAckWaitOk() (*int64, bool) {
	if o == nil || IsNil(o.AckWait) {
		return nil, false
	}
	return o.AckWait, true
}

// HasAckWait returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasAckWait() bool {
	if o != nil && !IsNil(o.AckWait) {
		return true
	}

	return false
}

// SetAckWait gets a reference to the given int64 and assigns it to the AckWait field.
func (o *JSPullConsumerUpdateRequest) SetAckWait(v int64) {
	o.AckWait = &v
}

// GetBackoff returns the Backoff field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetBackoff() []int64 {
	if o == nil || IsNil(o.Backoff) {
		var ret []int64
		return ret
	}
	return o.Backoff
}

// GetBackoffOk returns a tuple with the Backoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetBackoffOk() ([]int64, bool) {
	if o == nil || IsNil(o.Backoff) {
		return nil, false
	}
	return o.Backoff, true
}

// HasBackoff returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasBackoff() bool {
	if o != nil && !IsNil(o.Backoff) {
		return true
	}

	return false
}

// SetBackoff gets a reference to the given []int64 and assigns it to the Backoff field.
func (o *JSPullConsumerUpdateRequest) SetBackoff(v []int64) {
	o.Backoff = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JSPullConsumerUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFilterSubject returns the FilterSubject field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetFilterSubject() string {
	if o == nil || IsNil(o.FilterSubject) {
		var ret string
		return ret
	}
	return *o.FilterSubject
}

// GetFilterSubjectOk returns a tuple with the FilterSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetFilterSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.FilterSubject) {
		return nil, false
	}
	return o.FilterSubject, true
}

// HasFilterSubject returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasFilterSubject() bool {
	if o != nil && !IsNil(o.FilterSubject) {
		return true
	}

	return false
}

// SetFilterSubject gets a reference to the given string and assigns it to the FilterSubject field.
func (o *JSPullConsumerUpdateRequest) SetFilterSubject(v string) {
	o.FilterSubject = &v
}

// GetMaxAckPending returns the MaxAckPending field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetMaxAckPending() int32 {
	if o == nil || IsNil(o.MaxAckPending) {
		var ret int32
		return ret
	}
	return *o.MaxAckPending
}

// GetMaxAckPendingOk returns a tuple with the MaxAckPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetMaxAckPendingOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAckPending) {
		return nil, false
	}
	return o.MaxAckPending, true
}

// HasMaxAckPending returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasMaxAckPending() bool {
	if o != nil && !IsNil(o.MaxAckPending) {
		return true
	}

	return false
}

// SetMaxAckPending gets a reference to the given int32 and assigns it to the MaxAckPending field.
func (o *JSPullConsumerUpdateRequest) SetMaxAckPending(v int32) {
	o.MaxAckPending = &v
}

// GetMaxDeliver returns the MaxDeliver field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetMaxDeliver() int32 {
	if o == nil || IsNil(o.MaxDeliver) {
		var ret int32
		return ret
	}
	return *o.MaxDeliver
}

// GetMaxDeliverOk returns a tuple with the MaxDeliver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetMaxDeliverOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDeliver) {
		return nil, false
	}
	return o.MaxDeliver, true
}

// HasMaxDeliver returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasMaxDeliver() bool {
	if o != nil && !IsNil(o.MaxDeliver) {
		return true
	}

	return false
}

// SetMaxDeliver gets a reference to the given int32 and assigns it to the MaxDeliver field.
func (o *JSPullConsumerUpdateRequest) SetMaxDeliver(v int32) {
	o.MaxDeliver = &v
}

// GetSampleFreq returns the SampleFreq field value if set, zero value otherwise.
func (o *JSPullConsumerUpdateRequest) GetSampleFreq() string {
	if o == nil || IsNil(o.SampleFreq) {
		var ret string
		return ret
	}
	return *o.SampleFreq
}

// GetSampleFreqOk returns a tuple with the SampleFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerUpdateRequest) GetSampleFreqOk() (*string, bool) {
	if o == nil || IsNil(o.SampleFreq) {
		return nil, false
	}
	return o.SampleFreq, true
}

// HasSampleFreq returns a boolean if a field has been set.
func (o *JSPullConsumerUpdateRequest) HasSampleFreq() bool {
	if o != nil && !IsNil(o.SampleFreq) {
		return true
	}

	return false
}

// SetSampleFreq gets a reference to the given string and assigns it to the SampleFreq field.
func (o *JSPullConsumerUpdateRequest) SetSampleFreq(v string) {
	o.SampleFreq = &v
}

func (o JSPullConsumerUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSPullConsumerUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AckWait) {
		toSerialize["ack_wait"] = o.AckWait
	}
	if !IsNil(o.Backoff) {
		toSerialize["backoff"] = o.Backoff
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FilterSubject) {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	if !IsNil(o.MaxAckPending) {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if !IsNil(o.MaxDeliver) {
		toSerialize["max_deliver"] = o.MaxDeliver
	}
	if !IsNil(o.SampleFreq) {
		toSerialize["sample_freq"] = o.SampleFreq
	}
	return toSerialize, nil
}

type NullableJSPullConsumerUpdateRequest struct {
	value *JSPullConsumerUpdateRequest
	isSet bool
}

func (v NullableJSPullConsumerUpdateRequest) Get() *JSPullConsumerUpdateRequest {
	return v.value
}

func (v *NullableJSPullConsumerUpdateRequest) Set(val *JSPullConsumerUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJSPullConsumerUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJSPullConsumerUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSPullConsumerUpdateRequest(val *JSPullConsumerUpdateRequest) *NullableJSPullConsumerUpdateRequest {
	return &NullableJSPullConsumerUpdateRequest{value: val, isSet: true}
}

func (v NullableJSPullConsumerUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSPullConsumerUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}