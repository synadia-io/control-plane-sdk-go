/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the ClaimsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimsData{}

// ClaimsData ClaimsData is the base struct for all claims
type ClaimsData struct {
	Aud  *string `json:"aud,omitempty"`
	Exp  *int64  `json:"exp,omitempty"`
	Iat  *int64  `json:"iat,omitempty"`
	Iss  *string `json:"iss,omitempty"`
	Jti  *string `json:"jti,omitempty"`
	Name *string `json:"name,omitempty"`
	Nbf  *int64  `json:"nbf,omitempty"`
	Sub  *string `json:"sub,omitempty"`
}

func (o ClaimsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Aud != nil {
		toSerialize["aud"] = o.Aud
	}
	if o.Exp != nil {
		toSerialize["exp"] = o.Exp
	}
	if o.Iat != nil {
		toSerialize["iat"] = o.Iat
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.Jti != nil {
		toSerialize["jti"] = o.Jti
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Nbf != nil {
		toSerialize["nbf"] = o.Nbf
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}
	return toSerialize, nil
}
