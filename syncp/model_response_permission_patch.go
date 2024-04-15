/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the ResponsePermissionPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePermissionPatch{}

// ResponsePermissionPatch struct for ResponsePermissionPatch
type ResponsePermissionPatch struct {
	Max *int32 `json:"max,omitempty"`
	Ttl *int64 `json:"ttl,omitempty"`
}

func (o ResponsePermissionPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}