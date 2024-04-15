/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the ServiceAccountListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountListResponse{}

// ServiceAccountListResponse struct for ServiceAccountListResponse
type ServiceAccountListResponse struct {
	Items []ServiceAccountViewResponse `json:"items"`
}

func (o ServiceAccountListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}