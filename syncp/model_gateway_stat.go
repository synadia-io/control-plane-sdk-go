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

// checks if the GatewayStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayStat{}

// GatewayStat GatewayStat holds gateway statistics.
type GatewayStat struct {
	Gwid               int32     `json:"gwid"`
	InboundConnections int32     `json:"inbound_connections"`
	Name               string    `json:"name"`
	Received           DataStats `json:"received"`
	Sent               DataStats `json:"sent"`
}

// NewGatewayStat instantiates a new GatewayStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayStat(gwid int32, inboundConnections int32, name string, received DataStats, sent DataStats) *GatewayStat {
	this := GatewayStat{}
	this.Gwid = gwid
	this.InboundConnections = inboundConnections
	this.Name = name
	this.Received = received
	this.Sent = sent
	return &this
}

// NewGatewayStatWithDefaults instantiates a new GatewayStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayStatWithDefaults() *GatewayStat {
	this := GatewayStat{}
	return &this
}

// GetGwid returns the Gwid field value
func (o *GatewayStat) GetGwid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Gwid
}

// GetGwidOk returns a tuple with the Gwid field value
// and a boolean to check if the value has been set.
func (o *GatewayStat) GetGwidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gwid, true
}

// SetGwid sets field value
func (o *GatewayStat) SetGwid(v int32) {
	o.Gwid = v
}

// GetInboundConnections returns the InboundConnections field value
func (o *GatewayStat) GetInboundConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InboundConnections
}

// GetInboundConnectionsOk returns a tuple with the InboundConnections field value
// and a boolean to check if the value has been set.
func (o *GatewayStat) GetInboundConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InboundConnections, true
}

// SetInboundConnections sets field value
func (o *GatewayStat) SetInboundConnections(v int32) {
	o.InboundConnections = v
}

// GetName returns the Name field value
func (o *GatewayStat) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayStat) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayStat) SetName(v string) {
	o.Name = v
}

// GetReceived returns the Received field value
func (o *GatewayStat) GetReceived() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Received
}

// GetReceivedOk returns a tuple with the Received field value
// and a boolean to check if the value has been set.
func (o *GatewayStat) GetReceivedOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Received, true
}

// SetReceived sets field value
func (o *GatewayStat) SetReceived(v DataStats) {
	o.Received = v
}

// GetSent returns the Sent field value
func (o *GatewayStat) GetSent() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Sent
}

// GetSentOk returns a tuple with the Sent field value
// and a boolean to check if the value has been set.
func (o *GatewayStat) GetSentOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sent, true
}

// SetSent sets field value
func (o *GatewayStat) SetSent(v DataStats) {
	o.Sent = v
}

func (o GatewayStat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gwid"] = o.Gwid
	toSerialize["inbound_connections"] = o.InboundConnections
	toSerialize["name"] = o.Name
	toSerialize["received"] = o.Received
	toSerialize["sent"] = o.Sent
	return toSerialize, nil
}

type NullableGatewayStat struct {
	value *GatewayStat
	isSet bool
}

func (v NullableGatewayStat) Get() *GatewayStat {
	return v.value
}

func (v *NullableGatewayStat) Set(val *GatewayStat) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayStat) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayStat(val *GatewayStat) *NullableGatewayStat {
	return &NullableGatewayStat{value: val, isSet: true}
}

func (v NullableGatewayStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
