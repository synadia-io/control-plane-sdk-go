/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"time"
)

// checks if the JSPushConsumerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSPushConsumerInfoResponse{}

// JSPushConsumerInfoResponse struct for JSPushConsumerInfoResponse
type JSPushConsumerInfoResponse struct {
	AckFloor       SequenceInfo                 `json:"ack_floor"`
	Cluster        NullableConsumerInfoCluster  `json:"cluster,omitempty"`
	Config         *JSPushConsumerConfigRequest `json:"config,omitempty"`
	Created        time.Time                    `json:"created"`
	Delivered      SequenceInfo                 `json:"delivered"`
	Id             string                       `json:"id"`
	Name           string                       `json:"name"`
	NumAckPending  int32                        `json:"num_ack_pending"`
	NumPending     int32                        `json:"num_pending"`
	NumRedelivered int32                        `json:"num_redelivered"`
	PushBound      *bool                        `json:"push_bound,omitempty"`
	StreamName     string                       `json:"stream_name"`
}

// NewJSPushConsumerInfoResponse instantiates a new JSPushConsumerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSPushConsumerInfoResponse(ackFloor SequenceInfo, created time.Time, delivered SequenceInfo, id string, name string, numAckPending int32, numPending int32, numRedelivered int32, streamName string) *JSPushConsumerInfoResponse {
	this := JSPushConsumerInfoResponse{}
	this.AckFloor = ackFloor
	this.Created = created
	this.Delivered = delivered
	this.Id = id
	this.Name = name
	this.NumAckPending = numAckPending
	this.NumPending = numPending
	this.NumRedelivered = numRedelivered
	this.StreamName = streamName
	return &this
}

// NewJSPushConsumerInfoResponseWithDefaults instantiates a new JSPushConsumerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSPushConsumerInfoResponseWithDefaults() *JSPushConsumerInfoResponse {
	this := JSPushConsumerInfoResponse{}
	return &this
}

// GetAckFloor returns the AckFloor field value
func (o *JSPushConsumerInfoResponse) GetAckFloor() SequenceInfo {
	if o == nil {
		var ret SequenceInfo
		return ret
	}

	return o.AckFloor
}

// GetAckFloorOk returns a tuple with the AckFloor field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetAckFloorOk() (*SequenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckFloor, true
}

// SetAckFloor sets field value
func (o *JSPushConsumerInfoResponse) SetAckFloor(v SequenceInfo) {
	o.AckFloor = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSPushConsumerInfoResponse) GetCluster() ConsumerInfoCluster {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret ConsumerInfoCluster
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSPushConsumerInfoResponse) GetClusterOk() (*ConsumerInfoCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *JSPushConsumerInfoResponse) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableConsumerInfoCluster and assigns it to the Cluster field.
func (o *JSPushConsumerInfoResponse) SetCluster(v ConsumerInfoCluster) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *JSPushConsumerInfoResponse) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *JSPushConsumerInfoResponse) UnsetCluster() {
	o.Cluster.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *JSPushConsumerInfoResponse) GetConfig() JSPushConsumerConfigRequest {
	if o == nil || IsNil(o.Config) {
		var ret JSPushConsumerConfigRequest
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetConfigOk() (*JSPushConsumerConfigRequest, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *JSPushConsumerInfoResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given JSPushConsumerConfigRequest and assigns it to the Config field.
func (o *JSPushConsumerInfoResponse) SetConfig(v JSPushConsumerConfigRequest) {
	o.Config = &v
}

// GetCreated returns the Created field value
func (o *JSPushConsumerInfoResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *JSPushConsumerInfoResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDelivered returns the Delivered field value
func (o *JSPushConsumerInfoResponse) GetDelivered() SequenceInfo {
	if o == nil {
		var ret SequenceInfo
		return ret
	}

	return o.Delivered
}

// GetDeliveredOk returns a tuple with the Delivered field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetDeliveredOk() (*SequenceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivered, true
}

// SetDelivered sets field value
func (o *JSPushConsumerInfoResponse) SetDelivered(v SequenceInfo) {
	o.Delivered = v
}

// GetId returns the Id field value
func (o *JSPushConsumerInfoResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JSPushConsumerInfoResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *JSPushConsumerInfoResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JSPushConsumerInfoResponse) SetName(v string) {
	o.Name = v
}

// GetNumAckPending returns the NumAckPending field value
func (o *JSPushConsumerInfoResponse) GetNumAckPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumAckPending
}

// GetNumAckPendingOk returns a tuple with the NumAckPending field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetNumAckPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAckPending, true
}

// SetNumAckPending sets field value
func (o *JSPushConsumerInfoResponse) SetNumAckPending(v int32) {
	o.NumAckPending = v
}

// GetNumPending returns the NumPending field value
func (o *JSPushConsumerInfoResponse) GetNumPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPending
}

// GetNumPendingOk returns a tuple with the NumPending field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetNumPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPending, true
}

// SetNumPending sets field value
func (o *JSPushConsumerInfoResponse) SetNumPending(v int32) {
	o.NumPending = v
}

// GetNumRedelivered returns the NumRedelivered field value
func (o *JSPushConsumerInfoResponse) GetNumRedelivered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumRedelivered
}

// GetNumRedeliveredOk returns a tuple with the NumRedelivered field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetNumRedeliveredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumRedelivered, true
}

// SetNumRedelivered sets field value
func (o *JSPushConsumerInfoResponse) SetNumRedelivered(v int32) {
	o.NumRedelivered = v
}

// GetPushBound returns the PushBound field value if set, zero value otherwise.
func (o *JSPushConsumerInfoResponse) GetPushBound() bool {
	if o == nil || IsNil(o.PushBound) {
		var ret bool
		return ret
	}
	return *o.PushBound
}

// GetPushBoundOk returns a tuple with the PushBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetPushBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.PushBound) {
		return nil, false
	}
	return o.PushBound, true
}

// HasPushBound returns a boolean if a field has been set.
func (o *JSPushConsumerInfoResponse) HasPushBound() bool {
	if o != nil && !IsNil(o.PushBound) {
		return true
	}

	return false
}

// SetPushBound gets a reference to the given bool and assigns it to the PushBound field.
func (o *JSPushConsumerInfoResponse) SetPushBound(v bool) {
	o.PushBound = &v
}

// GetStreamName returns the StreamName field value
func (o *JSPushConsumerInfoResponse) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *JSPushConsumerInfoResponse) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *JSPushConsumerInfoResponse) SetStreamName(v string) {
	o.StreamName = v
}

func (o JSPushConsumerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSPushConsumerInfoResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PushBound) {
		toSerialize["push_bound"] = o.PushBound
	}
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}

type NullableJSPushConsumerInfoResponse struct {
	value *JSPushConsumerInfoResponse
	isSet bool
}

func (v NullableJSPushConsumerInfoResponse) Get() *JSPushConsumerInfoResponse {
	return v.value
}

func (v *NullableJSPushConsumerInfoResponse) Set(val *JSPushConsumerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJSPushConsumerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJSPushConsumerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSPushConsumerInfoResponse(val *JSPushConsumerInfoResponse) *NullableJSPushConsumerInfoResponse {
	return &NullableJSPushConsumerInfoResponse{value: val, isSet: true}
}

func (v NullableJSPushConsumerInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSPushConsumerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
