/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AppUserAssignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAssignRequest{}

// AppUserAssignRequest struct for AppUserAssignRequest
type AppUserAssignRequest struct {
	RoleId string `json:"role_id"`
}

func (o AppUserAssignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role_id"] = o.RoleId
	return toSerialize, nil
}
