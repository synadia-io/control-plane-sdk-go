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

// checks if the ServerStatsGatewaysInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStatsGatewaysInner{}

// ServerStatsGatewaysInner struct for ServerStatsGatewaysInner
type ServerStatsGatewaysInner struct {
	Gwid               int32     `json:"gwid"`
	InboundConnections int32     `json:"inbound_connections"`
	Name               string    `json:"name"`
	Received           DataStats `json:"received"`
	Sent               DataStats `json:"sent"`
}

// NewServerStatsGatewaysInner instantiates a new ServerStatsGatewaysInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStatsGatewaysInner(gwid int32, inboundConnections int32, name string, received DataStats, sent DataStats) *ServerStatsGatewaysInner {
	this := ServerStatsGatewaysInner{}
	this.Gwid = gwid
	this.InboundConnections = inboundConnections
	this.Name = name
	this.Received = received
	this.Sent = sent
	return &this
}

// NewServerStatsGatewaysInnerWithDefaults instantiates a new ServerStatsGatewaysInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatsGatewaysInnerWithDefaults() *ServerStatsGatewaysInner {
	this := ServerStatsGatewaysInner{}
	return &this
}

// GetGwid returns the Gwid field value
func (o *ServerStatsGatewaysInner) GetGwid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Gwid
}

// GetGwidOk returns a tuple with the Gwid field value
// and a boolean to check if the value has been set.
func (o *ServerStatsGatewaysInner) GetGwidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gwid, true
}

// SetGwid sets field value
func (o *ServerStatsGatewaysInner) SetGwid(v int32) {
	o.Gwid = v
}

// GetInboundConnections returns the InboundConnections field value
func (o *ServerStatsGatewaysInner) GetInboundConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InboundConnections
}

// GetInboundConnectionsOk returns a tuple with the InboundConnections field value
// and a boolean to check if the value has been set.
func (o *ServerStatsGatewaysInner) GetInboundConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InboundConnections, true
}

// SetInboundConnections sets field value
func (o *ServerStatsGatewaysInner) SetInboundConnections(v int32) {
	o.InboundConnections = v
}

// GetName returns the Name field value
func (o *ServerStatsGatewaysInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerStatsGatewaysInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerStatsGatewaysInner) SetName(v string) {
	o.Name = v
}

// GetReceived returns the Received field value
func (o *ServerStatsGatewaysInner) GetReceived() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Received
}

// GetReceivedOk returns a tuple with the Received field value
// and a boolean to check if the value has been set.
func (o *ServerStatsGatewaysInner) GetReceivedOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Received, true
}

// SetReceived sets field value
func (o *ServerStatsGatewaysInner) SetReceived(v DataStats) {
	o.Received = v
}

// GetSent returns the Sent field value
func (o *ServerStatsGatewaysInner) GetSent() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Sent
}

// GetSentOk returns a tuple with the Sent field value
// and a boolean to check if the value has been set.
func (o *ServerStatsGatewaysInner) GetSentOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sent, true
}

// SetSent sets field value
func (o *ServerStatsGatewaysInner) SetSent(v DataStats) {
	o.Sent = v
}

func (o ServerStatsGatewaysInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStatsGatewaysInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gwid"] = o.Gwid
	toSerialize["inbound_connections"] = o.InboundConnections
	toSerialize["name"] = o.Name
	toSerialize["received"] = o.Received
	toSerialize["sent"] = o.Sent
	return toSerialize, nil
}

type NullableServerStatsGatewaysInner struct {
	value *ServerStatsGatewaysInner
	isSet bool
}

func (v NullableServerStatsGatewaysInner) Get() *ServerStatsGatewaysInner {
	return v.value
}

func (v *NullableServerStatsGatewaysInner) Set(val *ServerStatsGatewaysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStatsGatewaysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStatsGatewaysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStatsGatewaysInner(val *ServerStatsGatewaysInner) *NullableServerStatsGatewaysInner {
	return &NullableServerStatsGatewaysInner{value: val, isSet: true}
}

func (v NullableServerStatsGatewaysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStatsGatewaysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
