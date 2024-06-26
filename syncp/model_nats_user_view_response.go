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

// checks if the NatsUserViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserViewResponse{}

// NatsUserViewResponse struct for NatsUserViewResponse
type NatsUserViewResponse struct {
	Account          AccountInfo         `json:"account"`
	Claims           UserClaims          `json:"claims"`
	Created          time.Time           `json:"created"`
	Id               string              `json:"id"`
	Jwt              string              `json:"jwt"`
	JwtExpiresAtMax  int64               `json:"jwt_expires_at_max"`
	JwtExpiresInSecs int64               `json:"jwt_expires_in_secs"`
	JwtSettings      NatsUserJwtSettings `json:"jwt_settings"`
	Name             string              `json:"name"`
	SkGroupId        *string             `json:"sk_group_id,omitempty"`
	System           SystemInfo          `json:"system"`
	Team             TeamInfo            `json:"team"`
	UserPublicKey    string              `json:"user_public_key"`
}

func (o NatsUserViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	toSerialize["claims"] = o.Claims
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["jwt"] = o.Jwt
	toSerialize["jwt_expires_at_max"] = o.JwtExpiresAtMax
	toSerialize["jwt_expires_in_secs"] = o.JwtExpiresInSecs
	toSerialize["jwt_settings"] = o.JwtSettings
	toSerialize["name"] = o.Name
	if o.SkGroupId != nil {
		toSerialize["sk_group_id"] = o.SkGroupId
	}
	toSerialize["system"] = o.System
	toSerialize["team"] = o.Team
	toSerialize["user_public_key"] = o.UserPublicKey
	return toSerialize, nil
}
