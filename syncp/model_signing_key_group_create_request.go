/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SigningKeyGroupCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyGroupCreateRequest{}

// SigningKeyGroupCreateRequest struct for SigningKeyGroupCreateRequest
type SigningKeyGroupCreateRequest struct {
	Name  string                `json:"name"`
	Scope *UserPermissionLimits `json:"scope,omitempty"`
	Seed  *string               `json:"seed,omitempty"`
}

func (o SigningKeyGroupCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Seed != nil {
		toSerialize["seed"] = o.Seed
	}
	return toSerialize, nil
}
