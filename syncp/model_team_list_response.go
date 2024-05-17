/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the TeamListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamListResponse{}

// TeamListResponse struct for TeamListResponse
type TeamListResponse struct {
	Items []TeamViewResponse `json:"items"`
}

func (o TeamListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
