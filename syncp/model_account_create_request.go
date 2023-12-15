/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AccountCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCreateRequest{}

// AccountCreateRequest struct for AccountCreateRequest
type AccountCreateRequest struct {
	JwtSettings          *AccountJWTSettings `json:"jwt_settings,omitempty"`
	Name                 string              `json:"name"`
	UserJwtExpiresInSecs *int64              `json:"user_jwt_expires_in_secs,omitempty"`
}

func (o AccountCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.JwtSettings != nil {
		toSerialize["jwt_settings"] = o.JwtSettings
	}
	toSerialize["name"] = o.Name
	if o.UserJwtExpiresInSecs != nil {
		toSerialize["user_jwt_expires_in_secs"] = o.UserJwtExpiresInSecs
	}
	return toSerialize, nil
}
