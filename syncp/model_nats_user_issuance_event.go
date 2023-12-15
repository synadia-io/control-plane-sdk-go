/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the NatsUserIssuanceEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserIssuanceEvent{}

// NatsUserIssuanceEvent struct for NatsUserIssuanceEvent
type NatsUserIssuanceEvent struct {
	AppUser AppUserInfo               `json:"app_user"`
	Exp     *int64                    `json:"exp,omitempty"`
	Iat     int64                     `json:"iat"`
	Jti     string                    `json:"jti"`
	Type    NatsUserIssuanceEventType `json:"type"`
}

func (o NatsUserIssuanceEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_user"] = o.AppUser
	if o.Exp != nil {
		toSerialize["exp"] = o.Exp
	}
	toSerialize["iat"] = o.Iat
	toSerialize["jti"] = o.Jti
	toSerialize["type"] = o.Type
	return toSerialize, nil
}