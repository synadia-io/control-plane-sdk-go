/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the PeerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeerInfo{}

// PeerInfo PeerInfo shows information about all the peers in the cluster that are supporting the stream or consumer.
type PeerInfo struct {
	Active  int64   `json:"active"`
	Current bool    `json:"current"`
	Lag     *uint64 `json:"lag,omitempty"`
	Name    string  `json:"name"`
	Offline *bool   `json:"offline,omitempty"`
}

func (o PeerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["current"] = o.Current
	if o.Lag != nil {
		toSerialize["lag"] = o.Lag
	}
	toSerialize["name"] = o.Name
	if o.Offline != nil {
		toSerialize["offline"] = o.Offline
	}
	return toSerialize, nil
}
