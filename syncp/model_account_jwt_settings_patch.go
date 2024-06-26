/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AccountJWTSettingsPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountJWTSettingsPatch{}

// AccountJWTSettingsPatch struct for AccountJWTSettingsPatch
type AccountJWTSettingsPatch struct {
	Authorization *Nullable[ExternalAuthorizationPatch] `json:"authorization,omitempty"`
	Description   *Nullable[string]                     `json:"description,omitempty"`
	InfoUrl       *Nullable[string]                     `json:"info_url,omitempty"`
	Limits        *Nullable[OperatorLimitsPatch]        `json:"limits,omitempty"`
	Mappings      map[string][]WeightedMapping          `json:"mappings,omitempty"`
	Revocations   map[string]int64                      `json:"revocations,omitempty"`
	Tags          []string                              `json:"tags,omitempty"`
}

func (o AccountJWTSettingsPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Authorization != nil && !o.Authorization.IsNull() {
		toSerialize["authorization"] = o.Authorization.Val
	}
	if o.Description != nil && !o.Description.IsNull() {
		toSerialize["description"] = o.Description.Val
	}
	if o.InfoUrl != nil && !o.InfoUrl.IsNull() {
		toSerialize["info_url"] = o.InfoUrl.Val
	}
	if o.Limits != nil && !o.Limits.IsNull() {
		toSerialize["limits"] = o.Limits.Val
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.Revocations != nil {
		toSerialize["revocations"] = o.Revocations
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}
