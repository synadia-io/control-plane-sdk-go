/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AccountUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUpdateRequest{}

// AccountUpdateRequest struct for AccountUpdateRequest
type AccountUpdateRequest struct {
	JwtSettings          *AccountJWTSettingsPatch `json:"jwt_settings,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	UserJwtExpiresInSecs *int64                   `json:"user_jwt_expires_in_secs,omitempty"`
}

func (o AccountUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.JwtSettings != nil {
		toSerialize["jwt_settings"] = o.JwtSettings
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UserJwtExpiresInSecs != nil {
		toSerialize["user_jwt_expires_in_secs"] = o.UserJwtExpiresInSecs
	}
	return toSerialize, nil
}
