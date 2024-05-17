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

// checks if the ServiceAccountViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountViewResponse{}

// ServiceAccountViewResponse struct for ServiceAccountViewResponse
type ServiceAccountViewResponse struct {
	Created    time.Time        `json:"created"`
	Id         string           `json:"id"`
	LastActive Nullable[string] `json:"last_active"`
	Name       string           `json:"name"`
	// resource keys must be in format ServiceAccountResourceScope:resource_id
	Resources map[string]AppUserAssignResponse `json:"resources"`
	RoleId    string                           `json:"role_id"`
	Scope     ServiceAccountScope              `json:"scope"`
	TeamId    *string                          `json:"team_id,omitempty"`
}

func (o ServiceAccountViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	if !o.LastActive.IsNull() {
		toSerialize["last_active"] = o.LastActive.Val
	}
	toSerialize["name"] = o.Name
	toSerialize["resources"] = o.Resources
	toSerialize["role_id"] = o.RoleId
	toSerialize["scope"] = o.Scope
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	return toSerialize, nil
}
