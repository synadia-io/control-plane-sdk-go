/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the ExternalAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalAuthorization{}

// ExternalAuthorization struct for ExternalAuthorization
type ExternalAuthorization struct {
	AllowedAccounts []string `json:"allowed_accounts,omitempty"`
	AuthUsers       []string `json:"auth_users"`
	Xkey            *string  `json:"xkey,omitempty"`
}

func (o ExternalAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if len(o.AllowedAccounts) != 0 {
		toSerialize["allowed_accounts"] = o.AllowedAccounts
	}
	if o.AuthUsers != nil {
		toSerialize["auth_users"] = o.AuthUsers
	}
	if o.Xkey != nil {
		toSerialize["xkey"] = o.Xkey
	}
	return toSerialize, nil
}
