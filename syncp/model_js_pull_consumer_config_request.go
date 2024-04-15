/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSPullConsumerConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSPullConsumerConfigRequest{}

// JSPullConsumerConfigRequest struct for JSPullConsumerConfigRequest
type JSPullConsumerConfigRequest struct {
	JSCommonConsumerConfigRequest
	MaxBatch   *int32 `json:"max_batch,omitempty"`
	MaxBytes   *int32 `json:"max_bytes,omitempty"`
	MaxExpires *int64 `json:"max_expires,omitempty"`
	// Pull based options.
	MaxWaiting *int32 `json:"max_waiting,omitempty"`
}

func (o JSPullConsumerConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxBatch != nil {
		toSerialize["max_batch"] = o.MaxBatch
	}
	if o.MaxBytes != nil {
		toSerialize["max_bytes"] = o.MaxBytes
	}
	if o.MaxExpires != nil {
		toSerialize["max_expires"] = o.MaxExpires
	}
	if o.MaxWaiting != nil {
		toSerialize["max_waiting"] = o.MaxWaiting
	}
	toSerialize["ack_policy"] = o.AckPolicy
	if o.AckWait != nil {
		toSerialize["ack_wait"] = o.AckWait
	}
	if len(o.Backoff) != 0 {
		toSerialize["backoff"] = o.Backoff
	}
	toSerialize["deliver_policy"] = o.DeliverPolicy
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Direct != nil {
		toSerialize["direct"] = o.Direct
	}
	if o.DurableName != nil {
		toSerialize["durable_name"] = o.DurableName
	}
	if o.FilterSubject != nil {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	if o.InactiveThreshold != nil {
		toSerialize["inactive_threshold"] = o.InactiveThreshold
	}
	if o.MaxAckPending != nil {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if o.MaxDeliver != nil {
		toSerialize["max_deliver"] = o.MaxDeliver
	}
	if o.MemStorage != nil {
		toSerialize["mem_storage"] = o.MemStorage
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["num_replicas"] = o.NumReplicas
	if o.OptStartSeq != nil {
		toSerialize["opt_start_seq"] = o.OptStartSeq
	}
	if o.OptStartTime != nil {
		toSerialize["opt_start_time"] = o.OptStartTime
	}
	toSerialize["replay_policy"] = o.ReplayPolicy
	if o.SampleFreq != nil {
		toSerialize["sample_freq"] = o.SampleFreq
	}
	return toSerialize, nil
}
