/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AuthzRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthzRequest{}

// AuthzRequest struct for AuthzRequest
type AuthzRequest struct {
	ResourceId string `json:"resource_id"`
	Service    string `json:"service"`
}

func (o AuthzRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["service"] = o.Service
	return toSerialize, nil
}
