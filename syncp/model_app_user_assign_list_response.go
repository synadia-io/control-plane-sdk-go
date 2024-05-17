/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AppUserAssignListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAssignListResponse{}

// AppUserAssignListResponse struct for AppUserAssignListResponse
type AppUserAssignListResponse struct {
	Items []AppUserAssignResponse `json:"items"`
}

func (o AppUserAssignListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
