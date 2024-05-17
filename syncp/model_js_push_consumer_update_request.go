/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSPushConsumerUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSPushConsumerUpdateRequest{}

// JSPushConsumerUpdateRequest struct for JSPushConsumerUpdateRequest
type JSPushConsumerUpdateRequest struct {
	AckWait       *int64  `json:"ack_wait,omitempty"`
	Backoff       []int64 `json:"backoff,omitempty"`
	Description   *string `json:"description,omitempty"`
	FilterSubject *string `json:"filter_subject,omitempty"`
	HeadersOnly   *bool   `json:"headers_only,omitempty"`
	MaxAckPending *int64  `json:"max_ack_pending,omitempty"`
	MaxDeliver    *int64  `json:"max_deliver,omitempty"`
	SampleFreq    *string `json:"sample_freq,omitempty"`
}

func (o JSPushConsumerUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AckWait != nil {
		toSerialize["ack_wait"] = o.AckWait
	}
	if len(o.Backoff) != 0 {
		toSerialize["backoff"] = o.Backoff
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FilterSubject != nil {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	if o.HeadersOnly != nil {
		toSerialize["headers_only"] = o.HeadersOnly
	}
	if o.MaxAckPending != nil {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if o.MaxDeliver != nil {
		toSerialize["max_deliver"] = o.MaxDeliver
	}
	if o.SampleFreq != nil {
		toSerialize["sample_freq"] = o.SampleFreq
	}
	return toSerialize, nil
}
