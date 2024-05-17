/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"time"
)

// checks if the SigningKeyViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyViewResponse{}

// SigningKeyViewResponse struct for SigningKeyViewResponse
type SigningKeyViewResponse struct {
	Created     time.Time  `json:"created"`
	Current     bool       `json:"current"`
	Disabled    bool       `json:"disabled"`
	DisabledAt  *time.Time `json:"disabled_at,omitempty"`
	GroupId     string     `json:"group_id"`
	Id          string     `json:"id"`
	MissingSeed bool       `json:"missing_seed"`
	PublicKey   string     `json:"public_key"`
	Rotated     bool       `json:"rotated"`
	RotatedAt   *time.Time `json:"rotated_at,omitempty"`
}

func (o SigningKeyViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["current"] = o.Current
	toSerialize["disabled"] = o.Disabled
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
	}
	toSerialize["group_id"] = o.GroupId
	toSerialize["id"] = o.Id
	toSerialize["missing_seed"] = o.MissingSeed
	toSerialize["public_key"] = o.PublicKey
	toSerialize["rotated"] = o.Rotated
	if o.RotatedAt != nil {
		toSerialize["rotated_at"] = o.RotatedAt
	}
	return toSerialize, nil
}
