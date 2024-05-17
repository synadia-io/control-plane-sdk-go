/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the NatsUserIssuancesListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NatsUserIssuancesListResponse{}

// NatsUserIssuancesListResponse struct for NatsUserIssuancesListResponse
type NatsUserIssuancesListResponse struct {
	Items []NatsUserIssuanceInfo `json:"items"`
}

func (o NatsUserIssuancesListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
