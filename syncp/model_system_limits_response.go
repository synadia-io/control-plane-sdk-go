/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SystemLimitsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemLimitsResponse{}

// SystemLimitsResponse struct for SystemLimitsResponse
type SystemLimitsResponse struct {
	Allocated TenantLimits `json:"allocated"`
	Available TenantLimits `json:"available"`
	Total     TenantLimits `json:"total"`
}

func (o SystemLimitsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allocated"] = o.Allocated
	toSerialize["available"] = o.Available
	toSerialize["total"] = o.Total
	return toSerialize, nil
}
