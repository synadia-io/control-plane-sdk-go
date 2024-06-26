/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the StreamShareViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamShareViewResponse{}

// StreamShareViewResponse struct for StreamShareViewResponse
type StreamShareViewResponse struct {
	TargetAccountNkeyPublic string `json:"target_account_nkey_public"`
}

func (o StreamShareViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_account_nkey_public"] = o.TargetAccountNkeyPublic
	return toSerialize, nil
}
