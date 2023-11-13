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

// checks if the JSPullConsumerConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSPullConsumerConfigRequest{}

// JSPullConsumerConfigRequest struct for JSPullConsumerConfigRequest
type JSPullConsumerConfigRequest struct {
	AckPolicy         AckPolicy      `json:"ack_policy"`
	AckWait           *int64         `json:"ack_wait,omitempty"`
	Backoff           []int64        `json:"backoff,omitempty"`
	DeliverPolicy     DeliverPolicy  `json:"deliver_policy"`
	Description       *string        `json:"description,omitempty"`
	Direct            *bool          `json:"direct,omitempty"`
	DurableName       *string        `json:"durable_name,omitempty"`
	FilterSubject     *string        `json:"filter_subject,omitempty"`
	InactiveThreshold *int64         `json:"inactive_threshold,omitempty"`
	MaxAckPending     *int32         `json:"max_ack_pending,omitempty"`
	MaxDeliver        *int32         `json:"max_deliver,omitempty"`
	MemStorage        *bool          `json:"mem_storage,omitempty"`
	Name              *string        `json:"name,omitempty"`
	NumReplicas       int32          `json:"num_replicas"`
	OptStartSeq       *int32         `json:"opt_start_seq,omitempty"`
	OptStartTime      NullableString `json:"opt_start_time,omitempty"`
	ReplayPolicy      ReplayPolicy   `json:"replay_policy"`
	SampleFreq        *string        `json:"sample_freq,omitempty"`
	MaxBatch          *int32         `json:"max_batch,omitempty"`
	MaxBytes          *int32         `json:"max_bytes,omitempty"`
	MaxExpires        *int64         `json:"max_expires,omitempty"`
	// Pull based options.
	MaxWaiting *int32 `json:"max_waiting,omitempty"`
}

type _JSPullConsumerConfigRequest JSPullConsumerConfigRequest

// NewJSPullConsumerConfigRequest instantiates a new JSPullConsumerConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSPullConsumerConfigRequest(ackPolicy AckPolicy, deliverPolicy DeliverPolicy, numReplicas int32, replayPolicy ReplayPolicy) *JSPullConsumerConfigRequest {
	this := JSPullConsumerConfigRequest{}
	this.AckPolicy = ackPolicy
	this.DeliverPolicy = deliverPolicy
	this.NumReplicas = numReplicas
	this.ReplayPolicy = replayPolicy
	return &this
}

// NewJSPullConsumerConfigRequestWithDefaults instantiates a new JSPullConsumerConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSPullConsumerConfigRequestWithDefaults() *JSPullConsumerConfigRequest {
	this := JSPullConsumerConfigRequest{}
	return &this
}

// GetAckPolicy returns the AckPolicy field value
func (o *JSPullConsumerConfigRequest) GetAckPolicy() AckPolicy {
	if o == nil {
		var ret AckPolicy
		return ret
	}

	return o.AckPolicy
}

// GetAckPolicyOk returns a tuple with the AckPolicy field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetAckPolicyOk() (*AckPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckPolicy, true
}

// SetAckPolicy sets field value
func (o *JSPullConsumerConfigRequest) SetAckPolicy(v AckPolicy) {
	o.AckPolicy = v
}

// GetAckWait returns the AckWait field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetAckWait() int64 {
	if o == nil || IsNil(o.AckWait) {
		var ret int64
		return ret
	}
	return *o.AckWait
}

// GetAckWaitOk returns a tuple with the AckWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetAckWaitOk() (*int64, bool) {
	if o == nil || IsNil(o.AckWait) {
		return nil, false
	}
	return o.AckWait, true
}

// HasAckWait returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasAckWait() bool {
	if o != nil && !IsNil(o.AckWait) {
		return true
	}

	return false
}

// SetAckWait gets a reference to the given int64 and assigns it to the AckWait field.
func (o *JSPullConsumerConfigRequest) SetAckWait(v int64) {
	o.AckWait = &v
}

// GetBackoff returns the Backoff field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetBackoff() []int64 {
	if o == nil || IsNil(o.Backoff) {
		var ret []int64
		return ret
	}
	return o.Backoff
}

// GetBackoffOk returns a tuple with the Backoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetBackoffOk() ([]int64, bool) {
	if o == nil || IsNil(o.Backoff) {
		return nil, false
	}
	return o.Backoff, true
}

// HasBackoff returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasBackoff() bool {
	if o != nil && !IsNil(o.Backoff) {
		return true
	}

	return false
}

// SetBackoff gets a reference to the given []int64 and assigns it to the Backoff field.
func (o *JSPullConsumerConfigRequest) SetBackoff(v []int64) {
	o.Backoff = v
}

// GetDeliverPolicy returns the DeliverPolicy field value
func (o *JSPullConsumerConfigRequest) GetDeliverPolicy() DeliverPolicy {
	if o == nil {
		var ret DeliverPolicy
		return ret
	}

	return o.DeliverPolicy
}

// GetDeliverPolicyOk returns a tuple with the DeliverPolicy field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetDeliverPolicyOk() (*DeliverPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliverPolicy, true
}

// SetDeliverPolicy sets field value
func (o *JSPullConsumerConfigRequest) SetDeliverPolicy(v DeliverPolicy) {
	o.DeliverPolicy = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JSPullConsumerConfigRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDirect returns the Direct field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetDirect() bool {
	if o == nil || IsNil(o.Direct) {
		var ret bool
		return ret
	}
	return *o.Direct
}

// GetDirectOk returns a tuple with the Direct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.Direct) {
		return nil, false
	}
	return o.Direct, true
}

// HasDirect returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasDirect() bool {
	if o != nil && !IsNil(o.Direct) {
		return true
	}

	return false
}

// SetDirect gets a reference to the given bool and assigns it to the Direct field.
func (o *JSPullConsumerConfigRequest) SetDirect(v bool) {
	o.Direct = &v
}

// GetDurableName returns the DurableName field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetDurableName() string {
	if o == nil || IsNil(o.DurableName) {
		var ret string
		return ret
	}
	return *o.DurableName
}

// GetDurableNameOk returns a tuple with the DurableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetDurableNameOk() (*string, bool) {
	if o == nil || IsNil(o.DurableName) {
		return nil, false
	}
	return o.DurableName, true
}

// HasDurableName returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasDurableName() bool {
	if o != nil && !IsNil(o.DurableName) {
		return true
	}

	return false
}

// SetDurableName gets a reference to the given string and assigns it to the DurableName field.
func (o *JSPullConsumerConfigRequest) SetDurableName(v string) {
	o.DurableName = &v
}

// GetFilterSubject returns the FilterSubject field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetFilterSubject() string {
	if o == nil || IsNil(o.FilterSubject) {
		var ret string
		return ret
	}
	return *o.FilterSubject
}

// GetFilterSubjectOk returns a tuple with the FilterSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetFilterSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.FilterSubject) {
		return nil, false
	}
	return o.FilterSubject, true
}

// HasFilterSubject returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasFilterSubject() bool {
	if o != nil && !IsNil(o.FilterSubject) {
		return true
	}

	return false
}

// SetFilterSubject gets a reference to the given string and assigns it to the FilterSubject field.
func (o *JSPullConsumerConfigRequest) SetFilterSubject(v string) {
	o.FilterSubject = &v
}

// GetInactiveThreshold returns the InactiveThreshold field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetInactiveThreshold() int64 {
	if o == nil || IsNil(o.InactiveThreshold) {
		var ret int64
		return ret
	}
	return *o.InactiveThreshold
}

// GetInactiveThresholdOk returns a tuple with the InactiveThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetInactiveThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.InactiveThreshold) {
		return nil, false
	}
	return o.InactiveThreshold, true
}

// HasInactiveThreshold returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasInactiveThreshold() bool {
	if o != nil && !IsNil(o.InactiveThreshold) {
		return true
	}

	return false
}

// SetInactiveThreshold gets a reference to the given int64 and assigns it to the InactiveThreshold field.
func (o *JSPullConsumerConfigRequest) SetInactiveThreshold(v int64) {
	o.InactiveThreshold = &v
}

// GetMaxAckPending returns the MaxAckPending field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMaxAckPending() int32 {
	if o == nil || IsNil(o.MaxAckPending) {
		var ret int32
		return ret
	}
	return *o.MaxAckPending
}

// GetMaxAckPendingOk returns a tuple with the MaxAckPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMaxAckPendingOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAckPending) {
		return nil, false
	}
	return o.MaxAckPending, true
}

// HasMaxAckPending returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMaxAckPending() bool {
	if o != nil && !IsNil(o.MaxAckPending) {
		return true
	}

	return false
}

// SetMaxAckPending gets a reference to the given int32 and assigns it to the MaxAckPending field.
func (o *JSPullConsumerConfigRequest) SetMaxAckPending(v int32) {
	o.MaxAckPending = &v
}

// GetMaxDeliver returns the MaxDeliver field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMaxDeliver() int32 {
	if o == nil || IsNil(o.MaxDeliver) {
		var ret int32
		return ret
	}
	return *o.MaxDeliver
}

// GetMaxDeliverOk returns a tuple with the MaxDeliver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMaxDeliverOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDeliver) {
		return nil, false
	}
	return o.MaxDeliver, true
}

// HasMaxDeliver returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMaxDeliver() bool {
	if o != nil && !IsNil(o.MaxDeliver) {
		return true
	}

	return false
}

// SetMaxDeliver gets a reference to the given int32 and assigns it to the MaxDeliver field.
func (o *JSPullConsumerConfigRequest) SetMaxDeliver(v int32) {
	o.MaxDeliver = &v
}

// GetMemStorage returns the MemStorage field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMemStorage() bool {
	if o == nil || IsNil(o.MemStorage) {
		var ret bool
		return ret
	}
	return *o.MemStorage
}

// GetMemStorageOk returns a tuple with the MemStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMemStorageOk() (*bool, bool) {
	if o == nil || IsNil(o.MemStorage) {
		return nil, false
	}
	return o.MemStorage, true
}

// HasMemStorage returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMemStorage() bool {
	if o != nil && !IsNil(o.MemStorage) {
		return true
	}

	return false
}

// SetMemStorage gets a reference to the given bool and assigns it to the MemStorage field.
func (o *JSPullConsumerConfigRequest) SetMemStorage(v bool) {
	o.MemStorage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JSPullConsumerConfigRequest) SetName(v string) {
	o.Name = &v
}

// GetNumReplicas returns the NumReplicas field value
func (o *JSPullConsumerConfigRequest) GetNumReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumReplicas
}

// GetNumReplicasOk returns a tuple with the NumReplicas field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetNumReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumReplicas, true
}

// SetNumReplicas sets field value
func (o *JSPullConsumerConfigRequest) SetNumReplicas(v int32) {
	o.NumReplicas = v
}

// GetOptStartSeq returns the OptStartSeq field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetOptStartSeq() int32 {
	if o == nil || IsNil(o.OptStartSeq) {
		var ret int32
		return ret
	}
	return *o.OptStartSeq
}

// GetOptStartSeqOk returns a tuple with the OptStartSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetOptStartSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.OptStartSeq) {
		return nil, false
	}
	return o.OptStartSeq, true
}

// HasOptStartSeq returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasOptStartSeq() bool {
	if o != nil && !IsNil(o.OptStartSeq) {
		return true
	}

	return false
}

// SetOptStartSeq gets a reference to the given int32 and assigns it to the OptStartSeq field.
func (o *JSPullConsumerConfigRequest) SetOptStartSeq(v int32) {
	o.OptStartSeq = &v
}

// GetOptStartTime returns the OptStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSPullConsumerConfigRequest) GetOptStartTime() string {
	if o == nil || IsNil(o.OptStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.OptStartTime.Get()
}

// GetOptStartTimeOk returns a tuple with the OptStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSPullConsumerConfigRequest) GetOptStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptStartTime.Get(), o.OptStartTime.IsSet()
}

// HasOptStartTime returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasOptStartTime() bool {
	if o != nil && o.OptStartTime.IsSet() {
		return true
	}

	return false
}

// SetOptStartTime gets a reference to the given NullableString and assigns it to the OptStartTime field.
func (o *JSPullConsumerConfigRequest) SetOptStartTime(v string) {
	o.OptStartTime.Set(&v)
}

// SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil
func (o *JSPullConsumerConfigRequest) SetOptStartTimeNil() {
	o.OptStartTime.Set(nil)
}

// UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil
func (o *JSPullConsumerConfigRequest) UnsetOptStartTime() {
	o.OptStartTime.Unset()
}

// GetReplayPolicy returns the ReplayPolicy field value
func (o *JSPullConsumerConfigRequest) GetReplayPolicy() ReplayPolicy {
	if o == nil {
		var ret ReplayPolicy
		return ret
	}

	return o.ReplayPolicy
}

// GetReplayPolicyOk returns a tuple with the ReplayPolicy field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetReplayPolicyOk() (*ReplayPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplayPolicy, true
}

// SetReplayPolicy sets field value
func (o *JSPullConsumerConfigRequest) SetReplayPolicy(v ReplayPolicy) {
	o.ReplayPolicy = v
}

// GetSampleFreq returns the SampleFreq field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetSampleFreq() string {
	if o == nil || IsNil(o.SampleFreq) {
		var ret string
		return ret
	}
	return *o.SampleFreq
}

// GetSampleFreqOk returns a tuple with the SampleFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetSampleFreqOk() (*string, bool) {
	if o == nil || IsNil(o.SampleFreq) {
		return nil, false
	}
	return o.SampleFreq, true
}

// HasSampleFreq returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasSampleFreq() bool {
	if o != nil && !IsNil(o.SampleFreq) {
		return true
	}

	return false
}

// SetSampleFreq gets a reference to the given string and assigns it to the SampleFreq field.
func (o *JSPullConsumerConfigRequest) SetSampleFreq(v string) {
	o.SampleFreq = &v
}

// GetMaxBatch returns the MaxBatch field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMaxBatch() int32 {
	if o == nil || IsNil(o.MaxBatch) {
		var ret int32
		return ret
	}
	return *o.MaxBatch
}

// GetMaxBatchOk returns a tuple with the MaxBatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMaxBatchOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxBatch) {
		return nil, false
	}
	return o.MaxBatch, true
}

// HasMaxBatch returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMaxBatch() bool {
	if o != nil && !IsNil(o.MaxBatch) {
		return true
	}

	return false
}

// SetMaxBatch gets a reference to the given int32 and assigns it to the MaxBatch field.
func (o *JSPullConsumerConfigRequest) SetMaxBatch(v int32) {
	o.MaxBatch = &v
}

// GetMaxBytes returns the MaxBytes field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMaxBytes() int32 {
	if o == nil || IsNil(o.MaxBytes) {
		var ret int32
		return ret
	}
	return *o.MaxBytes
}

// GetMaxBytesOk returns a tuple with the MaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMaxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxBytes) {
		return nil, false
	}
	return o.MaxBytes, true
}

// HasMaxBytes returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMaxBytes() bool {
	if o != nil && !IsNil(o.MaxBytes) {
		return true
	}

	return false
}

// SetMaxBytes gets a reference to the given int32 and assigns it to the MaxBytes field.
func (o *JSPullConsumerConfigRequest) SetMaxBytes(v int32) {
	o.MaxBytes = &v
}

// GetMaxExpires returns the MaxExpires field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMaxExpires() int64 {
	if o == nil || IsNil(o.MaxExpires) {
		var ret int64
		return ret
	}
	return *o.MaxExpires
}

// GetMaxExpiresOk returns a tuple with the MaxExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMaxExpiresOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxExpires) {
		return nil, false
	}
	return o.MaxExpires, true
}

// HasMaxExpires returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMaxExpires() bool {
	if o != nil && !IsNil(o.MaxExpires) {
		return true
	}

	return false
}

// SetMaxExpires gets a reference to the given int64 and assigns it to the MaxExpires field.
func (o *JSPullConsumerConfigRequest) SetMaxExpires(v int64) {
	o.MaxExpires = &v
}

// GetMaxWaiting returns the MaxWaiting field value if set, zero value otherwise.
func (o *JSPullConsumerConfigRequest) GetMaxWaiting() int32 {
	if o == nil || IsNil(o.MaxWaiting) {
		var ret int32
		return ret
	}
	return *o.MaxWaiting
}

// GetMaxWaitingOk returns a tuple with the MaxWaiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerConfigRequest) GetMaxWaitingOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxWaiting) {
		return nil, false
	}
	return o.MaxWaiting, true
}

// HasMaxWaiting returns a boolean if a field has been set.
func (o *JSPullConsumerConfigRequest) HasMaxWaiting() bool {
	if o != nil && !IsNil(o.MaxWaiting) {
		return true
	}

	return false
}

// SetMaxWaiting gets a reference to the given int32 and assigns it to the MaxWaiting field.
func (o *JSPullConsumerConfigRequest) SetMaxWaiting(v int32) {
	o.MaxWaiting = &v
}

func (o JSPullConsumerConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSPullConsumerConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ack_policy"] = o.AckPolicy
	if !IsNil(o.AckWait) {
		toSerialize["ack_wait"] = o.AckWait
	}
	if !IsNil(o.Backoff) {
		toSerialize["backoff"] = o.Backoff
	}
	toSerialize["deliver_policy"] = o.DeliverPolicy
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Direct) {
		toSerialize["direct"] = o.Direct
	}
	if !IsNil(o.DurableName) {
		toSerialize["durable_name"] = o.DurableName
	}
	if !IsNil(o.FilterSubject) {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	if !IsNil(o.InactiveThreshold) {
		toSerialize["inactive_threshold"] = o.InactiveThreshold
	}
	if !IsNil(o.MaxAckPending) {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if !IsNil(o.MaxDeliver) {
		toSerialize["max_deliver"] = o.MaxDeliver
	}
	if !IsNil(o.MemStorage) {
		toSerialize["mem_storage"] = o.MemStorage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["num_replicas"] = o.NumReplicas
	if !IsNil(o.OptStartSeq) {
		toSerialize["opt_start_seq"] = o.OptStartSeq
	}
	if o.OptStartTime.IsSet() {
		toSerialize["opt_start_time"] = o.OptStartTime.Get()
	}
	toSerialize["replay_policy"] = o.ReplayPolicy
	if !IsNil(o.SampleFreq) {
		toSerialize["sample_freq"] = o.SampleFreq
	}
	if !IsNil(o.MaxBatch) {
		toSerialize["max_batch"] = o.MaxBatch
	}
	if !IsNil(o.MaxBytes) {
		toSerialize["max_bytes"] = o.MaxBytes
	}
	if !IsNil(o.MaxExpires) {
		toSerialize["max_expires"] = o.MaxExpires
	}
	if !IsNil(o.MaxWaiting) {
		toSerialize["max_waiting"] = o.MaxWaiting
	}
	return toSerialize, nil
}

func (o *JSPullConsumerConfigRequest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ack_policy",
		"deliver_policy",
		"num_replicas",
		"replay_policy",
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

	varJSPullConsumerConfigRequest := _JSPullConsumerConfigRequest{}

	err = json.Unmarshal(bytes, &varJSPullConsumerConfigRequest)

	if err != nil {
		return err
	}

	*o = JSPullConsumerConfigRequest(varJSPullConsumerConfigRequest)

	return err
}

type NullableJSPullConsumerConfigRequest struct {
	value *JSPullConsumerConfigRequest
	isSet bool
}

func (v NullableJSPullConsumerConfigRequest) Get() *JSPullConsumerConfigRequest {
	return v.value
}

func (v *NullableJSPullConsumerConfigRequest) Set(val *JSPullConsumerConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJSPullConsumerConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJSPullConsumerConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSPullConsumerConfigRequest(val *JSPullConsumerConfigRequest) *NullableJSPullConsumerConfigRequest {
	return &NullableJSPullConsumerConfigRequest{value: val, isSet: true}
}

func (v NullableJSPullConsumerConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSPullConsumerConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
