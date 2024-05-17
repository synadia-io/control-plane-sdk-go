/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SigningKeyListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyListResponse{}

// SigningKeyListResponse struct for SigningKeyListResponse
type SigningKeyListResponse struct {
	Items []SigningKeyViewResponse `json:"items"`
}

func (o SigningKeyListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
