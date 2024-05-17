/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the NatsUserHTTPGWTokenCreateReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserHTTPGWTokenCreateReply{}

// NatsUserHTTPGWTokenCreateReply struct for NatsUserHTTPGWTokenCreateReply
type NatsUserHTTPGWTokenCreateReply struct {
	Token string `json:"token"`
}

func (o NatsUserHTTPGWTokenCreateReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}
