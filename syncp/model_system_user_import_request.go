/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SystemUserImportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemUserImportRequest{}

// SystemUserImportRequest struct for SystemUserImportRequest
type SystemUserImportRequest struct {
	Credential *string `json:"credential,omitempty"`
	Jwt        *string `json:"jwt,omitempty"`
	Seed       *string `json:"seed,omitempty"`
}

func (o SystemUserImportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	if o.Jwt != nil {
		toSerialize["jwt"] = o.Jwt
	}
	if o.Seed != nil {
		toSerialize["seed"] = o.Seed
	}
	return toSerialize, nil
}
