/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the NatsUserRevocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserRevocationRequest{}

// NatsUserRevocationRequest struct for NatsUserRevocationRequest
type NatsUserRevocationRequest struct {
	Before int64 `json:"before"`
}

func (o NatsUserRevocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["before"] = o.Before
	return toSerialize, nil
}
