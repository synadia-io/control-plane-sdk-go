/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the NatsCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsCluster{}

// NatsCluster struct for NatsCluster
type NatsCluster struct {
	Connections      int32  `json:"connections"`
	IncomingGateways int32  `json:"incomingGateways"`
	Name             string `json:"name"`
	NodeCount        int32  `json:"nodeCount"`
	OutgoingGateways int32  `json:"outgoingGateways"`
}

// NewNatsCluster instantiates a new NatsCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNatsCluster(connections int32, incomingGateways int32, name string, nodeCount int32, outgoingGateways int32) *NatsCluster {
	this := NatsCluster{}
	this.Connections = connections
	this.IncomingGateways = incomingGateways
	this.Name = name
	this.NodeCount = nodeCount
	this.OutgoingGateways = outgoingGateways
	return &this
}

// NewNatsClusterWithDefaults instantiates a new NatsCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNatsClusterWithDefaults() *NatsCluster {
	this := NatsCluster{}
	return &this
}

// GetConnections returns the Connections field value
func (o *NatsCluster) GetConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *NatsCluster) GetConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *NatsCluster) SetConnections(v int32) {
	o.Connections = v
}

// GetIncomingGateways returns the IncomingGateways field value
func (o *NatsCluster) GetIncomingGateways() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IncomingGateways
}

// GetIncomingGatewaysOk returns a tuple with the IncomingGateways field value
// and a boolean to check if the value has been set.
func (o *NatsCluster) GetIncomingGatewaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncomingGateways, true
}

// SetIncomingGateways sets field value
func (o *NatsCluster) SetIncomingGateways(v int32) {
	o.IncomingGateways = v
}

// GetName returns the Name field value
func (o *NatsCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NatsCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NatsCluster) SetName(v string) {
	o.Name = v
}

// GetNodeCount returns the NodeCount field value
func (o *NatsCluster) GetNodeCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
func (o *NatsCluster) GetNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeCount, true
}

// SetNodeCount sets field value
func (o *NatsCluster) SetNodeCount(v int32) {
	o.NodeCount = v
}

// GetOutgoingGateways returns the OutgoingGateways field value
func (o *NatsCluster) GetOutgoingGateways() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OutgoingGateways
}

// GetOutgoingGatewaysOk returns a tuple with the OutgoingGateways field value
// and a boolean to check if the value has been set.
func (o *NatsCluster) GetOutgoingGatewaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutgoingGateways, true
}

// SetOutgoingGateways sets field value
func (o *NatsCluster) SetOutgoingGateways(v int32) {
	o.OutgoingGateways = v
}

func (o NatsCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NatsCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["incomingGateways"] = o.IncomingGateways
	toSerialize["name"] = o.Name
	toSerialize["nodeCount"] = o.NodeCount
	toSerialize["outgoingGateways"] = o.OutgoingGateways
	return toSerialize, nil
}

type NullableNatsCluster struct {
	value *NatsCluster
	isSet bool
}

func (v NullableNatsCluster) Get() *NatsCluster {
	return v.value
}

func (v *NullableNatsCluster) Set(val *NatsCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableNatsCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableNatsCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatsCluster(val *NatsCluster) *NullableNatsCluster {
	return &NullableNatsCluster{value: val, isSet: true}
}

func (v NullableNatsCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatsCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
