/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

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

func (o NatsCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connections"] = o.Connections
	toSerialize["incomingGateways"] = o.IncomingGateways
	toSerialize["name"] = o.Name
	toSerialize["nodeCount"] = o.NodeCount
	toSerialize["outgoingGateways"] = o.OutgoingGateways
	return toSerialize, nil
}
