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

// checks if the AppUserAccessTokenViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAccessTokenViewResponse{}

// AppUserAccessTokenViewResponse struct for AppUserAccessTokenViewResponse
type AppUserAccessTokenViewResponse struct {
	Created time.Time  `json:"created"`
	Expires *time.Time `json:"expires,omitempty"`
	Id      string     `json:"id"`
	Name    string     `json:"name"`
}

func (o AppUserAccessTokenViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}
