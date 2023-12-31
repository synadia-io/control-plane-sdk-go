/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"time"
)

// checks if the ServerStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerStats{}

// ServerStats ServerStats hold various statistics that we will periodically send out.
type ServerStats struct {
	ActiveAccounts   int32                    `json:"active_accounts"`
	ActiveServers    *int32                   `json:"active_servers,omitempty"`
	Connections      int32                    `json:"connections"`
	Cores            int32                    `json:"cores"`
	Cpu              float64                  `json:"cpu"`
	Gateways         *[]GatewayStat           `json:"gateways,omitempty"`
	Jetstream        *Nullable[JetStreamVarz] `json:"jetstream,omitempty"`
	Mem              int64                    `json:"mem"`
	Received         DataStats                `json:"received"`
	Routes           *[]RouteStat             `json:"routes,omitempty"`
	Sent             DataStats                `json:"sent"`
	SlowConsumers    int64                    `json:"slow_consumers"`
	Start            time.Time                `json:"start"`
	Subscriptions    int32                    `json:"subscriptions"`
	TotalConnections int32                    `json:"total_connections"`
}

func (o ServerStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_accounts"] = o.ActiveAccounts
	if o.ActiveServers != nil {
		toSerialize["active_servers"] = o.ActiveServers
	}
	toSerialize["connections"] = o.Connections
	toSerialize["cores"] = o.Cores
	toSerialize["cpu"] = o.Cpu
	if o.Gateways != nil {
		toSerialize["gateways"] = o.Gateways
	}
	if o.Jetstream != nil && !o.Jetstream.IsNull() {
		toSerialize["jetstream"] = o.Jetstream.Val
	}
	toSerialize["mem"] = o.Mem
	toSerialize["received"] = o.Received
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	toSerialize["sent"] = o.Sent
	toSerialize["slow_consumers"] = o.SlowConsumers
	toSerialize["start"] = o.Start
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["total_connections"] = o.TotalConnections
	return toSerialize, nil
}
