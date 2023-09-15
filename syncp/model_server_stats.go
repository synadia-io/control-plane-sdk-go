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

// checks if the ServerStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStats{}

// ServerStats ServerStats hold various statistics that we will periodically send out.
type ServerStats struct {
	ActiveAccounts   int32                        `json:"active_accounts"`
	ActiveServers    *int32                       `json:"active_servers,omitempty"`
	Connections      int32                        `json:"connections"`
	Cores            int32                        `json:"cores"`
	Cpu              float64                      `json:"cpu"`
	Gateways         []ServerStatsGatewaysInner   `json:"gateways,omitempty"`
	Jetstream        NullableServerStatsJetstream `json:"jetstream,omitempty"`
	Mem              int64                        `json:"mem"`
	Received         DataStats                    `json:"received"`
	Routes           []ServerStatsRoutesInner     `json:"routes,omitempty"`
	Sent             DataStats                    `json:"sent"`
	SlowConsumers    int64                        `json:"slow_consumers"`
	Start            time.Time                    `json:"start"`
	Subscriptions    int32                        `json:"subscriptions"`
	TotalConnections int32                        `json:"total_connections"`
}

// NewServerStats instantiates a new ServerStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStats(activeAccounts int32, connections int32, cores int32, cpu float64, mem int64, received DataStats, sent DataStats, slowConsumers int64, start time.Time, subscriptions int32, totalConnections int32) *ServerStats {
	this := ServerStats{}
	this.ActiveAccounts = activeAccounts
	this.Connections = connections
	this.Cores = cores
	this.Cpu = cpu
	this.Mem = mem
	this.Received = received
	this.Sent = sent
	this.SlowConsumers = slowConsumers
	this.Start = start
	this.Subscriptions = subscriptions
	this.TotalConnections = totalConnections
	return &this
}

// NewServerStatsWithDefaults instantiates a new ServerStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatsWithDefaults() *ServerStats {
	this := ServerStats{}
	return &this
}

// GetActiveAccounts returns the ActiveAccounts field value
func (o *ServerStats) GetActiveAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveAccounts
}

// GetActiveAccountsOk returns a tuple with the ActiveAccounts field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetActiveAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveAccounts, true
}

// SetActiveAccounts sets field value
func (o *ServerStats) SetActiveAccounts(v int32) {
	o.ActiveAccounts = v
}

// GetActiveServers returns the ActiveServers field value if set, zero value otherwise.
func (o *ServerStats) GetActiveServers() int32 {
	if o == nil || IsNil(o.ActiveServers) {
		var ret int32
		return ret
	}
	return *o.ActiveServers
}

// GetActiveServersOk returns a tuple with the ActiveServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetActiveServersOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveServers) {
		return nil, false
	}
	return o.ActiveServers, true
}

// HasActiveServers returns a boolean if a field has been set.
func (o *ServerStats) HasActiveServers() bool {
	if o != nil && !IsNil(o.ActiveServers) {
		return true
	}

	return false
}

// SetActiveServers gets a reference to the given int32 and assigns it to the ActiveServers field.
func (o *ServerStats) SetActiveServers(v int32) {
	o.ActiveServers = &v
}

// GetConnections returns the Connections field value
func (o *ServerStats) GetConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *ServerStats) SetConnections(v int32) {
	o.Connections = v
}

// GetCores returns the Cores field value
func (o *ServerStats) GetCores() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cores
}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cores, true
}

// SetCores sets field value
func (o *ServerStats) SetCores(v int32) {
	o.Cores = v
}

// GetCpu returns the Cpu field value
func (o *ServerStats) GetCpu() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetCpuOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *ServerStats) SetCpu(v float64) {
	o.Cpu = v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *ServerStats) GetGateways() []ServerStatsGatewaysInner {
	if o == nil || IsNil(o.Gateways) {
		var ret []ServerStatsGatewaysInner
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetGatewaysOk() ([]ServerStatsGatewaysInner, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *ServerStats) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []ServerStatsGatewaysInner and assigns it to the Gateways field.
func (o *ServerStats) SetGateways(v []ServerStatsGatewaysInner) {
	o.Gateways = v
}

// GetJetstream returns the Jetstream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerStats) GetJetstream() ServerStatsJetstream {
	if o == nil || IsNil(o.Jetstream.Get()) {
		var ret ServerStatsJetstream
		return ret
	}
	return *o.Jetstream.Get()
}

// GetJetstreamOk returns a tuple with the Jetstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerStats) GetJetstreamOk() (*ServerStatsJetstream, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jetstream.Get(), o.Jetstream.IsSet()
}

// HasJetstream returns a boolean if a field has been set.
func (o *ServerStats) HasJetstream() bool {
	if o != nil && o.Jetstream.IsSet() {
		return true
	}

	return false
}

// SetJetstream gets a reference to the given NullableServerStatsJetstream and assigns it to the Jetstream field.
func (o *ServerStats) SetJetstream(v ServerStatsJetstream) {
	o.Jetstream.Set(&v)
}

// SetJetstreamNil sets the value for Jetstream to be an explicit nil
func (o *ServerStats) SetJetstreamNil() {
	o.Jetstream.Set(nil)
}

// UnsetJetstream ensures that no value is present for Jetstream, not even an explicit nil
func (o *ServerStats) UnsetJetstream() {
	o.Jetstream.Unset()
}

// GetMem returns the Mem field value
func (o *ServerStats) GetMem() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Mem
}

// GetMemOk returns a tuple with the Mem field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetMemOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mem, true
}

// SetMem sets field value
func (o *ServerStats) SetMem(v int64) {
	o.Mem = v
}

// GetReceived returns the Received field value
func (o *ServerStats) GetReceived() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Received
}

// GetReceivedOk returns a tuple with the Received field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetReceivedOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Received, true
}

// SetReceived sets field value
func (o *ServerStats) SetReceived(v DataStats) {
	o.Received = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *ServerStats) GetRoutes() []ServerStatsRoutesInner {
	if o == nil || IsNil(o.Routes) {
		var ret []ServerStatsRoutesInner
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStats) GetRoutesOk() ([]ServerStatsRoutesInner, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *ServerStats) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []ServerStatsRoutesInner and assigns it to the Routes field.
func (o *ServerStats) SetRoutes(v []ServerStatsRoutesInner) {
	o.Routes = v
}

// GetSent returns the Sent field value
func (o *ServerStats) GetSent() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Sent
}

// GetSentOk returns a tuple with the Sent field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetSentOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sent, true
}

// SetSent sets field value
func (o *ServerStats) SetSent(v DataStats) {
	o.Sent = v
}

// GetSlowConsumers returns the SlowConsumers field value
func (o *ServerStats) GetSlowConsumers() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowConsumers
}

// GetSlowConsumersOk returns a tuple with the SlowConsumers field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetSlowConsumersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowConsumers, true
}

// SetSlowConsumers sets field value
func (o *ServerStats) SetSlowConsumers(v int64) {
	o.SlowConsumers = v
}

// GetStart returns the Start field value
func (o *ServerStats) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ServerStats) SetStart(v time.Time) {
	o.Start = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *ServerStats) GetSubscriptions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetSubscriptionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *ServerStats) SetSubscriptions(v int32) {
	o.Subscriptions = v
}

// GetTotalConnections returns the TotalConnections field value
func (o *ServerStats) GetTotalConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalConnections
}

// GetTotalConnectionsOk returns a tuple with the TotalConnections field value
// and a boolean to check if the value has been set.
func (o *ServerStats) GetTotalConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalConnections, true
}

// SetTotalConnections sets field value
func (o *ServerStats) SetTotalConnections(v int32) {
	o.TotalConnections = v
}

func (o ServerStats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_accounts"] = o.ActiveAccounts
	if !IsNil(o.ActiveServers) {
		toSerialize["active_servers"] = o.ActiveServers
	}
	toSerialize["connections"] = o.Connections
	toSerialize["cores"] = o.Cores
	toSerialize["cpu"] = o.Cpu
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	if o.Jetstream.IsSet() {
		toSerialize["jetstream"] = o.Jetstream.Get()
	}
	toSerialize["mem"] = o.Mem
	toSerialize["received"] = o.Received
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	toSerialize["sent"] = o.Sent
	toSerialize["slow_consumers"] = o.SlowConsumers
	toSerialize["start"] = o.Start
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["total_connections"] = o.TotalConnections
	return toSerialize, nil
}

type NullableServerStats struct {
	value *ServerStats
	isSet bool
}

func (v NullableServerStats) Get() *ServerStats {
	return v.value
}

func (v *NullableServerStats) Set(val *ServerStats) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStats) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStats(val *ServerStats) *NullableServerStats {
	return &NullableServerStats{value: val, isSet: true}
}

func (v NullableServerStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
