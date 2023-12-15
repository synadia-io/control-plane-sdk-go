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

// checks if the AppUserAccessTokenCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAccessTokenCreateResponse{}

// AppUserAccessTokenCreateResponse struct for AppUserAccessTokenCreateResponse
type AppUserAccessTokenCreateResponse struct {
	Created time.Time        `json:"created"`
	Expires Nullable[string] `json:"expires"`
	Id      string           `json:"id"`
	Name    string           `json:"name"`
	Token   string           `json:"token"`
}

func (o AppUserAccessTokenCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	if !o.Expires.IsNull() {
		toSerialize["expires"] = o.Expires.Val
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["token"] = o.Token
	return toSerialize, nil
}
