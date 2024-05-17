/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the TeamAppUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamAppUserInfo{}

// TeamAppUserInfo struct for TeamAppUserInfo
type TeamAppUserInfo struct {
	Id                string `json:"id"`
	PendingInvitation bool   `json:"pending_invitation"`
	RoleId            string `json:"role_id"`
	RoleName          string `json:"role_name"`
}

func (o TeamAppUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["pending_invitation"] = o.PendingInvitation
	toSerialize["role_id"] = o.RoleId
	toSerialize["role_name"] = o.RoleName
	return toSerialize, nil
}
