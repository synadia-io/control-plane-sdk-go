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

// checks if the NatsUserIssuanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserIssuanceInfo{}

// NatsUserIssuanceInfo struct for NatsUserIssuanceInfo
type NatsUserIssuanceInfo struct {
	Created     time.Time `json:"created"`
	EventsCount int32     `json:"events_count"`
	ExpMax      *int64    `json:"exp_max,omitempty"`
	IatMax      int64     `json:"iat_max"`
	Id          string    `json:"id"`
	Iss         string    `json:"iss"`
	Name        string    `json:"name"`
	Sub         string    `json:"sub"`
}

func (o NatsUserIssuanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["events_count"] = o.EventsCount
	if o.ExpMax != nil {
		toSerialize["exp_max"] = o.ExpMax
	}
	toSerialize["iat_max"] = o.IatMax
	toSerialize["id"] = o.Id
	toSerialize["iss"] = o.Iss
	toSerialize["name"] = o.Name
	toSerialize["sub"] = o.Sub
	return toSerialize, nil
}
