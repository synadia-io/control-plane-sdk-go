/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSConsumerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSConsumerInfoResponse{}

// JSConsumerInfoResponse struct for JSConsumerInfoResponse
type JSConsumerInfoResponse struct {
	ConsumerInfo
	ConsumerType JSConsumerType `json:"consumer_type"`
	Id           string         `json:"id"`
}

func (o JSConsumerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consumer_type"] = o.ConsumerType
	toSerialize["id"] = o.Id
	toSerialize["ack_floor"] = o.AckFloor
	if o.Cluster != nil && !o.Cluster.IsNull() {
		toSerialize["cluster"] = o.Cluster.Val
	}
	toSerialize["config"] = o.Config
	toSerialize["created"] = o.Created
	toSerialize["delivered"] = o.Delivered
	toSerialize["name"] = o.Name
	toSerialize["num_ack_pending"] = o.NumAckPending
	toSerialize["num_pending"] = o.NumPending
	toSerialize["num_redelivered"] = o.NumRedelivered
	toSerialize["num_waiting"] = o.NumWaiting
	if o.PushBound != nil {
		toSerialize["push_bound"] = o.PushBound
	}
	toSerialize["stream_name"] = o.StreamName
	toSerialize["ts"] = o.Ts
	return toSerialize, nil
}
