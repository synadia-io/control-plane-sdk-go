/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AccountJWTSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountJWTSettings{}

// AccountJWTSettings struct for AccountJWTSettings
type AccountJWTSettings struct {
	Info
	Authorization *ExternalAuthorization        `json:"authorization,omitempty"`
	Limits        *OperatorLimits               `json:"limits,omitempty"`
	Mappings      *map[string][]WeightedMapping `json:"mappings,omitempty"`
	// RevocationList is used to store a mapping of public keys to unix timestamps
	Revocations *map[string]int64 `json:"revocations,omitempty"`
	// TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments
	Tags *Nullable[[]string] `json:"tags,omitempty"`
}

func (o AccountJWTSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Authorization != nil {
		toSerialize["authorization"] = o.Authorization
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.Revocations != nil {
		toSerialize["revocations"] = o.Revocations
	}
	if o.Tags != nil && !o.Tags.IsNull() {
		toSerialize["tags"] = o.Tags.Val
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.InfoUrl != nil {
		toSerialize["info_url"] = o.InfoUrl
	}
	return toSerialize, nil
}
