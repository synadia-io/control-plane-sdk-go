/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"time"
)

// checks if the JSPullConsumerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSPullConsumerInfoResponse{}

// JSPullConsumerInfoResponse struct for JSPullConsumerInfoResponse
type JSPullConsumerInfoResponse struct {
	AckFloor       SequenceInfo                 `json:"ack_floor"`
	Cluster        NullableConsumerInfoCluster  `json:"cluster,omitempty"`
	Config         *JSPullConsumerConfigRequest `json:"config,omitempty"`
	Created        time.Time                    `json:"created"`
	Delivered      SequenceInfo                 `json:"delivered"`
	Id             string                       `json:"id"`
	Name           string                       `json:"name"`
	NumAckPending  int32                        `json:"num_ack_pending"`
	NumPending     int32                        `json:"num_pending"`
	NumRedelivered int32                        `json:"num_redelivered"`
	NumWaiting     int32                        `json:"num_waiting"`
	StreamName     string                       `json:"stream_name"`
}

// NewJSPullConsumerInfoResponse instantiates a new JSPullConsumerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSPullConsumerInfoResponse(ackFloor SequenceInfo, created time.Time, delivered SequenceInfo, id string, name string, numAckPending int32, numPending int32, numRedelivered int32, numWaiting int32, streamName string) *JSPullConsumerInfoResponse {
	this := JSPullConsumerInfoResponse{}
	this.AckFloor = ackFloor
	this.Created = created
	this.Delivered = delivered
	this.Id = id
	this.Name = name
	this.NumAckPending = numAckPending
	this.NumPending = numPending
	this.NumRedelivered = numRedelivered
	this.NumWaiting = numWaiting
	this.StreamName = streamName
	return &this
}

// NewJSPullConsumerInfoResponseWithDefaults instantiates a new JSPullConsumerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSPullConsumerInfoResponseWithDefaults() *JSPullConsumerInfoResponse {
	this := JSPullConsumerInfoResponse{}
	return &this
}

// GetAckFloor returns the AckFloor field value
func (o *JSPullConsumerInfoResponse) GetAckFloor() SequenceInfo {
	if o == nil {
		var ret SequenceInfo
		return ret
	}

	return o.AckFloor
}

// GetAckFloorOk returns a tuple with the AckFloor field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetAckFloorOk() (*SequenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckFloor, true
}

// SetAckFloor sets field value
func (o *JSPullConsumerInfoResponse) SetAckFloor(v SequenceInfo) {
	o.AckFloor = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSPullConsumerInfoResponse) GetCluster() ConsumerInfoCluster {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret ConsumerInfoCluster
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSPullConsumerInfoResponse) GetClusterOk() (*ConsumerInfoCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *JSPullConsumerInfoResponse) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableConsumerInfoCluster and assigns it to the Cluster field.
func (o *JSPullConsumerInfoResponse) SetCluster(v ConsumerInfoCluster) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *JSPullConsumerInfoResponse) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *JSPullConsumerInfoResponse) UnsetCluster() {
	o.Cluster.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *JSPullConsumerInfoResponse) GetConfig() JSPullConsumerConfigRequest {
	if o == nil || IsNil(o.Config) {
		var ret JSPullConsumerConfigRequest
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetConfigOk() (*JSPullConsumerConfigRequest, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *JSPullConsumerInfoResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given JSPullConsumerConfigRequest and assigns it to the Config field.
func (o *JSPullConsumerInfoResponse) SetConfig(v JSPullConsumerConfigRequest) {
	o.Config = &v
}

// GetCreated returns the Created field value
func (o *JSPullConsumerInfoResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *JSPullConsumerInfoResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDelivered returns the Delivered field value
func (o *JSPullConsumerInfoResponse) GetDelivered() SequenceInfo {
	if o == nil {
		var ret SequenceInfo
		return ret
	}

	return o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetDeliveredOk() (*SequenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivered, true
}

// SetDelivered sets field value
func (o *JSPullConsumerInfoResponse) SetDelivered(v SequenceInfo) {
	o.Delivered = v
}

// GetId returns the Id field value
func (o *JSPullConsumerInfoResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JSPullConsumerInfoResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *JSPullConsumerInfoResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JSPullConsumerInfoResponse) SetName(v string) {
	o.Name = v
}

// GetNumAckPending returns the NumAckPending field value
func (o *JSPullConsumerInfoResponse) GetNumAckPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumAckPending
}

// GetNumAckPendingOk returns a tuple with the NumAckPending field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetNumAckPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAckPending, true
}

// SetNumAckPending sets field value
func (o *JSPullConsumerInfoResponse) SetNumAckPending(v int32) {
	o.NumAckPending = v
}

// GetNumPending returns the NumPending field value
func (o *JSPullConsumerInfoResponse) GetNumPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPending
}

// GetNumPendingOk returns a tuple with the NumPending field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetNumPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPending, true
}

// SetNumPending sets field value
func (o *JSPullConsumerInfoResponse) SetNumPending(v int32) {
	o.NumPending = v
}

// GetNumRedelivered returns the NumRedelivered field value
func (o *JSPullConsumerInfoResponse) GetNumRedelivered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumRedelivered
}

// GetNumRedeliveredOk returns a tuple with the NumRedelivered field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetNumRedeliveredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRedelivered, true
}

// SetNumRedelivered sets field value
func (o *JSPullConsumerInfoResponse) SetNumRedelivered(v int32) {
	o.NumRedelivered = v
}

// GetNumWaiting returns the NumWaiting field value
func (o *JSPullConsumerInfoResponse) GetNumWaiting() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumWaiting
}

// GetNumWaitingOk returns a tuple with the NumWaiting field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetNumWaitingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWaiting, true
}

// SetNumWaiting sets field value
func (o *JSPullConsumerInfoResponse) SetNumWaiting(v int32) {
	o.NumWaiting = v
}

// GetStreamName returns the StreamName field value
func (o *JSPullConsumerInfoResponse) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *JSPullConsumerInfoResponse) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *JSPullConsumerInfoResponse) SetStreamName(v string) {
	o.StreamName = v
}

func (o JSPullConsumerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSPullConsumerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ack_floor"] = o.AckFloor
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	toSerialize["created"] = o.Created
	toSerialize["delivered"] = o.Delivered
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["num_ack_pending"] = o.NumAckPending
	toSerialize["num_pending"] = o.NumPending
	toSerialize["num_redelivered"] = o.NumRedelivered
	toSerialize["num_waiting"] = o.NumWaiting
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}

type NullableJSPullConsumerInfoResponse struct {
	value *JSPullConsumerInfoResponse
	isSet bool
}

func (v NullableJSPullConsumerInfoResponse) Get() *JSPullConsumerInfoResponse {
	return v.value
}

func (v *NullableJSPullConsumerInfoResponse) Set(val *JSPullConsumerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJSPullConsumerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJSPullConsumerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSPullConsumerInfoResponse(val *JSPullConsumerInfoResponse) *NullableJSPullConsumerInfoResponse {
	return &NullableJSPullConsumerInfoResponse{value: val, isSet: true}
}

func (v NullableJSPullConsumerInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSPullConsumerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
