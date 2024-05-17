/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the OperatorClaims type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorClaims{}

// OperatorClaims OperatorClaims define the data for an operator JWT
type OperatorClaims struct {
	ClaimsData
	Nats *Operator `json:"nats,omitempty"`
}

func (o OperatorClaims) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Nats != nil {
		toSerialize["nats"] = o.Nats
	}
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
